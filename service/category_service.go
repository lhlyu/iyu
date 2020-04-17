package service

import (
	"github.com/lhlyu/iyu/cache"
	"github.com/lhlyu/iyu/common"
	"github.com/lhlyu/iyu/dao/po"
	"github.com/lhlyu/iyu/result"
	"github.com/lhlyu/iyu/service/vo"
	"github.com/lhlyu/iyu/trace"
)

type CategoryService struct {
	trace.BaseTracker
	*cache.CategoryCache
}

func NewCategoryService(tracker trace.ITracker) *CategoryService {
	return &CategoryService{
		BaseTracker:   trace.NewBaseTracker(tracker),
		CategoryCache: cache.NewCategoryCache(tracker),
	}
}

func (s *CategoryService) QueryCategory() *result.R {
	var datas []*vo.Category
	exists, err := s.CategoryCache.Get(&datas)
	if err != nil {
		return result.CacheErr
	}
	if exists {
		return result.Success.WithData(datas)
	}

	var items []*po.Category
	err = common.DB.Order("created_at").Find(&items).Error
	if err != nil {
		s.Error(err)
		return result.QueryError
	}

	for _, v := range items {
		data := &vo.Category{
			Id:    v.Id,
			Name:  v.Name,
			Count: v.Count,
		}
		datas = append(datas, data)
	}
	s.CategoryCache.Set(datas)
	return result.Success.WithData(datas)
}

func (s *CategoryService) AddCategory() *result.R {
	var items []*po.Category
	err := common.DB.Order("created_at").Find(&items).Error
	if err != nil {
		s.Error(err)
		return result.QueryError
	}
	return nil
}
