package service

import (
	"github.com/lhlyu/iyu/cache"
	"github.com/lhlyu/iyu/controller/vo"
	"github.com/lhlyu/iyu/errcode"
	"github.com/lhlyu/iyu/repository"
	"github.com/lhlyu/iyu/service/bo"
	"github.com/lhlyu/iyu/util"
)

type nailService struct {
	*Service
}

func NewNailService(traceId string) *nailService {
	return &nailService{
		Service: &Service{traceId},
	}
}

func (s *nailService) Query(reload bool, id ...int) *errcode.ErrCode {
	c := cache.NewCache(s.TraceId)
	var values []*bo.Nail
	if !reload {
		values = c.GetNail(id...)
	}
	if len(values) > 0 {
		return errcode.Success.WithData(values)
	}
	datas := repository.NewDao(s.TraceId).QueryNail(id...)
	if len(datas) == 0 {
		return errcode.EmptyData
	}
	for _, v := range datas {
		values = append(values, &bo.Nail{v.Id, v.Name, v.Color, v.IsDelete})
	}
	go c.SetNail(values...)
	return errcode.Success.WithData(values)
}

// add update
func (s *nailService) Edit(param *vo.NailVo) *errcode.ErrCode {
	dao := repository.NewDao(s.TraceId)
	if param.Id == 0 {
		data := dao.GetNailByName(param.Name)
		if data != nil {
			return errcode.ExsistData
		}
		id, err := dao.InsertNail(param)
		if err != nil {
			return errcode.InsertError
		}
		go s.Query(true, id)
		return errcode.Success
	}
	data := dao.GetNailById(param.Id)
	if data == nil {
		return errcode.NoExsistData
	}
	other := dao.GetNailByName(param.Name)
	if other != nil && other.Id != data.Id {
		return errcode.ExsistData
	}
	util.CompareIntSet(&data.IsDelete, &param.IsDelete)
	util.CompareStrSet(&data.Name, &param.Name)
	util.CompareStrSet(&data.Color, &param.Color)
	if err := dao.UpdateNail(data); err != nil {
		return errcode.UpdateError
	}
	go s.Query(true, data.Id)
	return errcode.Success
}
