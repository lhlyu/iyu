package dao

import (
	"github.com/jinzhu/gorm"
	"github.com/lhlyu/iyu/dao/po"
	"github.com/lhlyu/iyu/trace"
	"github.com/lhlyu/yutil/v2"
	"time"
)

type ArticleDao struct {
	BaseDao
}

func NewArticleDao(tracker trace.ITracker) ArticleDao {
	return ArticleDao{
		BaseDao: NewBaseDao(tracker),
	}
}

func (d ArticleDao) Add(param *po.Article) error {
	return d.Transact(func(tx *gorm.DB) error {
		if err := tx.Create(param).Error; err != nil {
			d.Error(err, yutil.Json.Marshal(param))
			return err
		}
		data := &po.Category{}
		if err := tx.First(data, param.Category).Error; err != nil {
			d.Error(err, yutil.Json.Marshal(param))
			return err
		}
		data.Count++
		now := time.Now()
		data.CreatedAt = now
		data.UpdatedAt = now
		if err := tx.Model(data).Updates(data).Error; err != nil {
			d.Error(err, yutil.Json.Marshal(param), yutil.Json.Marshal(data))
			return err
		}
		return nil
	})
}

func (d ArticleDao) Update(param *po.Article) error {
	return d.Transact(func(tx *gorm.DB) error {
		now := time.Now()
		old := &po.Article{}
		if err := tx.First(old, param.Id).Error; err != nil {
			d.Error(err, yutil.Json.Marshal(param))
			return err
		}
		if old.Category != param.Category {
			// 旧分类数量减一
			category1 := &po.Category{}
			if err := tx.First(category1, old.Category).Error; err != nil {
				d.Error(err, yutil.Json.Marshal(param), yutil.Json.Marshal(old))
				return err
			}
			if category1.Count > 0 {
				category1.Count--
			}
			category1.UpdatedAt = now
			if err := tx.Model(category1).Updates(category1).Error; err != nil {
				d.Error(err, yutil.Json.Marshal(param), yutil.Json.Marshal(category1))
				return err
			}
			// 新分类数量加一
			category2 := &po.Category{}
			if err := tx.First(category2, param.Category).Error; err != nil {
				d.Error(err, yutil.Json.Marshal(param))
				return err
			}
			category2.Count++
			category2.UpdatedAt = now
			if err := tx.Model(category2).Updates(category2).Error; err != nil {
				d.Error(err, yutil.Json.Marshal(param), yutil.Json.Marshal(category2))
				return err
			}
		}
		param.UpdatedAt = now
		if err := tx.Model(param).Updates(param).Error; err != nil {
			d.Error(err, yutil.Json.Marshal(param))
			return err
		}
		return nil
	})
}

// 删除
func (d ArticleDao) Del(param *po.Article) error {
	return d.Transact(func(tx *gorm.DB) error {
		now := time.Now()
		old := &po.Article{}
		if err := tx.First(old, param.Id).Error; err != nil {
			d.Error(err, yutil.Json.Marshal(param))
			return err
		}
		// 分类数量减一
		category := &po.Category{}
		if err := tx.First(category, old.Category).Error; err != nil {
			d.Error(err, yutil.Json.Marshal(param), yutil.Json.Marshal(old))
			return err
		}
		if category.Count > 0 {
			category.Count--
		}
		category.UpdatedAt = now
		if err := tx.Model(category).Updates(category).Error; err != nil {
			d.Error(err, yutil.Json.Marshal(param), yutil.Json.Marshal(category))
			return err
		}
		return nil
	})
}
