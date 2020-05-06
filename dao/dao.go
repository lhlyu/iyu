package dao

import (
	"github.com/jinzhu/gorm"
	"github.com/lhlyu/iyu/trace"
	"github.com/pkg/errors"
)

var (
	E_NX  = gorm.ErrRecordNotFound
	E_ASD = errors.New("存在关联关系，禁止删除")
	E_EX  = errors.New("数据已经存在")
)

type BaseDao struct {
	trace.BaseTracker
}

func NewBaseDao(tracker trace.ITracker) BaseDao {
	return BaseDao{
		BaseTracker: trace.NewBaseTracker(tracker),
	}
}
