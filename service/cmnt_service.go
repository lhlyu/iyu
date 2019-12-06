package service

import (
	"github.com/lhlyu/iyu/common"
	"github.com/lhlyu/iyu/controller/vo"
	"github.com/lhlyu/iyu/errcode"
	"github.com/lhlyu/iyu/repository"
	"github.com/lhlyu/iyu/service/bo"
	"github.com/lhlyu/iyu/util"
)

type cmntService struct {
}

func NewCmntService() *cmntService {
	return &cmntService{}
}

func (*cmntService) Insert(param *vo.CmntVo) *errcode.ErrCode {
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
	if err := dao.InsertCmnt(param); err != nil {
		return errcode.InsertError
	}
	return errcode.Success
}

func (*cmntService) Update(param *vo.CmntVo) *errcode.ErrCode {
	dao := repository.NewDao()
	data, err := dao.GetCmnt(param.Id)
	if err != nil {
		return errcode.QueryError
	}
	if data == nil {
		return errcode.NoExsistData
	}
	util.CompareStrSet(&data.Content, &param.Content)
	util.CompareIntSet(&data.IsCheck, &param.IsCheck)
	util.CompareIntSet(&data.IsDelete, &param.IsDelete)
	if err := dao.UpdateCmnt(data); err != nil {
		return errcode.InsertError
	}
	return errcode.Success
}

func (*cmntService) QueryPage(param *vo.CmntVo) *errcode.ErrCode {
	dao := repository.NewDao()
	total, err := dao.QueryCmntCount(param)
	if err != nil {
		return errcode.QueryError
	}
	page := common.NewPage(param.PageNum, param.PageSize)
	page.SetTotal(total)
	datas, err := dao.QueryCmntPage(page, param)
	if err != nil {
		return errcode.QueryError
	}
	var userIds []int
	for _, v := range datas {
		userIds = append(userIds, v.UserId)
	}
	svc := NewUserService()
	result := svc.Query(false, userIds...)
	m := make(map[int]*bo.User)
	if result.IsSuccess() {
		for _, v := range result.Data.([]*bo.User) {
			m[v.Id] = v
		}
	}
	var values []*bo.Cmnt
	for _, v := range datas {
		updatedAt := v.UpdatedAt.Unix()
		if updatedAt < 0 {
			updatedAt = v.CreatedAt.Unix()
		}
		value := &bo.Cmnt{
			Id:        v.Id,
			UserId:    v.UserId,
			ArticleId: v.ArticleId,
			Floor:     v.Floor,
			Content:   v.Content,
			IsCheck:   v.IsCheck,
			IsDelete:  v.IsDelete,
			CreatedAt: int(v.CreatedAt.Unix()),
			UpdatedAt: int(updatedAt),
		}
		if u, ok := m[v.UserId]; ok {
			value.UserName = u.UserName
			value.UserAvatar = u.AvatarUrl
		}
		values = append(values, value)
	}
	return errcode.Success.WithPage(page, values)
}

func (*cmntService) queryPost(cmntIds ...int) map[int]*bo.PostData {
	return nil
}
