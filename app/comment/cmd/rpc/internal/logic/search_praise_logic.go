package logic

import (
	"context"

	"Luckify/app/comment/cmd/rpc/internal/svc"
	"Luckify/app/comment/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchPraiseLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchPraiseLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchPraiseLogic {
	return &SearchPraiseLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SearchPraiseLogic) SearchPraise(in *pb.SearchPraiseReq) (*pb.SearchPraiseResp, error) {
	return &pb.SearchPraiseResp{}, nil
}
