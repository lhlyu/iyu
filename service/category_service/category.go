package category_service

import (
	"github.com/lhlyu/iyu/cache"
	"github.com/lhlyu/iyu/common"
	"github.com/lhlyu/iyu/controller/dto"
	"github.com/lhlyu/iyu/errcode"
	"github.com/lhlyu/iyu/repository/category_repository"
	"github.com/lhlyu/iyu/repository/po"
	"github.com/lhlyu/iyu/service/vo"
	"github.com/lhlyu/yutil"
)

type Service struct {
	common.BaseService
	dao *category_repository.Dao
	che *cache.Cache
}

func NewService(traceId string) *Service {
	svc := &Service{}
	svc.dao = category_repository.NewDao(traceId)
	svc.che = cache.NewCache(traceId)
	svc.SetTraceId(traceId)
	return svc
}

func (s *Service) QueryCategoryPage(param *dto.CategoryDto) *errcode.ErrCode {
	page := param.Page
	whr := &po.YuCategory{
		Id:       param.Id,
		Name:     param.Name,
		IsDelete: param.IsDelete,
	}
	total := s.dao.Count(whr)
	page.SetTotal(total)
	if total == 0 {
		return errcode.EmptyData
	}
	datas := s.dao.QueryPage(whr, page)
	items := _CategoryPosToCategoryVos(datas)
	return errcode.Success.WithPage(page, items)
}

func (s *Service) UpdateCategory(param *dto.CategoryDto) *errcode.ErrCode {
	whr := &po.YuCategory{
		Id: param.Id,
	}
	data := s.dao.Get(whr)
	if data == nil {
		return errcode.NoExsistData
	}
	whr = &po.YuCategory{
		Name: param.Name,
	}
	count := s.dao.Count(whr)
	if count > 0 {
		return errcode.ExsistData
	}
	data.IsDelete = yutil.CompareZInt(data.IsDelete, param.IsDelete)
	data.Color = yutil.CompareStr(data.Color, param.Color)
	data.Name = yutil.CompareStr(data.Name, param.Name)
	ok := s.dao.Update(data)
	if !ok {
		return errcode.UpdateError
	}
	go s.LoadCategory(data.Id)
	return errcode.Success
}

func (s *Service) LoadCategory(id int) *errcode.ErrCode {
	page := common.NewPageAll()
	whr := &po.YuCategory{}
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
		items := _CategoryPosToCategoryVos(datas)
		s.che.SetCategory(items)
	}
	return errcode.Success
}

func (s *Service) QueryCategoryByKeys(ids ...int) []*vo.CategoryVo {
	items := s.che.GetCategory(ids...)
	if len(items) == 0 {
		datas := s.dao.GetByIds(ids...)
		items = _CategoryPosToCategoryVos(datas)
	}
	if len(items) > 0 {
		go s.che.SetCategory(items)
	}
	return items
}

func _CategoryPosToCategoryVos(datas []*po.YuCategory) []*vo.CategoryVo {
	var items []*vo.CategoryVo
	for _, v := range datas {
		item := &vo.CategoryVo{
			Id:        v.Id,
			Name:      v.Name,
			Color:     v.Color,
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
