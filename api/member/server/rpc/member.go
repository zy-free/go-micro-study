package rpc

import (
	"context"
	proto "github.com/zy-free/micro-study/api/member/proto"
)

type MemberServer struct{}

func (s *MemberServer) GetByID(ctx context.Context, req *proto.GetByIDRequest, rsp *proto.MemberResponse) (err error) {
	member, err := svc.GetMember(req.Id)
	if err != nil {
		return err
	}
	rsp.Id = member.ID
	rsp.Name = member.Name
	rsp.Phone = member.Phone
	rsp.Age = int32(member.Age)
	return nil
}
