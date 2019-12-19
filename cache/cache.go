package cache

import "github.com/lhlyu/iyu/common"

type Cache struct {
	TraceId string
}

func NewCache(traceId string) *Cache {
	return &Cache{traceId}
}

func (s *Cache) Error(err error) bool {
	if err == nil {
		return false
	}
	common.Ylog.Log(3, "error", s.TraceId, "cache", err.Error())
	return true
}

func (s *Cache) Info(param ...interface{}) {
	common.Ylog.Log(3, "info", s.TraceId, "cache", param...)
}
