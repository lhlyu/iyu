package dao

import (
	"errors"
	"github.com/jinzhu/gorm"
	"github.com/lhlyu/iyu/common"
	"github.com/lhlyu/iyu/trace"
	"github.com/lhlyu/yutil/v2"
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

// 获取单个
func (d BaseDao) Get(data interface{}, whr string, params ...interface{}) error {
	if err := common.DB.Where(whr, params...).First(data).Error; err != nil {
		d.Error(err, yutil.Json.Marshal(data), whr, params)
		if err == gorm.ErrRecordNotFound {
			return nil
		}
		return err
	}
	return nil
}

// 查询
func (d BaseDao) Query(datas interface{}, whr string, params ...interface{}) error {
	if err := common.DB.Where(whr, params...).Find(datas).Error; err != nil {
		d.Error(err, yutil.Json.Marshal(datas), whr, params)
		if err == gorm.ErrRecordNotFound {
			return nil
		}
		return err
	}
	return nil
}

// 查询分页
func (d BaseDao) QueryPage(datas interface{}, page *common.Page, orderby, whr string, params ...interface{}) error {
	return d.Transact(func(tx *gorm.DB) error {
		var total int
		if err := tx.Model(datas).Where(whr, params...).Count(&total).Error; err != nil {
			d.Error(err, yutil.Json.Marshal(datas), yutil.Json.Marshal(page), orderby, whr, params)
			return err
		}
		page.SetTotal(total)
		if err := tx.Where(whr, params...).Limit(page.PageSize).Offset(page.StartRow).Order(orderby).Find(datas).Error; err != nil {
			d.Error(err, yutil.Json.Marshal(datas), yutil.Json.Marshal(page), orderby, whr, params)
			if err == gorm.ErrRecordNotFound {
				return nil
			}
			return err
		}
		return nil
	})
}

// 添加
func (d BaseDao) Add(data interface{}) error {
	if err := common.DB.Create(data).Error; err != nil {
		d.Error(err, yutil.Json.Marshal(data))
		return err
	}
	return nil
}

// 删除
func (d BaseDao) Del(data interface{}, whr string, params ...interface{}) error {
	if err := common.DB.Where(whr, params...).Delete(data).Error; err != nil {
		d.Error(err, yutil.Json.Marshal(data), whr, params)
		return err
	}
	return nil
}

// 更新
func (d BaseDao) Update(data interface{}, whr string, params ...interface{}) error {
	if err := common.DB.Model(data).Where(whr, params...).Updates(data).Error; err != nil {
		d.Error(err, yutil.Json.Marshal(data), whr, params)
		return err
	}
	return nil
}

// 事务
func (d BaseDao) Transact(fn func(tx *gorm.DB) error) error {
	tx := common.DB.Begin()
	if err := fn(tx); err != nil {
		d.Error(err)
		if err = tx.Rollback().Error; err != nil {
			d.Error(err)
		}
		return err
	}
	if err := tx.Commit().Error; err != nil {
		d.Error(err)
		return err
	}
	return nil
}
