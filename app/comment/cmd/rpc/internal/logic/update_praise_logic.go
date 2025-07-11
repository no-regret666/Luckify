package logic

import (
	"Luckify/app/comment/model"
	"Luckify/common/xerr"
	"context"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"

	"Luckify/app/comment/cmd/rpc/internal/svc"
	"Luckify/app/comment/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdatePraiseLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdatePraiseLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdatePraiseLogic {
	return &UpdatePraiseLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdatePraiseLogic) UpdatePraise(in *pb.UpdatePraiseReq) (*pb.UpdatePraiseResp, error) {
	dbPraise := &model.Praise{}
	_ = copier.Copy(dbPraise, in)
	err := l.svcCtx.PraiseModel.Update(l.ctx, nil, dbPraise)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_UPDATE_PRAISE_ERROR), "UpdatePraise rpc err: %v", err)
	}

	return &pb.UpdatePraiseResp{}, nil
}
