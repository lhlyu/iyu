package dao

import (
	"github.com/jinzhu/gorm"
	"github.com/lhlyu/iyu/common"
	"github.com/lhlyu/iyu/dao/po"
	"github.com/lhlyu/iyu/trace"
	"github.com/lhlyu/yutil/v2"
)

type CategoryDao struct {
	BaseDao
}

func NewCategoryDao(tracker trace.ITracker) CategoryDao {
	return CategoryDao{
		BaseDao: NewBaseDao(tracker),
	}
}

// 唯一
func (d CategoryDao) Add(param *po.Category) error {
	data := &po.Category{}
	tx := common.DB.Begin()
	if err := tx.Where("name = ?", param.Name).First(data).Error; err != nil {
		if err != gorm.ErrRecordNotFound {
			d.Error(err, yutil.Json.Marshal(param))
			return err
		}
	}
	if data != nil {
		return E_EX
	}
	if err := tx.Create(param).Error; err != nil {
		d.Error(err, yutil.Json.Marshal(param))
		if err := tx.Rollback().Error; err != nil {
			d.Error(err, yutil.Json.Marshal(param), "Rollback")
			return err
		}
		return err
	}
	return nil
}

// 删除关联数为0的记录
func (d CategoryDao) Del(id uint) error {
	data := &po.Category{
		Id: id,
	}
	tx := common.DB.Begin()
	if err := tx.First(data, id).Error; err != nil {
		d.Error(err, id)
		if err == gorm.ErrRecordNotFound {
			return E_NX
		}
	}
	if data.Count > 0 {
		return E_ASD
	}
	if err := tx.Where("count = 0").Delete(data).Error; err != nil {
		d.Error(err, id)
		if err := tx.Rollback().Error; err != nil {
			d.Error(err, id, "Rollback")
			return err
		}
		return err
	}
	if err := tx.Commit().Error; err != nil {
		d.Error(err, id, "Commit")
		return err
	}
	return nil
}
