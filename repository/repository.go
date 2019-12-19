package repository

import (
	"github.com/lhlyu/iyu/common"
)

type Dao struct {
	TraceId string
}

func NewDao(traceId string) *Dao {
	return &Dao{
		TraceId: traceId,
	}
}

func (s *Dao) Error(err error) bool {
	if err == nil {
		return false
	}
	common.Ylog.Log(3, "error", s.TraceId, "repository", err.Error())
	return true
}

func (s *Dao) Info(param ...interface{}) {
	common.Ylog.Log(3, "info", s.TraceId, "repository", param...)
}
