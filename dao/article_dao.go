package dao

import (
	"github.com/jinzhu/gorm"
	"github.com/lhlyu/iyu/common"
	"github.com/lhlyu/iyu/dao/po"
	"github.com/lhlyu/iyu/trace"
	"github.com/lhlyu/yutil/v2"
)

type ArticleDao struct {
	trace.BaseTracker
}

func NewArticleDao(tracker trace.ITracker) ArticleDao {
	return ArticleDao{
		BaseTracker: trace.NewBaseTracker(tracker),
	}
}

func (d ArticleDao) Get(whr string, params ...interface{}) (*po.Article, error) {
	data := &po.Article{}
	if err := common.DB.Where(whr, params...).First(data).Error; err != nil {
		d.Error(err, whr, params)
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return data, nil
}

func (d ArticleDao) Query(whr string, params ...interface{}) ([]*po.Article, error) {
	var datas []*po.Article
	if err := common.DB.Where(whr, params...).Find(&datas).Error; err != nil {
		d.Error(err, whr, params)
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return datas, nil
}

func (d ArticleDao) Add(param *po.Article) error {
	if err := common.DB.Create(param).Error; err != nil {
		d.Error(err, yutil.Json.Marshal(param))
		return err
	}
	return nil
}

func (d ArticleDao) Del(whr string, params ...interface{}) error {
	data := &po.Article{}
	if err := common.DB.Where(whr, params...).Delete(data).Error; err != nil {
		d.Error(err, whr, params)
		return err
	}
	return nil
}

func (d ArticleDao) Update(param *po.Article) error {
	if err := common.DB.Model(param).Updates(param).Error; err != nil {
		d.Error(err, yutil.Json.Marshal(param))
		return err
	}
	return nil
}
