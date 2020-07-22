package member

import (
	"flag"
	"fmt"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/zy-free/micro-study/api/member/conf"
	"github.com/zy-free/micro-study/api/member/dao/loginlog"
	"github.com/zy-free/micro-study/api/member/dao/member"
	memberModel "github.com/zy-free/micro-study/api/member/model/member"
	"github.com/zy-free/micro-study/lib/ecode"
)

type Service struct {
	memberDao *member.Dao
	logDao    *loginlog.Dao

	// rpc
}

func New() (s *Service) {
	flag.Parse()
	if err := conf.Init(); err != nil {
		panic(fmt.Sprintf("conf.Init() error(%v)", err))
	}
	s = &Service{
		memberDao: member.New(conf.Conf),
		logDao:    loginlog.New(conf.Conf),
	}

	// 可开启定时任务
	//s.initCron()

	return s
}

func (s *Service) AddMember(arg *memberModel.ArgMemberAdd) (id int64, err error) {
	id, err = s.memberDao.AddMember(arg)
	if err != nil {
		return 0, ecode.MemberAdd
	}
	return
}

func (s *Service) GetMember(id int64) (member *memberModel.Member, err error) {
	member, err = s.memberDao.GetMemberByID(id)
	if err != nil {
		logrus.Error(errors.Wrap(err, "GetMember"))
		return nil, ecode.MemberGet
	}
	return
}

func (s *Service) DeleteMember(id int64) (err error) {
	err = s.memberDao.DeleteMember(id)
	if err != nil {
		logrus.Error(errors.Wrap(err, "DelMember"))
		return ecode.MemberDelete
	}
	return
}

func (s *Service) SetMember(arg *memberModel.ArgMemberSet) (err error) {
	err = s.memberDao.SetMember(arg)
	if err != nil {
		logrus.Error(errors.Wrap(err, "SetMember"))
		return ecode.MemberUpdate
	}
	return
}

func (s *Service) UpdateMember(arg *memberModel.ArgMemberUpdate) (err error) {
	err = s.memberDao.UpdateMember(arg)
	if err != nil {
		logrus.Error(errors.Wrap(err, "UpdateMember"))
		return ecode.MemberUpdate
	}
	return
}
