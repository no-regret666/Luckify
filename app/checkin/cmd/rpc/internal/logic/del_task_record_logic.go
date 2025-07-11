package logic

import (
	"Luckify/common/xerr"
	"context"
	"github.com/pkg/errors"

	"Luckify/app/checkin/cmd/rpc/internal/svc"
	"Luckify/app/checkin/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelTaskRecordLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelTaskRecordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelTaskRecordLogic {
	return &DelTaskRecordLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DelTaskRecordLogic) DelTaskRecord(in *pb.DelTaskRecordReq) (*pb.DelTaskRecordResp, error) {
	err := l.svcCtx.TaskRecordModel.Delete(l.ctx, nil, in.Id)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_TASK_RECORD_DELETE_ERROR), "Failed to delete taskRecord data: %+v,err:%v", in, err)
	}

	return &pb.DelTaskRecordResp{}, nil
}
