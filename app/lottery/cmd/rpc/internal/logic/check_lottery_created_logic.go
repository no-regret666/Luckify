package logic

import (
	"context"

	"Luckify/app/lottery/cmd/rpc/internal/svc"
	"Luckify/app/lottery/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type CheckLotteryCreatedLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCheckLotteryCreatedLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CheckLotteryCreatedLogic {
	return &CheckLotteryCreatedLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CheckLotteryCreatedLogic) CheckLotteryCreated(in *pb.CheckLotteryCreatedReq) (*pb.CheckLotteryCreatedResp, error) {
	// todo: add your logic here and delete this line

	return &pb.CheckLotteryCreatedResp{}, nil
}
