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

type GetCommentLastIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取评论lastId
func NewGetCommentLastIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCommentLastIdLogic {
	return &GetCommentLastIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetCommentLastIdLogic) GetCommentLastId(req *types.GetCommentLastIdReq) (resp *types.GetCommentLastIdResp, err error) {
	pbResp, err := l.svcCtx.CommentRpc.GetCommentLastId(l.ctx, &comment.GetCommentLastIdReq{})
	if err != nil {
		return nil, errors.Wrapf(model.ErrGetLastCommentId, "rpc error:%v", err)
	}

	return &types.GetCommentLastIdResp{
		LastId: pbResp.LastId,
	}, nil
}
