package logic

import (
	"Luckify/common/xerr"
	"context"
	"github.com/pkg/errors"

	"Luckify/app/comment/cmd/rpc/internal/svc"
	"Luckify/app/comment/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type IsPraiseListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewIsPraiseListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *IsPraiseListLogic {
	return &IsPraiseListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *IsPraiseListLogic) IsPraiseList(in *pb.IsPraiseListReq) (*pb.IsPraiseListResp, error) {
	list, err := l.svcCtx.PraiseModel.IsPraiseList(l.ctx, in.CommentId, in.UserId)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_GET_PRAISE_LIST_ERROR), "PraiseModel IsPraiseList fail,req:%+v,err:%v", in, err)
	}

	return &pb.IsPraiseListResp{
		PraiseId: list,
	}, nil
}
