package service

import (
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/zy-free/micro-study/api/member/model"
	"github.com/zy-free/micro-study/lib/ecode"
)

func (s *Service) AddMember(arg *model.ArgMemberAdd) (id int64, err error) {
	id, err = s.dao.AddMember(arg)
	if err != nil {
		return 0, ecode.MemberAdd
	}
	return
}

func (s *Service) GetMember(id int64) (member *model.Member, err error) {
	member, err = s.dao.GetMemberByID(id)
	if err != nil {
		logrus.Error(errors.Wrap(err, "GetMember"))
		return nil, ecode.MemberGet
	}
	return
}

func (s *Service) DeleteMember(id int64) (err error) {
	err = s.dao.DeleteMember(id)
	if err != nil {
		logrus.Error(errors.Wrap(err, "DelMember"))
		return ecode.MemberDelete
	}
	return
}

func (s *Service) SetMember(arg *model.ArgMemberSet) (err error) {
	err = s.dao.SetMember(arg)
	if err != nil {
		logrus.Error(errors.Wrap(err, "SetMember"))
		return ecode.MemberUpdate
	}
	return
}

func (s *Service) UpdateMember(arg *model.ArgMemberUpdate) (err error) {
	err = s.dao.UpdateMember(arg)
	if err != nil {
		logrus.Error(errors.Wrap(err, "UpdateMember"))
		return ecode.MemberUpdate
	}
	return
}
