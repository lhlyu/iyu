package service

import (
	"github.com/lhlyu/iyu/cache"
	"github.com/lhlyu/iyu/common"
	"github.com/lhlyu/iyu/controller/vo"
	"github.com/lhlyu/iyu/errcode"
	"github.com/lhlyu/iyu/repository"
	"github.com/lhlyu/iyu/repository/po"
	"github.com/lhlyu/iyu/service/bo"
)

type quantaService struct {
}

func NewQuantaService() *quantaService {
	return &quantaService{}
}

// get all quantas
func (*quantaService) GetAll(page *common.Page, reload bool) *errcode.ErrCode {
	c := cache.NewCache()
	var quantas []*bo.Quanta
	if !reload {
		quantas = c.GetQuantaPage(page)
	}
	if len(quantas) != 0 {
		return errcode.Success.WithPage(page, quantas)
	}
	datas := repository.NewDao().GetQuantaAll()
	total := len(datas)
	if total == 0 {
		return errcode.EmptyData
	}
	page.SetTotal(total)
	var quantaAll []*bo.Quanta
	for _, v := range datas {
		quantaAll = append(quantaAll, &bo.Quanta{v.Id, v.Key, v.Value, v.Desc, v.IsEnable})
	}
	quantas = quantaAll[page.StartRow:page.StopRow]
	go c.LoadQuantas(quantaAll...)
	return errcode.Success.WithPage(page, quantas)
}

func (s *quantaService) Insert(param *vo.QuantaVo) *errcode.ErrCode {
	dao := repository.NewDao()
	data := dao.GetQuantaByKey(0, param.Key)
	if data != nil {
		return errcode.ExsistData
	}
	p := &po.YuQuanta{
		Key:      param.Key,
		Value:    param.Value,
		Desc:     param.Desc,
		IsEnable: param.IsEnable,
	}
	if err := dao.InsertQuanta(p); err != nil {
		return errcode.InsertError
	}
	go s.GetAll(nil, true)
	return errcode.Success
}

func (s *quantaService) Update(param *vo.QuantaVo) *errcode.ErrCode {
	dao := repository.NewDao()
	data := dao.GetQuantaById(param.Id)
	if data == nil {
		return errcode.NoExsistData
	}
	other := dao.GetQuantaByKey(param.Id, param.Key)
	if other != nil {
		return errcode.ExsistData
	}
	p := &po.YuQuanta{
		Id:       param.Id,
		Key:      param.Key,
		Value:    param.Value,
		Desc:     param.Desc,
		IsEnable: param.IsEnable,
	}
	if err := dao.UpdateQuanta(p); err != nil {
		return errcode.UpdateError
	}
	go s.GetAll(nil, true)
	return errcode.Success
}

// if real == 1 then delete from database
func (s *quantaService) Delete(param *vo.QuantaVo) *errcode.ErrCode {
	dao := repository.NewDao()
	data := dao.GetQuantaById(param.Id)
	if data == nil {
		return errcode.NoExsistData
	}
	if param.Real == 1 {
		if err := dao.DeleteQuantaById(data.Id); err != nil {
			return errcode.DeleteError
		}
		go s.GetAll(nil, true)
		return errcode.Success
	}
	p := &po.YuQuanta{
		Id:       param.Id,
		Key:      data.Key,
		Value:    data.Value,
		Desc:     data.Desc,
		IsEnable: common.TWO,
	}
	if err := dao.UpdateQuanta(p); err != nil {
		return errcode.UpdateError
	}
	go s.GetAll(nil, true)
	return errcode.Success
}
