package logic

import (
	"Luckify/app/mqueue/cmd/scheduler/internal/svc"
	"context"
)

type MqueueScheduler struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCronScheduler(ctx context.Context, svcCtx *svc.ServiceContext) *MqueueScheduler {
	return &MqueueScheduler{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MqueueScheduler) Register() {
	l.LotteryDrawScheduler()
	l.WishCheckinScheduler()
}
