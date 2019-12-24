package quanta_service

import (
	"github.com/lhlyu/iyu/cache"
	"github.com/lhlyu/iyu/common"
	"github.com/lhlyu/iyu/controller/dto"
	"github.com/lhlyu/iyu/errcode"
	"github.com/lhlyu/iyu/repository/po"
	"github.com/lhlyu/iyu/repository/quanta_repository"
	"github.com/lhlyu/iyu/service/vo"
	"github.com/lhlyu/yutil"
)

type Service struct {
	common.BaseService
	dao *quanta_repository.Dao
	che *cache.Cache
}

func NewService(traceId string) *Service {
	svc := &Service{}
	svc.dao = quanta_repository.NewDao(traceId)
	svc.che = cache.NewCache(traceId)
	svc.SetTraceId(traceId)
	return svc
}

func (s *Service) QueryQuantaPage(param *dto.QuantaDto) *errcode.ErrCode {
	page := param.Page
	whr := &po.YuQuanta{
		Id:       param.Id,
		Key:      param.Key,
		IsEnable: param.IsEnable,
	}
	total := s.dao.Count(whr)
	page.SetTotal(total)
	if total == 0 {
		return errcode.EmptyData
	}
	datas := s.dao.QueryPage(whr, page)
	items := _QuantaPosToQuantaVos(datas)
	return errcode.Success.WithPage(page, items)
}

func (s *Service) UpdateQuanta(param *dto.QuantaDto) *errcode.ErrCode {
	whr := &po.YuQuanta{
		Id: param.Id,
	}
	data := s.dao.Get(whr)
	if data == nil {
		return errcode.NoExsistData
	}
	data.IsEnable = yutil.CompareZInt(data.IsEnable, param.IsEnable)
	data.Value = yutil.CompareStr(data.Value, param.Value)
	ok := s.dao.Update(data)
	if !ok {
		return errcode.UpdateError
	}
	go s.LoadQuanta(data.Id)
	return errcode.Success
}

func (s *Service) LoadQuanta(id int) *errcode.ErrCode {
	page := common.NewPageAll()
	whr := &po.YuQuanta{}
	if id > 0 {
		whr.Id = id
	}
	total := s.dao.Count(whr)
	if total == 0 {
		return errcode.EmptyData
	}
	page.SetTotal(total)
	for page.Next() {
		datas := s.dao.QueryPage(whr, page)
		items := _QuantaPosToQuantaVos(datas)
		s.che.SetQuanta(items)
	}
	return errcode.Success
}

func (s *Service) QueryQuantaByKeys(keys ...string) []*vo.QuantaVo {
	items := s.che.GetQuanta(keys...)
	if len(items) == 0 {
		datas := s.dao.QueryByKeys(keys...)
		items = _QuantaPosToQuantaVos(datas)
	}
	if len(items) > 0 {
		go s.che.SetQuanta(items)
	}
	return items
}

func (s *Service) GetQuantaByKey(key string) *vo.QuantaVo {
	items := s.che.GetQuanta(key)
	if len(items) == 0 {
		return nil
	}
	return items[0]
}

func _QuantaPosToQuantaVos(datas []*po.YuQuanta) []*vo.QuantaVo {
	var items []*vo.QuantaVo
	for _, v := range datas {
		item := &vo.QuantaVo{
			Id:        v.Id,
			Key:       v.Key,
			Value:     v.Value,
			Desc:      v.Desc,
			IsEnable:  v.IsEnable,
			CreatedAt: v.CreatedAt.Unix(),
			UpdatedAt: v.UpdatedAt.Unix(),
		}
		if item.UpdatedAt <= 0 {
			item.UpdatedAt = item.CreatedAt
		}
		items = append(items, item)
	}
	return items
}
