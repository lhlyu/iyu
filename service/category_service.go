package service

import (
	"github.com/lhlyu/iyu/cache"
	"github.com/lhlyu/iyu/dao"
	"github.com/lhlyu/iyu/trace"
)

type CategoryService struct {
	BaseService
	dao.CategoryDao
	cache.CategoryCache
}

func NewCategoryService(tracker trace.ITracker) CategoryService {
	return CategoryService{
		BaseService:   NewBaseService(tracker),
		CategoryDao:   dao.NewCategoryDao(tracker),
		CategoryCache: cache.NewCategoryCache(tracker),
	}
}

func (c CategoryService) GetOne() {
}