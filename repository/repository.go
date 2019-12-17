package repository

import (
    "github.com/lhlyu/iyu/common"
)

type dao struct {
	TraceId string
}

func NewDao(traceId string) *dao {
	return &dao{
		TraceId: traceId,
	}
}

func (s *dao) Error(param ...interface{}) {
	common.Ylog.Log(3, "error", s.TraceId, "repository", param...)
}

func (s *dao) Info(param ...interface{}) {
	common.Ylog.Log(3, "info", s.TraceId, "repository", param...)
}
