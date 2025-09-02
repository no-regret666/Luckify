package logic

import (
	"Luckify/app/mqueue/cmd/job/internal/svc"
	"Luckify/app/mqueue/cmd/job/jobtype"
	"context"
	"github.com/hibiken/asynq"
)

type CronJob struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCronJob(ctx context.Context, svcCtx *svc.ServiceContext) *CronJob {
	return &CronJob{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CronJob) Register() *asynq.ServeMux {
	mux := asynq.NewServeMux()

	// schedule job
	mux.Handle(jobtype.ScheduleLotteryDraw, NewLotteryDrawHandler(l.svcCtx))
	mux.Handle(jobtype.ScheduleWishCheckin, NewWishCheckinHandler(l.svcCtx))
	return mux
}
