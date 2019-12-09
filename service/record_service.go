package service

import (
	"github.com/lhlyu/iyu/controller/vo"
	"github.com/lhlyu/iyu/errcode"
	"github.com/lhlyu/iyu/repository"
)

type recordService struct {
	*Service
}

func NewRecordService(traceId string) *recordService {
	return &recordService{
		Service: &Service{traceId},
	}
}

func (s *recordService) Insert(param *vo.RecordParam) *errcode.ErrCode {
	if err := repository.NewDao(s.TraceId).InsertRecord([]*vo.RecordParam{param}); err != nil {
		return errcode.InsertError
	}
	svc := NewArticleService(s.TraceId)
	svc.Query(true, param.BusinessId)
	return errcode.Success
}

func (s *recordService) BatchInsert(param []*vo.RecordParam) *errcode.ErrCode {
	if err := repository.NewDao(s.TraceId).InsertRecord(param); err != nil {
		return errcode.InsertError
	}
	return errcode.Success
}
