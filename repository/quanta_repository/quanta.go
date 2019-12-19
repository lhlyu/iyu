package quanta_repository

import (
	"github.com/lhlyu/iyu/common"
)

type QuantaDao struct {
	common.BaseDao
}

func NewQuantaDao(traceId string) *QuantaDao {
	dao := &QuantaDao{}
	dao.SetTraceId(traceId)
	return dao
}

func (d *QuantaDao) Query() {
	d.Info("test")
}
