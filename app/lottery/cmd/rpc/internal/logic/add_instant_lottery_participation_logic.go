package logic

import (
	"context"

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
	// todo: add your logic here and delete this line

	return &pb.AddInstantLotteryParticipationResp{}, nil
}
