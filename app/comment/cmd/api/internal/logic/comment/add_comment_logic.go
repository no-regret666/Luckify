package comment

import (
	"Luckify/app/comment/cmd/rpc/comment"
	"Luckify/app/comment/model"
	"Luckify/common/utility"
	"context"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"

	"Luckify/app/comment/cmd/api/internal/svc"
	"Luckify/app/comment/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddCommentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 增加评论
func NewAddCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddCommentLogic {
	return &AddCommentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddCommentLogic) AddComment(req *types.CommentAddReq) (resp *types.CommentAddResp, err error) {
	pbReq := &comment.AddCommentReq{}
	_ = copier.Copy(pbReq, req)
	pbReq.UserId = utility.GetUserIdFromCtx(l.ctx)
	_, err = l.svcCtx.CommentRpc.AddComment(l.ctx, pbReq)
	if err != nil {
		return nil, errors.Wrapf(model.ErrInsertComment, "AddComment rpc error:%v", err)
	}

	return
}
