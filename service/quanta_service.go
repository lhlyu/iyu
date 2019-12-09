package service

import (
	"github.com/lhlyu/iyu/cache"
	"github.com/lhlyu/iyu/common"
	"github.com/lhlyu/iyu/controller/vo"
	"github.com/lhlyu/iyu/errcode"
	"github.com/lhlyu/iyu/repository"
	"github.com/lhlyu/iyu/service/bo"
	"github.com/lhlyu/iyu/util"
)

type quantaService struct {
	*Service
}

func NewQuantaService(traceId string) *quantaService {
	return &quantaService{
		Service: &Service{traceId},
	}
}

func (s *quantaService) QueryPage(page *common.Page) *errcode.ErrCode {
	dao := repository.NewDao(s.TraceId)
	total := dao.QueryQuantaCount()
	page.SetTotal(total)
	datas := dao.QueryQuantaPage(page)
	if len(datas) == 0 {
		return errcode.EmptyData
	}
	return errcode.Success.WithPage(page, datas)
}

func (s *quantaService) Query(reload bool, id ...int) *errcode.ErrCode {
	c := cache.NewCache(s.TraceId)
	var values []*bo.Quanta
	if !reload {
		values = c.GetQuanta(id...)
	}
	if len(values) > 0 {
		return errcode.Success.WithData(values)
	}
	datas := repository.NewDao(s.TraceId).QueryQuanta(id...)
	if len(datas) == 0 {
		return errcode.EmptyData
	}
	for _, v := range datas {
		values = append(values, &bo.Quanta{v.Id, v.Key, v.Value, v.Desc, v.IsEnable})
	}
	go c.SetQuanta(values...)
	return errcode.Success.WithData(values)
}

// add update
func (s *quantaService) Edit(param *vo.QuantaVo) *errcode.ErrCode {
	dao := repository.NewDao(s.TraceId)
	if param.Id == 0 {
		data := dao.GetQuantaByKey(param.Key)
		if data != nil {
			return errcode.ExsistData
		}
		id, err := dao.InsertQuanta(param)
		if err != nil {
			return errcode.InsertError
		}
		go s.Query(true, id)
		return errcode.Success
	}
	data := dao.GetQuantaById(param.Id)
	if data == nil {
		return errcode.NoExsistData
	}
	other := dao.GetQuantaByKey(param.Key)
	if len(other) > 0 && other[0].Id != data.Id {
		return errcode.ExsistData
	}
	util.CompareIntSet(&data.IsEnable, &param.IsEnable)
	util.CompareStrSet(&data.Key, &param.Key)
	util.CompareStrSet(&data.Value, &param.Value)
	util.CompareStrSet(&data.Desc, &param.Desc)
	if err := dao.UpdateQuanta(data); err != nil {
		return errcode.UpdateError
	}
	go s.Query(true, data.Id)
	return errcode.Success
}

func (s *quantaService) QueryByKey(key ...string) map[string]*bo.Quanta {
	che := cache.NewCache(s.TraceId)
	datas := che.GetQuantaByKeys(key...)
	m := make(map[string]*bo.Quanta)
	if len(datas) > 0 {
		for _, v := range datas {
			m[v.Key] = v
		}
		return m
	}
	dao := repository.NewDao(s.TraceId)
	list := dao.GetQuantaByKey(key...)
	if len(list) == 0 {
		return nil
	}
	for _, v := range list {
		value := &bo.Quanta{
			Id:       v.Id,
			Key:      v.Key,
			Value:    v.Value,
			Desc:     v.Desc,
			IsEnable: v.IsEnable,
		}
		m[v.Key] = value
	}
	return m
}
