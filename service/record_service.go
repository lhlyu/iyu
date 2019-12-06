package service

import (
	"github.com/lhlyu/iyu/controller/vo"
	"github.com/lhlyu/iyu/errcode"
	"github.com/lhlyu/iyu/repository"
)

type recordService struct {
}

func NewRecordService() *recordService {
	return &recordService{}
}

func (*recordService) Insert(param *vo.RecordParam) *errcode.ErrCode {
	if err := repository.NewDao().InsertRecord([]*vo.RecordParam{param}); err != nil {
		return errcode.InsertError
	}
	svc := NewArticleService()
	svc.Query(true, param.BusinessId)
	return errcode.Success
}

func (*recordService) BatchInsert(param []*vo.RecordParam) *errcode.ErrCode {
	if err := repository.NewDao().InsertRecord(param); err != nil {
		return errcode.InsertError
	}
	return errcode.Success
}
