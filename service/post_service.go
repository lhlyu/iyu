package service

import (
	"github.com/lhlyu/iyu/common"
	"github.com/lhlyu/iyu/controller/vo"
	"github.com/lhlyu/iyu/errcode"
	"github.com/lhlyu/iyu/repository"
	"github.com/lhlyu/iyu/util"
)

type postService struct {
}

func NewPostService() *postService {
	return &postService{}
}

func (*postService) Insert(param *vo.PostVo) *errcode.ErrCode {
	qs := NewQuantaService()
	m := qs.QueryByKey(common.KEY_2, common.KEY_3)
	param.IsCheck = common.TWO
	if len(m) > 0 {
		if v, ok := m[common.KEY_2]; ok {
			if v.IsEnable == common.TWO {
				return errcode.NoOpenCmntError
			}
		}
		if v, ok := m[common.KEY_3]; ok {
			if v.IsEnable == common.ONE {
				param.IsCheck = common.ONE
			}
		}
	}
	dao := repository.NewDao()
	if err := dao.InsertPost(param); err != nil {
		return errcode.InsertError
	}
	return errcode.Success
}

func (*postService) Update(param *vo.PostVo) *errcode.ErrCode {
	dao := repository.NewDao()
	data, err := dao.GetPost(param.Id)
	if err != nil {
		return errcode.QueryError
	}
	if data == nil {
		return errcode.NoExsistData
	}
	util.CompareStrSet(&data.Content, &param.Content)
	util.CompareIntSet(&data.IsCheck, &param.IsCheck)
	util.CompareIntSet(&data.IsDelete, &param.IsDelete)
	if err := dao.UpdatePost(data); err != nil {
		return errcode.InsertError
	}
	return errcode.Success
}
