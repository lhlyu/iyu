package record_service

import (
	"github.com/lhlyu/iyu/common"
	"github.com/lhlyu/iyu/controller/dto"
	"github.com/lhlyu/iyu/repository/po"
	"github.com/lhlyu/iyu/repository/record_repository"
)

type Service struct {
	common.BaseService
	dao *record_repository.Dao
}

func NewService(tracker *common.Tracker) *Service {
	svc := &Service{}
	svc.dao = record_repository.NewDao(tracker)
	svc.SetTracker(tracker)
	return svc
}

func (s *Service) AddRecord(param *dto.RecordDto) bool {
	whr := &po.YuRecord{
		UserId:       param.UserId,
		BusinessId:   param.BusinessId,
		Content:      param.Content,
		BusinessKind: param.BusinessKind,
		Agent:        param.Agent,
		Ip:           param.Ip,
	}
	return s.dao.Add(whr)
}
