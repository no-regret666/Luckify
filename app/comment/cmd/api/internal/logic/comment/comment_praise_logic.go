package comment

import (
	"Luckify/app/comment/cmd/rpc/comment"
	"Luckify/app/comment/model"
	"Luckify/common/utility"
	"context"
	"github.com/pkg/errors"

	"Luckify/app/comment/cmd/api/internal/svc"
	"Luckify/app/comment/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CommentPraiseLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 评论点赞/取消点赞
func NewCommentPraiseLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CommentPraiseLogic {
	return &CommentPraiseLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CommentPraiseLogic) CommentPraise(req *types.CommentPraiseReq) (resp *types.CommentPraiseResp, err error) {
	pbReq := &comment.PraiseCommentReq{}
	pbReq.CommentId = req.Id
	pbReq.UserId = utility.GetUserIdFromCtx(l.ctx)
	_, err = l.svcCtx.CommentRpc.PraiseComment(l.ctx, pbReq)
	if err != nil {
		return nil, errors.Wrapf(model.ErrInsertPraise, "PraiseComment rpc error:%v", err)
	}

	return
}
