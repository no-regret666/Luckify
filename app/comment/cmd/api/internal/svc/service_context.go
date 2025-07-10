package svc

import (
	"Luckify/app/comment/cmd/api/internal/config"
	"Luckify/app/comment/cmd/rpc/comment"
	"Luckify/app/usercenter/cmd/rpc/usercenter"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config        config.Config
	CommentRpc    comment.CommentZrpcClient
	UsercenterRpc usercenter.Usercenter
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:        c,
		CommentRpc:    comment.NewCommentZrpcClient(zrpc.MustNewClient(c.CommentRpcConf)),
		UsercenterRpc: usercenter.NewUsercenter(zrpc.MustNewClient(c.UsercenterRpcConf)),
	}
}
