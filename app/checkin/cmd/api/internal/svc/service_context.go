package svc

import (
	"Luckify/app/checkin/cmd/api/internal/config"
	"Luckify/app/checkin/cmd/rpc/checkin"
	"Luckify/app/usercenter/cmd/rpc/usercenter"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config        config.Config
	CheckinRpc    checkin.Checkin
	UsercenterRpc usercenter.Usercenter
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:        c,
		CheckinRpc:    checkin.NewCheckin(zrpc.MustNewClient(c.CheckinRpcConf)),
		UsercenterRpc: usercenter.NewUsercenter(zrpc.MustNewClient(c.UsercenterRpcConf)),
	}
}
