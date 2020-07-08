package member

import (
	"context"
	micro "github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
)

type rpcHandler struct {
	client MemberService
}

func New() rpcHandler {
	reg := etcd.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{
			"49.235.139.144:2379",
		}
	})
	// Create a new service
	service := micro.NewService(
		micro.Name("member.client"),
		micro.Registry(reg),
	)
	// Initialise the client and parse command line flags
	service.Init()

	// Create new user client
	member := NewMemberService("member", service.Client())

	return rpcHandler{client: member}
}

func (handler *rpcHandler) GetByID(id int64) (member *MemberResponse, err error) {
	member, err = handler.client.GetByID(context.TODO(), &GetByIDRequest{Id: id})
	return
}
