package logic

import (
	"Luckify/app/lottery/cmd/rpc/internal/svc"
	"Luckify/app/lottery/cmd/rpc/pb"
	"Luckify/app/lottery/model"
	"Luckify/common/constants"
	"Luckify/common/xerr"
	"context"
	"github.com/IBM/sarama"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"google.golang.org/protobuf/proto"
	"strconv"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type MessageValue struct {
	LotteryId int64  `json:"lottery_id"`
	UserId    int64  `json:"user_id"`
	Timestamp int64  `json:"timestamp"`
	RequestId string `json:"request_id"`
}

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
	//// 使用分布式锁控制并发，同一时刻只有一个实例可以处理对该LotteryId的抽奖请求，防止超卖
	//mutexKey := constants.InstantLotteryRedisKey + strconv.Itoa(int(in.LotteryId))
	//mutex := l.svcCtx.RedsyncClient.NewMutex(mutexKey)
	//if err := mutex.Lock(); err != nil {
	//	return nil, errors.Wrapf(err, "failed to lock mutex %s", mutexKey)
	//}
	//defer mutex.Unlock()

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

	msgValue := &pb.MessageValue{
		LotteryId: in.LotteryId,
		UserId:    in.UserId,
		Timestamp: time.Now().Unix(),
		RequestId: uuid.New().String(), // 可以考虑换成雪花 ID
	}

	messageValue, _ := proto.Marshal(msgValue)
	msg := &sarama.ProducerMessage{
		Topic: l.svcCtx.Config.Kafka.Topic,
		Key:   sarama.StringEncoder(strconv.Itoa(int(in.LotteryId))),
		Value: sarama.ByteEncoder(messageValue),
	}

	go func() {
		partition, offset, err := l.svcCtx.Producer.SendMessage(msg)
		if err != nil {
			l.Logger.Errorf("Async Send Kafka message failed: %v", err)
		} else {
			l.Logger.Infof("Kafka async sent: partition %d offset %d", partition, offset)
		}
	}()

	return &pb.AddInstantLotteryParticipationResp{
		Id: 0,
	}, nil
}
