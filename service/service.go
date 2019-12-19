package service

import (
	"github.com/lhlyu/iyu/cache"
	"github.com/lhlyu/iyu/common"
	"github.com/lhlyu/iyu/repository"
)

type Service struct {
	TraceId string
	Cache   *cache.Cache
	Dao     *repository.Dao
}

func NewService(traceId string) *Service {
	return &Service{
		TraceId: traceId,
		Cache:   cache.NewCache(traceId),
		Dao:     repository.NewDao(traceId),
	}
}

func (s *Service) Error(err error) bool {
	if err == nil {
		return false
	}
	common.Ylog.Log(3, "error", s.TraceId, "service", err.Error())
	return true
}

func (s *Service) Info(param ...interface{}) {
	common.Ylog.Log(3, "info", s.TraceId, "service", param...)
}

/**
standard: 命名规范

XxxQuery:      普通条件查询
XxxQueryPage:  分页查询
XxxGet:        获取单条记录
XxxGetByField: 根据某个字段获取单条记录

XxxDelByField: 根据某个字段删除
XxxBatchDel:   批量删除

XxxUpd:        更新记录

XxxAdd:        添加
XxxBatchAdd:   批量添加
*/
