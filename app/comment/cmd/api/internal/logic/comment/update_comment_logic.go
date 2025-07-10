package comment

import (
	"Luckify/app/comment/cmd/rpc/comment"
	"Luckify/app/comment/model"
	"context"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"

	"Luckify/app/comment/cmd/api/internal/svc"
	"Luckify/app/comment/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateCommentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 更新评论
func NewUpdateCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateCommentLogic {
	return &UpdateCommentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateCommentLogic) UpdateComment(req *types.CommentUpdateReq) (resp *types.CommentUpdateResp, err error) {
	pbReq := &comment.UpdateCommentReq{}
	_ = copier.Copy(pbReq, req)
	_, err = l.svcCtx.CommentRpc.UpdateComment(l.ctx, pbReq)
	if err != nil {
		return nil, errors.Wrapf(model.ErrUpdateComment, "UpdateComment rpc error:%v", err)
	}

	return &types.CommentUpdateResp{}, nil
}
