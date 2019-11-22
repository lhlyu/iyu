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
func (*tagService) GetTagAll() *errcode.ErrCode {
	// from cache
	c := cache.NewCache()
	tags := c.GetTagAll()
	if len(tags) == 0 {
		// from database
		datas := repository.NewDao().GetTagAll()
		if len(datas) == 0 {
			return errcode.EmptyData
		}
		for _, v := range datas {
			tags = append(tags, &bo.Tag{v.Id, v.Name, v.IsDelete})
		}
	}
	// load to cache
	c.LoadTags(tags)
	return errcode.Success.WithData(tags)
}

func (*tagService) InsertTag(name string) *errcode.ErrCode {
	// if exists
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
	return errcode.Success
}

func (*tagService) UpdateTag(id, status int, name string) *errcode.ErrCode {
	// if exists
	dao := repository.NewDao()
	data := dao.GetTagByName(name)
	if data == nil {
		if err := dao.UpdateTag(id, status, name); err != nil {
			return errcode.UpdateError
		}
	} else {
		return errcode.ExsistData
	}
	return errcode.Success
}
