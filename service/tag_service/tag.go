package tag_service

import (
	"github.com/lhlyu/iyu/cache"
	"github.com/lhlyu/iyu/common"
	"github.com/lhlyu/iyu/controller/dto"
	"github.com/lhlyu/iyu/errcode"
	"github.com/lhlyu/iyu/repository/po"
	"github.com/lhlyu/iyu/repository/tag_repository"
	"github.com/lhlyu/iyu/service/vo"
	"github.com/lhlyu/yutil"
)

type Service struct {
	common.BaseService
	dao *tag_repository.Dao
	che *cache.Cache
}

func NewService(tracker *common.Tracker) *Service {
	svc := &Service{}
	svc.dao = tag_repository.NewDao(tracker)
	svc.che = cache.NewCache(tracker)
	svc.SetTracker(tracker)
	return svc
}

func (s *Service) QueryTagPage(param *dto.TagDto) *errcode.ErrCode {
	page := common.NewPage(param.PageNum, param.PageSize)
	whr := &po.YuTag{
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
	items := _TagPosToTagVos(datas)
	return errcode.Success.WithPage(page, items)
}

func (s *Service) UpdateTag(param *dto.TagDto) *errcode.ErrCode {
	whr := &po.YuTag{
		Id: param.Id,
	}
	data := s.dao.Get(whr)
	if data == nil {
		return errcode.NoExsistData
	}
	whr = &po.YuTag{
		Name: param.Name,
	}
	count := s.dao.Count(whr)
	if count > 0 {
		return errcode.ExsistData
	}
	data.IsDelete = yutil.CompareZInt(data.IsDelete, param.IsDelete)
	data.Name = yutil.CompareStr(data.Name, param.Name)
	ok := s.dao.Update(data)
	if !ok {
		return errcode.UpdateError
	}
	go s.LoadTag(data.Id)
	return errcode.Success
}

func (s *Service) LoadTag(id int) *errcode.ErrCode {
	page := common.NewPageAll()
	whr := &po.YuTag{}
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
		items := _TagPosToTagVos(datas)
		s.che.SetTag(items)
	}
	return errcode.Success
}

func (s *Service) QueryTagByIds(ids ...int) []*vo.TagVo {
	items := s.che.GetTag(ids...)
	if len(items) == 0 {
		datas := s.dao.GetByIds(ids...)
		items = _TagPosToTagVos(datas)
	}
	if len(items) > 0 {
		go s.che.SetTag(items)
	}
	return items
}

func _TagPosToTagVos(datas []*po.YuTag) []*vo.TagVo {
	var items []*vo.TagVo
	for _, v := range datas {
		item := &vo.TagVo{
			Id:        v.Id,
			Name:      v.Name,
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
