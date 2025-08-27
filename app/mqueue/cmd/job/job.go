package main

import (
	"Luckify/app/mqueue/cmd/job/internal/config"
	"Luckify/app/mqueue/cmd/job/internal/logic"
	"Luckify/app/mqueue/cmd/job/internal/svc"
	"context"
	"flag"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"os"
)

var configFile = flag.String("f", "etc/job.yaml", "Specify the config file")

func main() {
	flag.Parse()

	// 定义解析配置
	var c config.Config
	conf.MustLoad(*configFile, &c, conf.UseEnv())

	// 该操作之前是由rest.MustNewServer(c.RestConf)完成的，主要是初始化了prometheus、trace、metricsUrl等
	if err := c.SetUp(); err != nil {
		panic(err)
	}

	logx.DisableStat()

	// 实例化
	svcContext := svc.NewServiceContext(c)
	ctx := context.Background()
	// 注册处理程序
	cronJob := logic.NewCronJob(ctx, svcContext)
	mux := cronJob.Register()

	// 绑定，启动
	if err := svcContext.AsynqServer.Run(mux); err != nil {
		logx.WithContext(ctx).Errorf("!!!CronJobErr!!! run err:%+v", err)
		os.Exit(1)
	}
}
