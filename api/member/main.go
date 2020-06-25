package main

import (
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/web"
	"github.com/zy-free/micro-study/api/member/proto"
	"github.com/zy-free/micro-study/api/member/server/rpc"

	"github.com/zy-free/micro-study/api/member/server/http"
)

func main() {
	// Create webService
	webService := web.NewService(
		web.Address(":8080"),
		web.Handler(http.InitRouter()),
	)
	webService.Init()
	// Run webService
	webService.Run()

	// Create rpcService
	rpcService := micro.NewService(
		micro.Name("member"),
	)
	rpcService.Init()
	// Register handler
	proto.RegisterMemberHandler(rpcService.Server(), new(rpc.MemberServer))
	// Run the server
	if err := rpcService.Run(); err != nil {
		panic(err)
	}
}
