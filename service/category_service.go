package service

import (
	"github.com/lhlyu/iyu/cache"
	"github.com/lhlyu/iyu/controller/vo"
	"github.com/lhlyu/iyu/errcode"
	"github.com/lhlyu/iyu/repository"
	"github.com/lhlyu/iyu/service/bo"
	"github.com/lhlyu/iyu/util"
)

type categoryService struct {
}

func NewCategoryService() *categoryService {
	return &categoryService{}
}

func (s *categoryService) Query(reload bool, id ...int) *errcode.ErrCode {
	c := cache.NewCache()
	var values []*bo.Category
	if !reload {
		values = c.GetCategory(id...)
	}
	if len(values) > 0 {
		return errcode.Success.WithData(values)
	}
	datas := repository.NewDao().QueryCategory(id...)
	if len(datas) == 0 {
		return errcode.EmptyData
	}
	for _, v := range datas {
		values = append(values, &bo.Category{v.Id, v.Name, v.IsDelete})
	}
	go c.SetCategory(values...)
	return errcode.Success.WithData(values)
}

// add update
func (s *categoryService) Edit(param *vo.CategoryVo) *errcode.ErrCode {
	dao := repository.NewDao()
	if param.Id == 0 {
		data := dao.GetCategoryByName(param.Name)
		if data != nil {
			return errcode.ExsistData
		}
		id, err := dao.InsertCategory(param)
		if err != nil {
			return errcode.InsertError
		}
		go s.Query(true, id)
		return errcode.Success
	}
	data := dao.GetCategoryById(param.Id)
	if data == nil {
		return errcode.NoExsistData
	}
	other := dao.GetCategoryByName(param.Name)
	if other != nil && other.Id != data.Id {
		return errcode.ExsistData
	}
	util.CompareIntSet(&data.IsDelete, &param.IsDelete)
	util.CompareStrSet(&data.Name, &param.Name)
	if err := dao.UpdateCategory(data); err != nil {
		return errcode.UpdateError
	}
	go s.Query(true, data.Id)
	return errcode.Success
}
