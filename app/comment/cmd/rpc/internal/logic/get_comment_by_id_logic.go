package logic

import (
	"Luckify/common/xerr"
	"context"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"

	"Luckify/app/comment/cmd/rpc/internal/svc"
	"Luckify/app/comment/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCommentByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetCommentByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCommentByIdLogic {
	return &GetCommentByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetCommentByIdLogic) GetCommentById(in *pb.GetCommentByIdReq) (*pb.GetCommentByIdResp, error) {
	dbComment, err := l.svcCtx.CommentModel.FindOne(l.ctx, in.Id)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_FIND_COMMENT_ERROR), "Failed to find comment by id, err: %v", err)
	}

	pbComment := &pb.Comment{}
	_ = copier.Copy(pbComment, dbComment)
	pbComment.CreateTime = dbComment.CreateTime.Unix()
	pbComment.UpdateTime = dbComment.UpdateTime.Unix()

	return &pb.GetCommentByIdResp{
		Comment: pbComment,
	}, nil
}
