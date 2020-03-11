package service

import (
	"fmt"
	"github.com/kataras/iris/v12"
	"github.com/lhlyu/iyu/cache"
	"github.com/lhlyu/iyu/dao"
	"github.com/lhlyu/iyu/result"
	"github.com/lhlyu/iyu/trace"
)

type IndexService struct {
	trace.BaseTracker
	*cache.IndexCache
}

func NewIndexService(ctx iris.Context) *IndexService {
	tracker := ctx.Values().Get(trace.TRACKER).(*trace.Tracker)
	return &IndexService{
		BaseTracker: trace.NewBaseTracker(tracker),
		IndexCache:  cache.NewIndexCache(tracker),
	}
}

func (s *IndexService) Hello(name string, age int) *result.R {
	s.Info("IndexService.Hello", name, age)
	// 查询缓存
	s.Get(name)

	// 查询数据库
	d := dao.NewIndexDao(s.ITracker)
	v := d.Query(name)

	s.Debug(fmt.Sprintf("%s is real age is %d", name, v))

	if age > v {
		s.Info(name, " is real age less than ", age)
		return result.Failure.WithData(fmt.Sprintf("%s is real age less than %d", name, age))
	}

	return result.Success.WithData(fmt.Sprintf("%s is real age greater than %d", name, age))
}