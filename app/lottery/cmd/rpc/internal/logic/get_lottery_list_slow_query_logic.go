package logic

import (
	"context"

	"Luckify/app/lottery/cmd/rpc/internal/svc"
	"Luckify/app/lottery/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetLotteryListSlowQueryLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetLotteryListSlowQueryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetLotteryListSlowQueryLogic {
	return &GetLotteryListSlowQueryLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetLotteryListSlowQueryLogic) GetLotteryListSlowQuery(in *pb.GetLotteryListSlowQueryReq) (*pb.GetLotteryListSlowQueryResp, error) {
	// todo: add your logic here and delete this line

	return &pb.GetLotteryListSlowQueryResp{}, nil
}
