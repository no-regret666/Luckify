package svc

import (
	"Luckify/app/checkin/cmd/rpc/checkin"
	"Luckify/app/lottery/cmd/rpc/lottery"
	"Luckify/app/mqueue/cmd/job/internal/config"
	"github.com/hibiken/asynq"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config      config.Config
	AsynqServer *asynq.Server
	LotteryRpc  lottery.LotteryZrpcClient
	CheckinRpc  checkin.Checkin
}

func newAsynqServer(c config.Config) *asynq.Server {
	return asynq.NewServer(
		asynq.RedisClientOpt{Addr: c.Redis.Host, Password: c.Redis.Pass},
		asynq.Config{
			IsFailure: func(err error) bool {
				logx.Errorf(" [asynq server] exec task IsFailure ======= >>>>>>> err: %+v", err)
				return true
			},
			Concurrency: 20, // max concurrent process job task num
		},
	)
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:      c,
		AsynqServer: newAsynqServer(c),
		LotteryRpc:  lottery.NewLotteryZrpcClient(zrpc.MustNewClient(c.LotteryRpcConf)),
		CheckinRpc:  checkin.NewCheckin(zrpc.MustNewClient(c.CheckinRpcConf)),
	}
}
