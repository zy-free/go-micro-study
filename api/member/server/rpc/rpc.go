package rpc

import (
	"github.com/zy-free/micro-study/api/member/service/member"
)

var svc *member.Service

func init() {
	svc = member.New()
}
