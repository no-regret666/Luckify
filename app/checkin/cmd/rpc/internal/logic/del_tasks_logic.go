package logic

import (
	"Luckify/common/xerr"
	"context"
	"github.com/pkg/errors"

	"Luckify/app/checkin/cmd/rpc/internal/svc"
	"Luckify/app/checkin/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelTasksLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelTasksLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelTasksLogic {
	return &DelTasksLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DelTasksLogic) DelTasks(in *pb.DelTasksReq) (*pb.DelTasksResp, error) {
	err := l.svcCtx.TasksModel.Delete(l.ctx, nil, in.Id)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_TASKS_DELETE_ERROR), "Failed to delete tasks data: %+v,err:%v", in, err)
	}

	return &pb.DelTasksResp{}, nil
}
