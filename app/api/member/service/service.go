package service

type Service struct {
	//dao  //
}
func New() (s *Service) {
	s = &Service{

	}

	// 可开启定时任务
	//s.initCron()

	return s
}
