package logic

import (
	"context"

	"Luckify/app/comment/cmd/rpc/internal/svc"
	"Luckify/app/comment/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserCommentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserCommentLogic {
	return &GetUserCommentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserCommentLogic) GetUserComment(in *pb.GetUserCommentReq) (*pb.GetUserCommentResp, error) {
	// todo: add your logic here and delete this line

	return &pb.GetUserCommentResp{}, nil
}
