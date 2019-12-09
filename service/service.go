package service

import "github.com/lhlyu/iyu/common"

type Service struct {
	TraceId string
}

func (s *Service) Error(param ...interface{}) {
	common.Ylog.Log(3, "error", s.TraceId, "service", param...)
}

func (s *Service) Info(param ...interface{}) {
	common.Ylog.Log(3, "info", s.TraceId, "service", param...)
}
