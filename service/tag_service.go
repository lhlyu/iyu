package service

import (
	"github.com/lhlyu/iyu/cache"
	"github.com/lhlyu/iyu/common"
	"github.com/lhlyu/iyu/errcode"
	"github.com/lhlyu/iyu/repository"
	"github.com/lhlyu/iyu/service/bo"
)

type tagService struct {
}

func NewTagService() *tagService {
	return &tagService{}
}

// get all tags
func (*tagService) GetAll(reload bool) *errcode.ErrCode {
	c := cache.NewCache()
	var tags []*bo.Tag
	if !reload {
		tags = c.GetTagAll()
	}
	if len(tags) != 0 {
		return errcode.Success.WithData(tags)
	}
	datas := repository.NewDao().GetTagAll()
	if len(datas) == 0 {
		return errcode.EmptyData
	}
	for _, v := range datas {
		tags = append(tags, &bo.Tag{v.Id, v.Name, v.IsDelete})
	}
	go c.LoadTags(tags...)
	return errcode.Success.WithData(tags)
}

func (s *tagService) Insert(name string) *errcode.ErrCode {
	dao := repository.NewDao()
	data := dao.GetTagByName(name)
	if data == nil {
		if err := dao.InsertTag(name); err != nil {
			return errcode.InsertError
		}
	} else {
		if data.IsDelete == common.ONE {
			return errcode.ExsistData
		}
		if err := dao.UpdateTag(data.Id, common.ONE, data.Name); err != nil {
			return errcode.InsertError.AddMsg(1)
		}
	}
	s.GetAll(true)
	return errcode.Success
}

func (s *tagService) Update(id, status int, name string) *errcode.ErrCode {
	dao := repository.NewDao()
	data := dao.GetTagById(id)
	if data == nil {
		return errcode.NoExsistData
	}
	other := dao.GetTagByName(name)
	if other != nil {
		return errcode.ExsistData
	}
	if err := dao.UpdateTag(data.Id, status, name); err != nil {
		return errcode.UpdateError
	}
	go s.GetAll(true)
	return errcode.Success
}

// if real == 1 then delete from database
func (s *tagService) Delete(id, real int) *errcode.ErrCode {
	dao := repository.NewDao()
	data := dao.GetTagById(id)
	if data == nil {
		return errcode.NoExsistData
	}
	if real == 1 {
		if err := dao.DeleteTagById(data.Id); err != nil {
			return errcode.DeleteError
		}
		s.GetAll(true)
		return errcode.Success
	}
	if err := dao.UpdateTag(data.Id, common.TWO, data.Name); err != nil {
		return errcode.UpdateError
	}
	go s.GetAll(true)
	return errcode.Success
}
