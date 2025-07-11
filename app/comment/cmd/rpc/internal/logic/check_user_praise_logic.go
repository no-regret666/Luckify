package logic

import (
	"context"

	"Luckify/app/comment/cmd/rpc/internal/svc"
	"Luckify/app/comment/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type CheckUserPraiseLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCheckUserPraiseLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CheckUserPraiseLogic {
	return &CheckUserPraiseLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CheckUserPraiseLogic) CheckUserPraise(in *pb.CheckUserPraiseReq) (*pb.CheckUserPraiseResp, error) {
	
	return &pb.CheckUserPraiseResp{}, nil
}
