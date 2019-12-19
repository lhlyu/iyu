package quanta_service

import (
	"github.com/lhlyu/iyu/cache"
	"github.com/lhlyu/iyu/common"
	"github.com/lhlyu/iyu/repository"
)

type QuantaService struct {
	common.BaseService
	dao *repository.Dao
	che *cache.Cache
}

func NewService(traceId string) *QuantaService {
	svc := &QuantaService{}
	svc.dao = repository.NewDao(traceId)
	svc.che = cache.NewCache(traceId)
	svc.SetTraceId(traceId)
	return svc
}

func (s *QuantaService) QueryQuanta() {
	s.Info("QuantaService")
}
