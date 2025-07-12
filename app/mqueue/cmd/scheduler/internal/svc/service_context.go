package svc

import (
	"Luckify/app/mqueue/cmd/scheduler/internal/config"
	"github.com/hibiken/asynq"
	"github.com/zeromicro/go-zero/core/logx"
	"time"
)

type ServiceContext struct {
	config    config.Config
	Scheduler *asynq.Scheduler
}

func newScheduler(c config.Config) *asynq.Scheduler {
	location, _ := time.LoadLocation("Asia/Shanghai")
	return asynq.NewScheduler(asynq.RedisClientOpt{
		Addr:     c.Redis.Host,
		Password: c.Redis.Pass,
	}, &asynq.SchedulerOpts{
		Location: location,
		PostEnqueueFunc: func(info *asynq.TaskInfo, err error) {
			if err != nil {
				logx.Errorf("[Scheduler PostEnqueueFunc] <<<<<<<<== ==>>>>>>>> err: %+v,info:%+v", err, info)
			} else {
				logx.Infof(" [Scheduler PostEnqueueFunc] <<<<<<<<== ==>>>>>>>> info: %+v", info)
			}
		},
	})
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		config:    c,
		Scheduler: newScheduler(c),
	}
}
