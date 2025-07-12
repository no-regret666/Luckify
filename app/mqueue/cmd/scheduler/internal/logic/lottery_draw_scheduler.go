package logic

import (
	"Luckify/app/mqueue/cmd/job/jobtype"
	"github.com/hibiken/asynq"
	"github.com/zeromicro/go-zero/core/logx"
)

// 定时任务调度器
func (l *MqueueScheduler) LotteryDrawScheduler() {
	// 每分钟投递一个lottery:draw任务到Redis,由worker消费
	task := asynq.NewTask(jobtype.ScheduleLotteryDraw, nil)
	entryID, err := l.svcCtx.Scheduler.Register("*/1 * * * *", task)

	if err != nil {
		logx.WithContext(l.ctx).Errorf("!!!MqueueSchedulerErr!!! ====> [LotteryDrawScheduler] registered err:%+v,task:%+v", err, task)
	}
	logx.Infof("[LotteryDrawScheduler] registered an entry: %q \n", entryID)
}
