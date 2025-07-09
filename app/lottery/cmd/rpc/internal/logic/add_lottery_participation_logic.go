package logic

import (
	"Luckify/app/lottery/model"
	"Luckify/common/constants"
	"Luckify/common/xerr"
	"context"
	"github.com/pkg/errors"

	"Luckify/app/lottery/cmd/rpc/internal/svc"
	"Luckify/app/lottery/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddLotteryParticipationLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddLotteryParticipationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddLotteryParticipationLogic {
	return &AddLotteryParticipationLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddLotteryParticipationLogic) AddLotteryParticipation(in *pb.AddLotteryParticipationReq) (*pb.AddLotteryParticipationResp, error) {
	dbLottery, err := l.svcCtx.LotteryModel.FindOne(l.ctx, in.LotteryId)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR_NOT_FOUND), "LotteryModel FindOne err: %+v", err)
	}

	if dbLottery.IsAnnounced == constants.LotteryHasAnnounced {
		return nil, errors.Wrapf(model.ErrHasAnnounced, "lottery has been announced")
	}

	lottery := &model.LotteryParticipation{
		UserId:    in.UserId,
		LotteryId: in.LotteryId,
	}
	err = l.svcCtx.LotteryParticipationModel.Insert(l.ctx, nil, lottery)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_PARTICIPATE_LOTTERY), "LotteryParticipation Insert err: %+v", err)
	}

	return &pb.AddLotteryParticipationResp{
		Id: lottery.Id,
	}, nil
}
