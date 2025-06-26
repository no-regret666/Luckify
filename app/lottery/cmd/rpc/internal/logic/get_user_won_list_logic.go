package logic

import (
	"context"

	"Luckify/app/lottery/cmd/rpc/internal/svc"
	"Luckify/app/lottery/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserWonListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserWonListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserWonListLogic {
	return &GetUserWonListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserWonListLogic) GetUserWonList(in *pb.GetUserWonListReq) (*pb.GetUserWonListResp, error) {
	// todo: add your logic here and delete this line

	return &pb.GetUserWonListResp{}, nil
}
