package service

import (
	"github.com/lhlyu/iyu/cache"
	"github.com/lhlyu/iyu/controller/vo"
	"github.com/lhlyu/iyu/errcode"
	"github.com/lhlyu/iyu/repository"
	"github.com/lhlyu/iyu/service/bo"
)

type categoryService struct {
}

func NewCategoryService() *categoryService {
	return &categoryService{}
}

// get all categorys
func (*categoryService) GetAll(reload bool) *errcode.ErrCode {
	c := cache.NewCache()
	var categorys []*bo.Category
	if !reload {
		categorys = c.GetCategoryAll()
	}
	if len(categorys) != 0 {
		return errcode.Success.WithData(categorys)
	}
	datas := repository.NewDao().GetCategoryAll()
	if len(datas) == 0 {
		return errcode.EmptyData
	}
	for _, v := range datas {
		categorys = append(categorys, &bo.Category{v.Id, v.Name, v.IsDelete})
	}
	go c.LoadCategorys(categorys...)
	return errcode.Success.WithData(categorys)
}

func (s *categoryService) Insert(param *vo.CategoryVo) *errcode.ErrCode {
	dao := repository.NewDao()
	data := dao.GetCategoryByName(param.Name)
	if data != nil {
		return errcode.ExsistData
	}
	if err := dao.InsertCategory(param); err != nil {
		return errcode.InsertError
	}
	go s.GetAll(true)
	return errcode.Success
}

func (s *categoryService) Update(param *vo.CategoryVo) *errcode.ErrCode {
	dao := repository.NewDao()
	data := dao.GetCategoryById(param.Id)
	if data == nil {
		return errcode.NoExsistData
	}
	other := dao.GetCategoryByName(param.Name)
	if other != nil && other.Id != param.Id {
		return errcode.ExsistData
	}
	if err := dao.UpdateCategory(param); err != nil {
		return errcode.UpdateError
	}
	go s.GetAll(true)
	return errcode.Success
}

// if real == 1 then delete from database
func (s *categoryService) Delete(param *vo.CategoryVo) *errcode.ErrCode {
	dao := repository.NewDao()
	data := dao.GetCategoryById(param.Id)
	if data == nil {
		return errcode.NoExsistData
	}
	if param.Real == 1 {
		if err := dao.DeleteCategoryById(data.Id); err != nil {
			return errcode.DeleteError
		}
		go s.GetAll(true)
		return errcode.Success
	}
	if err := dao.UpdateCategory(param); err != nil {
		return errcode.UpdateError
	}
	go s.GetAll(true)
	return errcode.Success
}
