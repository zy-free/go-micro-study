package rpc

import (
	"context"
	"github.com/zy-free/micro-study/api/member/proto"
)

type MemberServer struct{}

func (s *MemberServer) Hello(ctx context.Context, req *proto.HelloRequest, rsp *proto.HelloResponse) error {
	rsp.Say = "memberServer Hello: " + req.Name
	return nil
}
