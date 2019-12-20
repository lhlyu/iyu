package user_service

import (
	"github.com/lhlyu/iyu/cache"
	"github.com/lhlyu/iyu/common"
	"github.com/lhlyu/iyu/service/vo"
)

type Service struct {
	common.BaseService
	che *cache.Cache
}

func NewService(traceId string) *Service {
	svc := &Service{}
	svc.che = cache.NewCache(traceId)
	svc.SetTraceId(traceId)
	return svc
}

func (*Service) QueryUserById(ids ...int) []*vo.UserVo {
	return nil
}
