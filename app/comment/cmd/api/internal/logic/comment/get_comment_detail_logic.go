package comment

import (
	"Luckify/app/comment/cmd/rpc/pb"
	"Luckify/app/comment/model"
	"context"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"

	"Luckify/app/comment/cmd/api/internal/svc"
	"Luckify/app/comment/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCommentDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取评论详情
func NewGetCommentDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCommentDetailLogic {
	return &GetCommentDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetCommentDetailLogic) GetCommentDetail(req *types.CommentDetailReq) (resp *types.CommentDetailResp, err error) {
	pbResp, err := l.svcCtx.CommentRpc.GetCommentById(l.ctx, &pb.GetCommentByIdReq{
		Id: req.Id,
	})
	if err != nil {
		return nil, errors.Wrapf(model.ErrGetCommentDetail, "CommentRpc GetCommentById fail,req: %+v,err: %v", req, err)
	}

	respComment := new(types.Comment)
	_ = copier.Copy(respComment, &pbResp.Comment)

	return &types.CommentDetailResp{
		Comment: *respComment,
	}, nil
}
