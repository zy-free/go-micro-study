package service

import (
	"flag"
	"fmt"
	"github.com/zy-free/micro-study/api/member/conf"
	"github.com/zy-free/micro-study/api/member/dao"
)

type Service struct {
	dao *dao.Dao

	// rpc
}

func New() (s *Service) {
	flag.Parse()
	if err := conf.Init(); err != nil {
		panic(fmt.Sprintf("conf.Init() error(%v)",err))
	}
	s = &Service{
		dao: dao.New(conf.Conf),
	}

	// 可开启定时任务
	//s.initCron()

	return s
}
