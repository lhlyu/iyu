package service

import (
	"github.com/lhlyu/iyu/cache"
	"github.com/lhlyu/iyu/common"
	"github.com/lhlyu/iyu/errcode"
	"github.com/lhlyu/iyu/repository"
	"github.com/lhlyu/iyu/service/bo"
)

type nailService struct {
}

func NewNailService() *nailService {
	return &nailService{}
}

// get all nails
func (*nailService) GetAll(reload bool) *errcode.ErrCode {
	c := cache.NewCache()
	var nails []*bo.Nail
	if !reload {
		nails = c.GetNailAll()
	}
	if len(nails) != 0 {
		return errcode.Success.WithData(nails)
	}
	datas := repository.NewDao().GetNailAll()
	if len(datas) == 0 {
		return errcode.EmptyData
	}
	for _, v := range datas {
		nails = append(nails, &bo.Nail{v.Id, v.Name, v.Color, v.IsDelete})
	}
	go c.LoadNails(nails...)
	return errcode.Success.WithData(nails)
}

func (s *nailService) Insert(name, color string) *errcode.ErrCode {
	dao := repository.NewDao()
	data := dao.GetNailByName(0, name)
	if data != nil {
		return errcode.ExsistData
	}
	if err := dao.InsertNail(name, color); err != nil {
		return errcode.InsertError
	}
	go s.GetAll(true)
	return errcode.Success
}

func (s *nailService) Update(id, status int, name, color string) *errcode.ErrCode {
	dao := repository.NewDao()
	data := dao.GetNailById(id)
	if data == nil {
		return errcode.NoExsistData
	}
	other := dao.GetNailByName(id, name)
	if other != nil {
		return errcode.ExsistData
	}
	if err := dao.UpdateNail(data.Id, status, name, color); err != nil {
		return errcode.UpdateError
	}
	go s.GetAll(true)
	return errcode.Success
}

// if real == 1 then delete from database
func (s *nailService) Delete(id, real int) *errcode.ErrCode {
	dao := repository.NewDao()
	data := dao.GetNailById(id)
	if data == nil {
		return errcode.NoExsistData
	}
	if real == 1 {
		if err := dao.DeleteNailById(data.Id); err != nil {
			return errcode.DeleteError
		}
		go s.GetAll(true)
		return errcode.Success
	}
	if err := dao.UpdateNail(data.Id, common.TWO, data.Name, data.Color); err != nil {
		return errcode.UpdateError
	}
	go s.GetAll(true)
	return errcode.Success
}
