package service

import (
	"github.com/lhlyu/iyu/cache"
	"github.com/lhlyu/iyu/common"
	"github.com/lhlyu/iyu/controller/vo"
	"github.com/lhlyu/iyu/errcode"
	"github.com/lhlyu/iyu/repository"
	"github.com/lhlyu/iyu/service/bo"
	"github.com/lhlyu/iyu/util"
)

type userService struct {
	*Service
}

func NewUserService(traceId string) *userService {
	return &userService{
		Service: &Service{traceId},
	}
}

func (s *userService) QueryPage(param *vo.UserParam) *errcode.ErrCode {
	dao := repository.NewDao(s.TraceId)
	total, err := dao.GetUsersCount(param)
	if err != nil {
		return errcode.QueryError
	}
	if total == 0 {
		return errcode.EmptyData
	}
	page := common.NewPage(param.PageNum, param.PageSize)
	page.SetTotal(total)
	datas, err := dao.QueryUserPage(param, page)
	if err != nil {
		return errcode.QueryError
	}
	result := s.Query(false, datas...)
	if result.IsSuccess() {
		return errcode.Success.WithPage(page, result.Data)
	}
	return result
}

func (s *userService) Get(id int) *errcode.ErrCode {
	return s.Query(false, id)
}

func (s *userService) Query(reload bool, id ...int) *errcode.ErrCode {
	c := cache.NewCache(s.TraceId)
	var values []*bo.User
	if !reload {
		values = c.GetUser(id...)
	}
	if len(values) > 0 {
		return errcode.Success.WithData(values)
	}
	datas := repository.NewDao(s.TraceId).QueryUser(id...)
	if len(datas) == 0 {
		return errcode.EmptyData
	}
	for _, v := range datas {
		updatedAt := v.UpdatedAt.Unix()
		if updatedAt < 0 {
			updatedAt = v.CreatedAt.Unix()
		}
		values = append(values, &bo.User{
			Id:        v.Id,
			Ip:        v.Ip,
			Bio:       v.Bio,
			Role:      v.Role,
			ThirdId:   v.ThirdId,
			From:      v.From,
			Status:    v.Status,
			AvatarUrl: v.AvatarUrl,
			UserUrl:   v.UserUrl,
			UserName:  v.UserName,
			UpdatedAt: int(updatedAt),
			CreatedAt: int(v.CreatedAt.Unix()),
		})
	}
	go c.SetUser(values...)
	return errcode.Success.WithData(values)
}

// add update
func (s *userService) Edit(param *vo.UserEditParam) *errcode.ErrCode {
	dao := repository.NewDao(s.TraceId)
	if param.Id == 0 {
		id, err := dao.InsertUser(param)
		if err != nil {
			return errcode.InsertError
		}
		go s.Query(true, id)
		return errcode.Success
	}
	data, err := dao.GetUserById(param.Id)
	if err != nil {
		return errcode.QueryError
	}
	if data == nil {
		return errcode.NoExsistData
	}
	util.CompareStrSet(&data.UserName, &param.UserName)
	util.CompareStrSet(&data.AvatarUrl, &param.AvatarUrl)
	util.CompareStrSet(&data.Ip, &param.Ip)
	util.CompareStrSet(&data.Bio, &param.Bio)
	util.CompareIntSet(&data.Status, &param.Status)
	util.CompareIntSet(&data.Role, &param.Role)
	if err := dao.UpdateUser(data); err != nil {
		return errcode.UpdateError
	}
	go s.Query(true, data.Id)
	return errcode.Success
}
