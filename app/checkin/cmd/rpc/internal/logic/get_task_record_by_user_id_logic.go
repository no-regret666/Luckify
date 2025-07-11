package logic

import (
	"Luckify/common/xerr"
	"context"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"

	"Luckify/app/checkin/cmd/rpc/internal/svc"
	"Luckify/app/checkin/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetTaskRecordByUserIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetTaskRecordByUserIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTaskRecordByUserIdLogic {
	return &GetTaskRecordByUserIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetTaskRecordByUserIdLogic) GetTaskRecordByUserId(in *pb.GetTaskRecordByUserIdReq) (*pb.GetTaskRecordByUserIdResp, error) {
	dbTasks, err := l.svcCtx.TasksModel.FindAllTasks(l.ctx)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_TASKS_FIND_ALL), "Failed to find all tasks,err: %v", err)
	}

	var taskList []*pb.Tasks
	_ = copier.Copy(&taskList, &dbTasks)
	dbTaskRecords, err := l.svcCtx.TaskRecordModel.FindByUserId(l.ctx, in.UserId)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_TASK_RECORD_FIND_BY_USER_ID), "Failed to find taskRecordByUserId, err: %v", err)
	}

	isFinishedMap := make(map[int64]int64)
	for _, dbTaskRecord := range dbTaskRecords {
		isFinishedMap[dbTaskRecord.TaskId] = dbTaskRecord.IsFinished
	}

	for i := range taskList {
		if isFinished, ok := isFinishedMap[taskList[i].Id]; ok {
			taskList[i].IsFinished = isFinished
		}
	}

	return &pb.GetTaskRecordByUserIdResp{
		TaskList: taskList,
	}, nil
}
