package logic

import (
	"Luckify/common/xerr"
	"context"
	"github.com/pkg/errors"

	"Luckify/app/checkin/cmd/rpc/internal/svc"
	"Luckify/app/checkin/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelIntegralRecordLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelIntegralRecordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelIntegralRecordLogic {
	return &DelIntegralRecordLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DelIntegralRecordLogic) DelIntegralRecord(in *pb.DelIntegralRecordReq) (*pb.DelIntegralRecordResp, error) {
	err := l.svcCtx.IntegralRecordModel.Delete(l.ctx, nil, in.Id)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_INTEGRAL_RECORD_DELETE_ERROR), "Failed to delete integralRecord data: %+v,err:%v", in, err)
	}

	return &pb.DelIntegralRecordResp{}, nil
}
