package logic

import (
	"Luckify/app/lottery/model"
	"Luckify/common/constants"
	"Luckify/common/utility"
	"Luckify/common/xerr"
	"context"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"strconv"

	"Luckify/app/lottery/cmd/rpc/internal/svc"
	"Luckify/app/lottery/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddInstantLotteryParticipationLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddInstantLotteryParticipationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddInstantLotteryParticipationLogic {
	return &AddInstantLotteryParticipationLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddInstantLotteryParticipationLogic) AddInstantLotteryParticipation(in *pb.AddInstantLotteryParticipationReq) (*pb.AddInstantLotteryParticipationResp, error) {
	mutexKey := constants.InstantLotteryRedisKey + strconv.Itoa(int(in.LotteryId))
	mutex := l.svcCtx.RedsyncClient.NewMutex(mutexKey)
	if err := mutex.Lock(); err != nil {
		return nil, errors.Wrapf(err, "failed to lock mutex %s", mutexKey)
	}
	defer mutex.Unlock()

	// 1. 检查参与的抽奖是否为即抽即中类型
	dbLottery, err := l.svcCtx.LotteryModel.FindOne(l.ctx, in.LotteryId)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_FIND_LOTTERY_BY_ID_ERR), "failed to find lottery by id %d", in.LotteryId)
	}
	if dbLottery.AnnounceType != constants.AnnounceTypeInstant {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.ANNOUNCE_LOTTERY_FAIL), "lottery %d is not instant lottery", in.LotteryId)
	}

	// 2. 检查用户是否已经参与
	dbPartitions, err := l.svcCtx.LotteryParticipationModel.FindOneByLotteryIdUserId(l.ctx, in.LotteryId, in.UserId)
	if err != nil && !errors.Is(err, model.ErrNotFound) {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_FIND_PARTICIPATION_BY_LOTTERY_ID_AND_USER_ID_ERR), "failed to find lottery participation by lotteryId %d and userId %d", in.LotteryId, in.UserId)
	}

	if dbPartitions != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_CHECK_IS_PARTICIPATED), "user %d has participated in lottery %d", in.UserId, in.LotteryId)
	}

	// 3. 参与抽奖
	randomCode := utility.Random(constants.ProbabilityMax)
	dbPrizeList, err := l.svcCtx.PrizeModel.FindByLotteryId(l.ctx, in.LotteryId)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_FIND_PRIZES_BY_LOTTERY_ID_ERR), "failed to find prize by lotteryId %d", in.LotteryId)
	}

	isWon := false
	wonPrize := &model.Prize{}
	for _, dbPrize := range dbPrizeList {
		probability := constants.ProbabilityMap[int(dbPrize.Level)]
		if randomCode >= probability[0] && randomCode < probability[1] && dbPrize.Stock > 0 {
			isWon = true
			wonPrize = dbPrize
			break
		}
	}
	if !isWon {
		err = l.svcCtx.LotteryParticipationModel.Insert(l.ctx, nil, &model.LotteryParticipation{
			LotteryId: in.LotteryId,
			UserId:    in.UserId,
			IsWon:     constants.PrizeNotWon,
		})
		if err != nil {
			return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_PARTICIPATE_LOTTERY), "failed to participate lottery %d", in.LotteryId)
		}
		return &pb.AddInstantLotteryParticipationResp{}, nil
	}

	err = l.svcCtx.TransactCtx(l.ctx, func(db *gorm.DB) error {
		err = l.svcCtx.PrizeModel.DecrStock(l.ctx, wonPrize.Id)
		if err != nil {
			return errors.Wrapf(xerr.NewErrCode(xerr.DB_DECR_PRIZE_STOCK_ERR), "failed to decr prize count by prizeId %d", wonPrize.Id)
		}

		err = l.svcCtx.LotteryParticipationModel.Insert(l.ctx, db, &model.LotteryParticipation{
			LotteryId: in.LotteryId,
			UserId:    in.UserId,
			IsWon:     constants.PrizeHasWon,
			PrizeId:   wonPrize.Id,
		})
		if err != nil {
			return errors.Wrapf(xerr.NewErrCode(xerr.DB_PARTICIPATE_LOTTERY), "failed to participate lottery %d", in.LotteryId)
		}
		return nil
	})
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_TRANSACTION_ERROR), "failed to participate lottery %d", in.LotteryId)
	}
	return &pb.AddInstantLotteryParticipationResp{
		Id: wonPrize.Id,
	}, nil
}
