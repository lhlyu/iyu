package dao

import (
	"github.com/lhlyu/iyu/common"
	"github.com/lhlyu/iyu/dao/po"
	"github.com/lhlyu/iyu/trace"
)

type PlateDao struct {
	trace.BaseTracker
}

func NewPlateDao(tracker trace.ITracker) *PlateDao {
	return &PlateDao{
		BaseTracker: trace.NewBaseTracker(tracker),
	}
}

// 修改
func (d *PlateDao) Edit(v *po.Plate) bool {
	tx := common.DB.Begin()
	old := &po.Plate{}
	if err := tx.First(old, v.Id).Error; err != nil {
		d.Error(err, v.Id)
		return false
	}
	if old == nil {
		return false
	}

	return true
}

// 查询
func (*PlateDao) Query() []*po.Plate {
	return nil
}
