package logic

import (
	"context"

	"Luckify/app/lottery/cmd/rpc/internal/svc"
	"Luckify/app/lottery/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserAllListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserAllListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserAllListLogic {
	return &GetUserAllListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserAllListLogic) GetUserAllList(in *pb.GetUserAllListReq) (*pb.GetUserAllListResp, error) {
	// todo: add your logic here and delete this line

	return &pb.GetUserAllListResp{}, nil
}
