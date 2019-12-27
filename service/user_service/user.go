package user_service

import (
	"github.com/lhlyu/iyu/cache"
	"github.com/lhlyu/iyu/common"
	"github.com/lhlyu/iyu/repository/po"
	"github.com/lhlyu/iyu/repository/user_repository"
	"github.com/lhlyu/iyu/service/vo"
)

type Service struct {
	common.BaseService
	dao *user_repository.Dao
	che *cache.Cache
}

func NewService(tracker *common.Tracker) *Service {
	svc := &Service{}
	svc.che = cache.NewCache(tracker)
	svc.dao = user_repository.NewDao(tracker)
	svc.SetTracker(tracker)
	return svc
}

func (s *Service) QueryUserById(ids ...int) []*vo.UserVo {
	items := s.che.GetUser(ids...)
	if len(items) == 0 {
		datas := s.dao.GetByIds(ids...)
		items = _UserPosToUserVos(datas)
	}
	if len(items) > 0 {
		go s.che.SetUser(items)
	}
	return items
}

func _UserPosToUserVos(datas []*po.YuUser) []*vo.UserVo {
	var items []*vo.UserVo
	for _, v := range datas {
		item := &vo.UserVo{
			Id:        v.Id,
			ThirdId:   v.ThirdId,
			Role:      v.Role,
			From:      v.From,
			Status:    v.Status,
			AvatarUrl: v.AvatarUrl,
			UserUrl:   v.UserUrl,
			UserName:  v.UserName,
			Bio:       v.Bio,
			Sign:      v.Sign,
			Ip:        v.Sign,
			LastLogin: v.LastLogin.Unix(),
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
