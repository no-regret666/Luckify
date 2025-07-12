package svc

import (
	"Luckify/app/notice/cmd/rpc/internal/config"
	"Luckify/app/usercenter/cmd/rpc/usercenter"
	"github.com/hibiken/asynq"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config        config.Config
	AsynqClient   *asynq.Client
	UsercenterRpc usercenter.Usercenter
}

func newAsynqClient(c config.Config) *asynq.Client {
	return asynq.NewClient(asynq.RedisClientOpt{Addr: c.RedisForAsynq.Host, Password: c.RedisForAsynq.Pass})
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:        c,
		AsynqClient:   newAsynqClient(c),
		UsercenterRpc: usercenter.NewUsercenter(zrpc.MustNewClient(c.UsercenterRpcConf)),
	}
}
