package service

import "github.com/zy-free/micro-study/api/member/model"

func (s *Service) CreateMember(arg *model.ArgMemberCreate) (int64, error) {
	return s.dao.CreateMember()
}
