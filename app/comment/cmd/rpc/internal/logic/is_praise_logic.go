package logic

import (
	"Luckify/common/xerr"
	"context"
	"github.com/pkg/errors"

	"Luckify/app/comment/cmd/rpc/internal/svc"
	"Luckify/app/comment/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type IsPraiseLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewIsPraiseLogic(ctx context.Context, svcCtx *svc.ServiceContext) *IsPraiseLogic {
	return &IsPraiseLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *IsPraiseLogic) IsPraise(in *pb.IsPraiseReq) (*pb.IsPraiseResp, error) {
	dbPraise, err := l.svcCtx.PraiseModel.IsPraise(l.ctx, in.CommentId, in.UserId)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_FIND_PRAISE_ERROR), "PraiseModel IsPraise fail,req: %+v,err: %v", in, err)
	}

	return &pb.IsPraiseResp{
		PraiseId: dbPraise.Id,
	}, nil
}
