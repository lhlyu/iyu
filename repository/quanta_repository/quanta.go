package quanta_repository

import (
	"github.com/lhlyu/iyu/common"
	"github.com/lhlyu/iyu/repository/po"
	"github.com/lhlyu/yutil"
)

type Dao struct {
	common.BaseDao
}

func NewDao(tracker *common.Tracker) *Dao {
	dao := &Dao{}
	dao.SetTracker(tracker)
	return dao
}

func (d *Dao) Count(whr *po.YuQuanta) int {
	qb := &yutil.SqlBuffer{}
	qb.Add("select count(*) from yu_quanta where 1 = 1")
	qb.AddWhrGtZero(" and id = ?", whr.Id)
	qb.AddWhrGtZero(" and is_enable = ?", whr.IsEnable)
	qb.AddWhrNeqEmpty(" and key = ?", whr.Key)
	sql, params := qb.Build()
	var value int
	if err := common.DB.Get(&value, sql, params...); err != nil {
		d.Error(err, sql, params)
		return 0
	}
	return value
}

func (d *Dao) QueryPage(whr *po.YuQuanta, page *common.Page) []*po.YuQuanta {
	qb := &yutil.SqlBuffer{}
	qb.Add("select * from yu_quanta where 1 = 1")
	qb.AddWhrGtZero(" and id = ?", whr.Id)
	qb.AddWhrGtZero(" and is_enable = ?", whr.IsEnable)
	qb.AddWhrNeqEmpty(" and key = ?", whr.Key)
	qb.Add(" limit ?,?", page.StartRow, page.PageSize)
	sql, params := qb.Build()
	var values []*po.YuQuanta
	if err := common.DB.Select(&values, sql, params...); err != nil {
		d.Error(err, sql, params)
		return nil
	}
	return values
}

func (d *Dao) Get(whr *po.YuQuanta) *po.YuQuanta {
	qb := &yutil.SqlBuffer{}
	qb.Add("select * from yu_quanta where 1 = 1")
	qb.AddWhrGtZero(" and id = ?", whr.Id)
	qb.AddWhrNeqEmpty(" and key = ?", whr.Key)
	qb.Add(" limit 1")
	sql, params := qb.Build()
	var values []*po.YuQuanta
	if err := common.DB.Select(&values, sql, params...); err != nil {
		d.Error(err, sql, params)
		return nil
	}
	if len(values) == 0 {
		return nil
	}
	return values[0]
}

func (d *Dao) QueryByKeys(keys ...string) []*po.YuQuanta {
	qb := &yutil.SqlBuffer{}
	qb.Add("select * from yu_quanta where 1 = 1")
	qb.AddWhrIn(" and key in (%s)", yutil.ConvertStrToInterface(keys)...)
	sql, params := qb.Build()
	var values []*po.YuQuanta
	if err := common.DB.Select(&values, sql, params...); err != nil {
		d.Error(err, sql, params)
		return nil
	}
	return values
}

func (d *Dao) Update(whr *po.YuQuanta) bool {
	qb := &yutil.SqlBuffer{}
	qb.Add("update yu_quanta set updated_at = now()")
	qb.AddWhrGtZero(",is_enable = ?", whr.IsEnable)
	qb.AddWhrNeqEmpty(",value = ?", whr.Value)
	qb.Add(" where id = ?", whr.Id)
	sql, params := qb.Build()
	_, err := common.DB.Exec(sql, params...)
	return !d.Error(err, sql, params)
}
