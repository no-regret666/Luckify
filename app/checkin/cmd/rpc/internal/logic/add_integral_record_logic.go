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

type AddIntegralRecordLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddIntegralRecordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddIntegralRecordLogic {
	return &AddIntegralRecordLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddIntegralRecordLogic) AddIntegralRecord(in *pb.AddIntegralRecordReq) (*pb.AddIntegralRecordResp, error) {
	dbIntegralRecord := &model.IntegralRecord{}
	_ = copier.Copy(dbIntegralRecord, in)
	err := l.svcCtx.IntegralRecordModel.Insert(l.ctx, nil, dbIntegralRecord)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_INTEGRAL_RECORD_INSERT_ERROR), "IntegralRecordModel Insert:%+v, error: %v", dbIntegralRecord, err)
	}

	return &pb.AddIntegralRecordResp{}, nil
}
