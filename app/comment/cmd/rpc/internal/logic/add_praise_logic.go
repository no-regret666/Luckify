package logic

import (
	"context"

	"Luckify/app/comment/cmd/rpc/internal/svc"
	"Luckify/app/comment/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddPraiseLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddPraiseLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddPraiseLogic {
	return &AddPraiseLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddPraiseLogic) AddPraise(in *pb.AddPraiseReq) (*pb.AddPraiseResp, error) {
	// todo: add your logic here and delete this line

	return &pb.AddPraiseResp{}, nil
}
