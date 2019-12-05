package service

import (
	"github.com/lhlyu/iyu/cache"
	"github.com/lhlyu/iyu/controller/vo"
	"github.com/lhlyu/iyu/errcode"
	"github.com/lhlyu/iyu/repository"
	"github.com/lhlyu/iyu/service/bo"
	"github.com/lhlyu/iyu/util"
)

type tagService struct {
}

func NewTagService() *tagService {
	return &tagService{}
}

func (s *tagService) Query(reload bool, id ...int) *errcode.ErrCode {
	c := cache.NewCache()
	var values []*bo.Tag
	if !reload {
		values = c.GetTag(id...)
	}
	if len(values) > 0 {
		return errcode.Success.WithData(values)
	}
	datas := repository.NewDao().QueryTag(id...)
	if len(datas) == 0 {
		return errcode.EmptyData
	}
	for _, v := range datas {
		values = append(values, &bo.Tag{v.Id, v.Name, v.IsDelete})
	}
	go c.SetTag(values...)
	return errcode.Success.WithData(values)
}

// add update
func (s *tagService) Edit(param *vo.TagVo) *errcode.ErrCode {
	dao := repository.NewDao()
	if param.Id == 0 {
		data := dao.GetTagByName(param.Name)
		if data != nil {
			return errcode.ExsistData
		}
		id, err := dao.InsertTag(param)
		if err != nil {
			return errcode.InsertError
		}
		go s.Query(true, id)
		return errcode.Success
	}
	data := dao.GetTagById(param.Id)
	if data == nil {
		return errcode.NoExsistData
	}
	other := dao.GetTagByName(param.Name)
	if other != nil && other.Id != data.Id {
		return errcode.ExsistData
	}
	util.CompareIntSet(&data.IsDelete, &param.IsDelete)
	util.CompareStrSet(&data.Name, &param.Name)
	if err := dao.UpdateTag(data); err != nil {
		return errcode.UpdateError
	}
	go s.Query(true, data.Id)
	return errcode.Success
}
