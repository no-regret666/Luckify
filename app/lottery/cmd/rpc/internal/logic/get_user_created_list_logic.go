package logic

import (
	"context"

	"Luckify/app/lottery/cmd/rpc/internal/svc"
	"Luckify/app/lottery/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserCreatedListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserCreatedListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserCreatedListLogic {
	return &GetUserCreatedListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserCreatedListLogic) GetUserCreatedList(in *pb.GetUserCreatedListReq) (*pb.GetUserCreatedListResp, error) {
	// todo: add your logic here and delete this line

	return &pb.GetUserCreatedListResp{}, nil
}
