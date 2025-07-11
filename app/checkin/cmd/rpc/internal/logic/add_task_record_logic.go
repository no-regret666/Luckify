package logic

import (
	"Luckify/app/checkin/model"
	"Luckify/common/constants"
	"Luckify/common/xerr"
	"context"
	"github.com/pkg/errors"

	"Luckify/app/checkin/cmd/rpc/internal/svc"
	"Luckify/app/checkin/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddTaskRecordLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddTaskRecordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddTaskRecordLogic {
	return &AddTaskRecordLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddTaskRecordLogic) AddTaskRecord(in *pb.AddTaskRecordReq) (*pb.AddTaskRecordResp, error) {
	task, err := l.svcCtx.TasksModel.FindOne(l.ctx, in.TaskId)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_TASK_FIND_ONE_ERROR), "TaskModel FindOne error: %v", err)
	}
	taskRecord := &model.TaskRecord{}
	taskRecord.TaskId = in.TaskId
	taskRecord.UserId = in.UserId
	taskRecord.Type = task.Type
	taskRecord.IsFinished = constants.TaskIsFinished
	err = l.svcCtx.TaskRecordModel.Insert(l.ctx, nil, taskRecord)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_TASK_RECORD_INSERT_ERROR), "TaskRecordModel Insert error: %v", err)
	}

	return &pb.AddTaskRecordResp{}, nil
}
