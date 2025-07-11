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

type SearchCommentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchCommentLogic {
	return &SearchCommentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SearchCommentLogic) SearchComment(in *pb.SearchCommentReq) (*pb.SearchCommentResp, error) {
	if in.LastId == 0 {
		id, err := l.svcCtx.CommentModel.GetCommentLastId(l.ctx)
		if err != nil {
			return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_GET_COMMENT_LAST_ID_ERROR), "model err: %v", err)
		}
		in.LastId = id + 1
	}

	dbList, err := l.svcCtx.CommentModel.GetCommentList(l.ctx, in.Limit, in.LastId, in.Sort)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_GET_COMMENT_LIST_ERROR), "model err: %v", err)
	}

	commentList := make([]*pb.Comment, 0)
	for _, dbComment := range dbList {
		commentItem := &pb.Comment{}
		_ = copier.Copy(commentItem, dbComment)
		commentItem.CreateTime = dbComment.CreateTime.Unix()
		commentItem.UpdateTime = dbComment.UpdateTime.Unix()
		commentList = append(commentList, commentItem)
	}

	return &pb.SearchCommentResp{
		Comment: commentList,
	}, nil
}
