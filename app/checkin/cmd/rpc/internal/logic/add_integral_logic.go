package logic

import (
	"Luckify/app/checkin/model"
	"Luckify/common/xerr"
	"context"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"

	"Luckify/app/checkin/cmd/rpc/internal/svc"
	"Luckify/app/checkin/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddIntegralLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddIntegralLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddIntegralLogic {
	return &AddIntegralLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddIntegralLogic) AddIntegral(in *pb.AddIntegralReq) (*pb.AddIntegralResp, error) {
	dbIntegral := &model.Integral{}
	_ = copier.Copy(dbIntegral, in)
	err := l.svcCtx.IntegralModel.Insert(l.ctx, nil, dbIntegral)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_INTEGRAL_INSERT_ERROR), "IntegralModel Insert:%+v, error: %v", dbIntegral, err)
	}

	return &pb.AddIntegralResp{}, nil
}
