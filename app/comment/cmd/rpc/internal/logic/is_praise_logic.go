package logic

import (
	"context"

	"Luckify/app/comment/cmd/rpc/internal/svc"
	"Luckify/app/comment/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type IsPraiseLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewIsPraiseLogic(ctx context.Context, svcCtx *svc.ServiceContext) *IsPraiseLogic {
	return &IsPraiseLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *IsPraiseLogic) IsPraise(in *pb.IsPraiseReq) (*pb.IsPraiseResp, error) {
	// todo: add your logic here and delete this line

	return &pb.IsPraiseResp{}, nil
}
