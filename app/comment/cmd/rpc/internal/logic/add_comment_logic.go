package logic

import (
	"Luckify/app/comment/model"
	"Luckify/common/xerr"
	"context"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"

	"Luckify/app/comment/cmd/rpc/internal/svc"
	"Luckify/app/comment/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddCommentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddCommentLogic {
	return &AddCommentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddCommentLogic) AddComment(in *pb.AddCommentReq) (*pb.AddCommentResp, error) {
	dbComment := &model.Comment{}
	_ = copier.Copy(dbComment, in)
	err := l.svcCtx.CommentModel.Insert(l.ctx, nil, dbComment)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_INSERT_COMMENT_ERROR), "AddComment err: %v", err)
	}

	return &pb.AddCommentResp{}, nil
}
