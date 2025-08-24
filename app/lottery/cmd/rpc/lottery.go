package main

import (
	"Luckify/app/lottery/cmd/rpc/internal/logic"
	"Luckify/common/interceptor/rpcserver"
	"flag"
	"fmt"

	"Luckify/app/lottery/cmd/rpc/internal/config"
	"Luckify/app/lottery/cmd/rpc/internal/server"
	"Luckify/app/lottery/cmd/rpc/internal/svc"
	"Luckify/app/lottery/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/lottery.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		pb.RegisterLotteryServer(grpcServer, server.NewLotteryServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	s.AddUnaryInterceptors(rpcserver.LoggerInterceptor)

	// 开启协程消费kafka消息
	go logic.HandleInstantLotteryParticipationMessage(ctx)

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
