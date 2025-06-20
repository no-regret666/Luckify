package svc

import (
	"Luckify/app/upload/cmd/api/internal/config"
	"Luckify/app/upload/cmd/rpc/upload"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config    config.Config
	UploadRpc upload.Upload
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:    c,
		UploadRpc: upload.NewUpload(zrpc.MustNewClient(c.UploadRpcConf)),
	}
}
