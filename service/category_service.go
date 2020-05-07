package service

import (
	"github.com/lhlyu/iyu/cache"
	"github.com/lhlyu/iyu/controller/dto"
	"github.com/lhlyu/iyu/dao"
	"github.com/lhlyu/iyu/dao/po"
	"github.com/lhlyu/iyu/result"
	"github.com/lhlyu/iyu/service/vo"
	"github.com/lhlyu/iyu/trace"
	"github.com/lhlyu/iyu/util"
	"strings"
	"time"
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

// get all categorys
func (c CategoryService) GetAll() ([]*vo.Category, *result.R) {
	items := []*vo.Category{}
	// read by redis
	if ok, _ := c.CategoryCache.Get(&items); ok {
		if len(items) == 0 {
			return nil, result.EmptyData
		}
		return items, nil
	}
	datas := []*po.Category{}
	if err := c.CategoryDao.Query(&datas, ""); err != nil {
		return nil, result.QueryError
	}
	if len(datas) == 0 {
		c.CategoryCache.Set(items)
		return nil, result.EmptyData
	}

	for _, v := range datas {
		item := &vo.Category{
			Id:    v.Id,
			Name:  v.Name,
			Count: v.Count,
		}
		items = append(items, item)
	}
	c.CategoryCache.Set(items)
	return items, nil
}

// query by name
func (c CategoryService) QueryByName(name string) ([]*vo.Category, *result.R) {
	items := []*vo.Category{}
	// read by redis
	if ok, _ := c.CategoryCache.Get(&items); ok {
		values := []*vo.Category{}
		for _, v := range items {
			if strings.Contains(v.Name, name) {
				values = append(values, v)
			}
		}
		if len(values) == 0 {
			return nil, result.EmptyData
		}
		return values, nil
	}

	datas := []*po.Category{}
	if err := c.CategoryDao.Query(&datas, "name like ?", util.Concat("%", name, "%")); err != nil {
		return nil, result.QueryError
	}
	if len(datas) == 0 {
		return nil, result.EmptyData
	}

	for _, v := range datas {
		item := &vo.Category{
			Id:    v.Id,
			Name:  v.Name,
			Count: v.Count,
		}
		items = append(items, item)
	}
	return items, nil
}

func (c CategoryService) AddOne(param *dto.Category) *result.R {
	now := time.Now()
	data := &po.Category{
		Name:      param.Name,
		CreatedAt: now,
		UpdatedAt: now,
	}
	if err := c.CategoryDao.Add(data); err != nil {
		return result.InsertError.WithMsg(err.Error())
	}
	c.CategoryCache.Clear()
	return result.Success
}

func (c CategoryService) UpdateOne(param *dto.Category) *result.R {
	now := time.Now()
	data := &po.Category{
		Id:        param.Id,
		Name:      param.Name,
		UpdatedAt: now,
	}
	if err := c.CategoryDao.Update(data, "id = ?", param.Id); err != nil {
		return result.UpdateError
	}
	c.CategoryCache.Clear()
	return result.Success
}

func (c CategoryService) DelById(param *dto.Category) *result.R {
	if err := c.CategoryDao.Del(param.Id); err != nil {
		return result.DeleteError.WithMsg(err.Error())
	}
	c.CategoryCache.Clear()
	return result.Success
}
