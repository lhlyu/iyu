package dao

import (
	"github.com/lhlyu/iyu/common"
	"github.com/lhlyu/iyu/dao/po"
	"github.com/lhlyu/iyu/trace"
)

type RecordDao struct {
	trace.BaseTracker
}

func NewRecordDao(tracker trace.ITracker) *RecordDao {
	return &RecordDao{
		BaseTracker: trace.NewBaseTracker(tracker),
	}
}

// 添加记录
func (d *RecordDao) Add(v *po.Record) bool {
	err := common.DB.Create(v).Error
	if err != nil {
		d.Error(err, v)
		return false
	}
	return true
}

// 查询浏览量
func (d *RecordDao) QueryVisit(ids ...int) map[int]int {
	m := make(map[int]int)
	rows, err := common.DB.Model(&po.Record{}).Select("target_id,count(id)").Where("kind = ?", po.RECORD_KIND_SYSTEM).Where("target_id in (?)", ids).Group("target_id").Rows()
	if err != nil {
		d.Error(err, ids)
		return nil
	}
	for rows.Next() {
		var targetId, count int
		err = rows.Scan(&targetId, &count)
		if err != nil {
			d.Error(err, ids)
			return nil
		}
		m[targetId] = count
	}
	return m
}

// 查询记录
func (d *RecordDao) Query(v *po.Record, page *common.Page) []*po.Record {
	items := []*po.Record{}
	var total int
	err := common.DB.Model(v).Where(v).Count(&total).Error
	if err != nil {
		d.Error(err, v)
		return nil
	}
	page.SetTotal(total)
	err = common.DB.Where(v).Offset(page.StartRow).Limit(page.PageSize).Find(&items).Error
	if err != nil {
		d.Error(err, v)
		return nil
	}
	return items
}
