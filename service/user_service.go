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
}

func NewUserService() *userService {
	return &userService{}
}

func (*userService) LoadUsers(id int) {
	dao := repository.NewDao()
	param := &vo.UserParam{
		PageSize: 100,
		PageNum:  1,
		Id:       id,
	}
	total, err := dao.GetUsersCount(param)
	if err != nil {
		return
	}
	if total == 0 {
		return
	}
	page := common.NewPage(param.PageNum, param.PageSize)
	page.SetTotal(total)
	for i := 0; i < page.PageMax; i++ {
		param.PageNum = i
		datas, err := dao.QueryUser(param, page)
		if err != nil {
			return
		}
		var result []*bo.UserData
		for _, v := range datas {
			updatedAt := v.UpdatedAt.Unix()
			if updatedAt < 0 {
				updatedAt = v.CreatedAt.Unix()
			}
			user := &bo.UserData{
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
			}
			result = append(result, user)
		}
		cache.NewCache().LoadUser(result)
	}
	return
}

func (s *userService) GetById(id int) *errcode.ErrCode {
	che := cache.NewCache()
	users := che.GetUsers(id)
	if len(users) > 0 {
		return errcode.Success.WithData(users[0])
	}
	dao := repository.NewDao()
	user, err := dao.GetUserById(id)
	if err != nil {
		return errcode.QueryError
	}
	if user == nil {
		return errcode.NoExsistData
	}
	updatedAt := user.UpdatedAt.Unix()
	if updatedAt < 0 {
		updatedAt = user.CreatedAt.Unix()
	}
	result := &bo.UserData{
		Id:        user.Id,
		Ip:        user.Ip,
		Bio:       user.Bio,
		Role:      user.Role,
		ThirdId:   user.ThirdId,
		From:      user.From,
		Status:    user.Status,
		AvatarUrl: user.AvatarUrl,
		UserUrl:   user.UserUrl,
		UserName:  user.UserName,
		UpdatedAt: int(updatedAt),
		CreatedAt: int(user.CreatedAt.Unix()),
	}
	che.LoadUser([]*bo.UserData{result})
	return errcode.Success.WithData(result)
}

func (s *userService) Update(param *vo.UserEditParam) *errcode.ErrCode {
	dao := repository.NewDao()
	user, err := dao.GetUserById(param.Id)
	if err != nil {
		return errcode.UpdateError
	}
	if user == nil {
		return errcode.NoExsistData
	}
	util.CompareStrSet(&user.UserName, &param.UserName)
	util.CompareStrSet(&user.AvatarUrl, &param.AvatarUrl)
	util.CompareStrSet(&user.Ip, &param.Ip)
	util.CompareStrSet(&user.Bio, &param.Bio)
	util.CompareIntSet(&user.Status, &param.Status)
	util.CompareIntSet(&user.Role, &param.Role)
	err = dao.UpdateUser(user)
	if err != nil {
		return errcode.UpdateError
	}
	s.LoadUsers(user.Id)
	return errcode.Success
}

func (s *userService) Insert(param *vo.UserEditParam) *errcode.ErrCode {
	dao := repository.NewDao()
	id, err := dao.InsertUser(param)
	if err != nil {
		return errcode.UpdateError
	}
	s.LoadUsers(id)
	return errcode.Success
}

func (s *userService) Query(param *vo.UserParam) *errcode.ErrCode {
	dao := repository.NewDao()
	total, err := dao.GetUsersCount(param)
	if err != nil {
		return errcode.QueryError
	}
	if total == 0 {
		return errcode.EmptyData
	}
	page := common.NewPage(param.PageNum, param.PageSize)
	page.SetTotal(total)
	datas, err := dao.QueryUser(param, page)
	if err != nil {
		return errcode.QueryError
	}
	var result []*bo.UserData
	for _, v := range datas {
		updatedAt := v.UpdatedAt.Unix()
		if updatedAt < 0 {
			updatedAt = v.CreatedAt.Unix()
		}
		user := &bo.UserData{
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
		}
		result = append(result, user)
	}
	return errcode.Success.WithPage(page, result)
}
