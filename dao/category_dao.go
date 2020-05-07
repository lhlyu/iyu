package dao

import (
	"github.com/jinzhu/gorm"
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
	return d.Transact(func(tx *gorm.DB) error {
		data := &po.Category{}
		if err := tx.Where("name = ?", param.Name).First(data).Error; err != nil {
			d.Error(err, yutil.Json.Marshal(param))
			if err != gorm.ErrRecordNotFound {
				return err
			}
		}
		if data.Id > 0 {
			return E_EX
		}
		if err := tx.Create(param).Error; err != nil {
			d.Error(err, yutil.Json.Marshal(param))
			return err
		}
		return nil
	})
}

// 删除关联数为0的记录
func (d CategoryDao) Del(id uint) error {
	return d.Transact(func(tx *gorm.DB) error {
		data := &po.Category{
			Id: id,
		}
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
			return err
		}
		return nil
	})
}
