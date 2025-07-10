package comment

import (
	"Luckify/app/comment/cmd/rpc/comment"
	"Luckify/app/comment/model"
	"context"
	"github.com/pkg/errors"

	"Luckify/app/comment/cmd/api/internal/svc"
	"Luckify/app/comment/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteCommentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 删除评论
func NewDeleteCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteCommentLogic {
	return &DeleteCommentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteCommentLogic) DeleteComment(req *types.CommentDelReq) (resp *types.CommentDelResp, err error) {
	_, err = l.svcCtx.CommentRpc.DelComment(l.ctx, &comment.DelCommentReq{
		Id: req.Id,
	})
	if err != nil {
		return nil, errors.Wrapf(model.ErrDeleteComment, "DeleteComment rpc error:%v", err)
	}

	return
}
