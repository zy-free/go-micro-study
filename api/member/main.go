package main

import (
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
	"github.com/micro/go-micro/v2/web"
	proto "github.com/zy-free/micro-study/api/member/proto"

	"github.com/zy-free/micro-study/api/member/server/http"
	"github.com/zy-free/micro-study/api/member/server/rpc"
	"github.com/zy-free/micro-study/lib/log"
)

func main() {
	log.Init()
	reg := etcd.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{
			"49.235.139.144:2379",
		}
	})
	// Create webService
	webService := web.NewService(
		web.Address(":8080"),
		web.Handler(http.InitRouter()),
	)
	webService.Init()
	// Run webService
	go webService.Run()

	// Create rpcService
	rpcService := micro.NewService(
		micro.Name("member"),
		micro.Registry(reg),
	)
	rpcService.Init()
	// Register handler
	proto.RegisterMemberHandler(rpcService.Server(), new(rpc.MemberServer))
	// Run the server
	if err := rpcService.Run(); err != nil {
		panic(err)
	}
}
