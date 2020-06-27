package service

import (
	"github.com/zy-free/micro-study/api/member/model"
	"github.com/zy-free/micro-study/lib/ecode"
)

func (s *Service) AddMember(arg *model.ArgMemberAdd) (int64, error) {
	id, err := s.dao.AddMember(arg)
	if err != nil {
		return 0, ecode.ChannelNotExist
	}
	return id, nil
}
