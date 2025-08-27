package main

import (
	"Luckify/app/mqueue/cmd/scheduler/internal/config"
	"Luckify/app/mqueue/cmd/scheduler/internal/logic"
	"Luckify/app/mqueue/cmd/scheduler/internal/svc"
	"context"
	"flag"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"os"
)

var configFile = flag.String("f", "etc/scheduler.yaml", "Specify the config file")

func main() {
	flag.Parse()
	var c config.Config

	conf.MustLoad(*configFile, &c)

	logx.DisableStat()
	if err := c.SetUp(); err != nil {
		panic(err)
	}

	// 实例化
	svcContext := svc.NewServiceContext(c)
	ctx := context.Background()
	// 注册定时任务
	mqueueScheduler := logic.NewCronScheduler(ctx, svcContext)
	mqueueScheduler.Register()

	// 启动
	if err := svcContext.Scheduler.Run(); err != nil {
		logx.Errorf("!!!MqueueSchedulerErr!!! run err:%+v", err)
		os.Exit(1)
	}
}
