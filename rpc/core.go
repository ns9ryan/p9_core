package main

import (
	"flag"
	"fmt"

	"github.com/ns9ryan/p9_core/rpc/internal/config"
	baseserviceServer "github.com/ns9ryan/p9_core/rpc/internal/server/baseservice"
	pingserviceServer "github.com/ns9ryan/p9_core/rpc/internal/server/pingservice"
	roleserviceServer "github.com/ns9ryan/p9_core/rpc/internal/server/roleservice"
	"github.com/ns9ryan/p9_core/rpc/internal/svc"
	"github.com/ns9ryan/p9_core/rpc/pb/core"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/core.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		core.RegisterPingServiceServer(grpcServer, pingserviceServer.NewPingServiceServer(ctx))
		core.RegisterBaseServiceServer(grpcServer, baseserviceServer.NewBaseServiceServer(ctx))
		core.RegisterRoleServiceServer(grpcServer, roleserviceServer.NewRoleServiceServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
