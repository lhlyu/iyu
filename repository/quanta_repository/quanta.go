package quanta_repository

import (
	"github.com/lhlyu/iyu/common"
	"github.com/lhlyu/iyu/repository/po"
	"github.com/lhlyu/yutil"
)

type QuantaDao struct {
	common.BaseDao
}

func NewQuantaDao(traceId string) *QuantaDao {
	dao := &QuantaDao{}
	dao.SetTraceId(traceId)
	return dao
}

func (d *QuantaDao) Count(whr *po.YuQuanta) int {
	qb := &yutil.SqlBuffer{}
	qb.Add("select count(*) from yu_quanta where 1 = 1")
	qb.AddWhrGtZero(" and id = ?", whr.Id)
	qb.AddWhrGtZero(" and is_enable = ?", whr.IsEnable)
	qb.AddWhrNeqEmpty(" and key = ?", whr.Key)
	sql, params := qb.Build()
	var value int
	if err := common.DB.Get(&value, sql, params); err != nil {
		d.Error(err)
		return 0
	}
	return value
}

func (d *QuantaDao) QueryPage(whr *po.YuQuanta, page *common.Page) []*po.YuQuanta {
	qb := &yutil.SqlBuffer{}
	qb.Add("select * from yu_quanta where 1 = 1")
	qb.AddWhrGtZero(" and id = ?", whr.Id)
	qb.AddWhrGtZero(" and is_enable = ?", whr.IsEnable)
	qb.AddWhrNeqEmpty(" and key = ?", whr.Key)
	qb.Add(" limit ?,?", page.StartRow, page.PageSize)
	sql, params := qb.Build()
	var values []*po.YuQuanta
	if err := common.DB.Select(&values, sql, params); err != nil {
		d.Error(err)
		return nil
	}
	return values
}

func (d *QuantaDao) Get(whr *po.YuQuanta) *po.YuQuanta {
	qb := &yutil.SqlBuffer{}
	qb.Add("select * from yu_quanta where 1 = 1")
	qb.AddWhrGtZero(" and id = ?", whr.Id)
	qb.AddWhrNeqEmpty(" and key = ?", whr.Key)
	qb.Add(" limit 1")
	sql, params := qb.Build()
	var values []*po.YuQuanta
	if err := common.DB.Select(&values, sql, params); err != nil {
		d.Error(err)
		return nil
	}
	if len(values) == 0 {
		return nil
	}
	return values[0]
}

func (d *QuantaDao) Update(whr *po.YuQuanta) bool {
	qb := &yutil.SqlBuffer{}
	qb.Add("update yu_quanta set updated_at = now()")
	qb.AddWhrGtZero(",is_enable = ?", whr.Id)
	qb.AddWhrNeqEmpty(",value = ?", whr.Value)
	qb.Add(" where id = ?", whr.Id)
	sql, params := qb.Build()
	_, err := common.DB.Exec(sql, params...)
	return !d.Error(err)
}
