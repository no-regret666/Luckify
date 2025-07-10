package logic

import (
	"context"

	"Luckify/app/comment/cmd/rpc/internal/svc"
	"Luckify/app/comment/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelPraiseLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelPraiseLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelPraiseLogic {
	return &DelPraiseLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DelPraiseLogic) DelPraise(in *pb.DelPraiseReq) (*pb.DelPraiseResp, error) {
	// todo: add your logic here and delete this line

	return &pb.DelPraiseResp{}, nil
}
