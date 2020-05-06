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

func (d BaseDao) Get(data interface{},whr string,params ...interface{}) error{
    if err := common.DB.Where(whr,params...).First(data).Error;err != nil{
        d.Error(err, yutil.Json.Marshal(data),whr,params)
        if err == gorm.ErrRecordNotFound {
            return nil
        }
        return err
    }
    return nil
}

func (d BaseDao) Query(datas interface{},whr string,params ...interface{}) error{
    if err := common.DB.Where(whr,params...).Find(datas).Error;err != nil{
        d.Error(err, yutil.Json.Marshal(datas),whr,params)
        if err == gorm.ErrRecordNotFound {
            return nil
        }
        return err
    }
    return nil
}

func (d BaseDao) Add(data interface{}) error {
    if err := common.DB.Create(data).Error; err != nil {
        d.Error(err, yutil.Json.Marshal(data))
        return err
    }
    return nil
}

func (d BaseDao) Del(data interface{},whr string, params ...interface{}) error {
    if err := common.DB.Where(whr, params...).Delete(data).Error; err != nil {
        d.Error(err, yutil.Json.Marshal(data), whr, params)
        return err
    }
    return nil
}

func (d BaseDao) Update(data interface{},whr string, params ...interface{}) error {
    if err := common.DB.Model(data).Where(whr,params...).Updates(data).Error; err != nil {
        d.Error(err, yutil.Json.Marshal(data), whr, params)
        return err
    }
    return nil
}
