package service

import "github.com/zy-free/micro-study/api/member/dao"

type Service struct {
	dao *dao.Dao

	// rpc
}

func New() (s *Service) {
	s = &Service{
		dao: dao.New(),
	}

	// 可开启定时任务
	//s.initCron()

	return s
}
