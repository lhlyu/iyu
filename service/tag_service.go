package service

import (
	"github.com/lhlyu/iyu/cache"
	"github.com/lhlyu/iyu/controller/vo"
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

func (s *tagService) Insert(param *vo.TagVo) *errcode.ErrCode {
	dao := repository.NewDao()
	data := dao.GetTagByName(param.Name)
	if data != nil {
		return errcode.ExsistData
	}
	if err := dao.InsertTag(param); err != nil {
		return errcode.InsertError
	}
	go s.GetAll(true)
	return errcode.Success
}

func (s *tagService) Update(param *vo.TagVo) *errcode.ErrCode {
	dao := repository.NewDao()
	data := dao.GetTagById(param.Id)
	if data == nil {
		return errcode.NoExsistData
	}
	other := dao.GetTagByName(param.Name)
	if other != nil && other.Id != param.Id {
		return errcode.ExsistData
	}
	if err := dao.UpdateTag(param); err != nil {
		return errcode.UpdateError
	}
	go s.GetAll(true)
	return errcode.Success
}

// if real == 1 then delete from database
func (s *tagService) Delete(param *vo.TagVo) *errcode.ErrCode {
	dao := repository.NewDao()
	data := dao.GetTagById(param.Id)
	if data == nil {
		return errcode.NoExsistData
	}
	if param.Real == 1 {
		if err := dao.DeleteTagById(data.Id); err != nil {
			return errcode.DeleteError
		}
		go s.GetAll(true)
		return errcode.Success
	}
	if err := dao.UpdateTag(param); err != nil {
		return errcode.UpdateError
	}
	go s.GetAll(true)
	return errcode.Success
}
