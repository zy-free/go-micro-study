package loginlog

import (
	"flag"
	"fmt"
	"github.com/zy-free/micro-study/api/member/conf"
	"github.com/zy-free/micro-study/api/member/dao/loginlog"
	loginlogModel "github.com/zy-free/micro-study/api/member/model/loginlog"
)

type Service struct {
	dao *loginlog.Dao

	// rpc
}

func New() (s *Service) {
	flag.Parse()
	if err := conf.Init(); err != nil {
		panic(fmt.Sprintf("conf.Init() error(%v)", err))
	}
	s = &Service{
		dao: loginlog.New(conf.Conf),
	}

	// 可开启定时任务
	//s.initCron()

	return s
}

func (s *Service) AddLoginLog(arg *loginlogModel.ArgLogAdd) (id int64, err error) {
	arg.Phone = "1"
	return
}
