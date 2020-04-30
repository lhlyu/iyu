package dao

import (
	"github.com/jinzhu/gorm"
	"github.com/lhlyu/iyu/common"
	"github.com/lhlyu/iyu/dao/po"
	"github.com/lhlyu/iyu/trace"
	"github.com/lhlyu/yutil/v2"
)

type CategoryDao struct {
	trace.BaseTracker
}

func NewCategoryDao(tracker trace.ITracker) CategoryDao {
	return CategoryDao{
		BaseTracker: trace.NewBaseTracker(tracker),
	}
}

func (d CategoryDao) Get(id uint) (*po.Category, error) {
	data := &po.Category{}
	if err := common.DB.First(data, id).Error; err != nil {
		d.Error(err, id)
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return data, nil
}

func (d CategoryDao) Query(whr string, params ...interface{}) ([]*po.Category, error) {
	var datas []*po.Category
	if err := common.DB.Where(whr, params...).Find(&datas).Error; err != nil {
		d.Error(err, whr, params)
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return datas, nil
}

func (d CategoryDao) Add(param *po.Category) error {
	if err := common.DB.Create(param).Error; err != nil {
		d.Error(err, yutil.Json.Marshal(param))
		return err
	}
	return nil
}

func (d CategoryDao) Del(id uint) error {
	data := &po.Category{
		Id: id,
	}
	if err := common.DB.Where("count = 0").Delete(data).Error; err != nil {
		d.Error(err, id)
		return err
	}
	return nil
}

func (d CategoryDao) Update(param *po.Category) error {
	if err := common.DB.Model(param).Updates(param).Error; err != nil {
		d.Error(err, yutil.Json.Marshal(param))
		return err
	}
	return nil
}
