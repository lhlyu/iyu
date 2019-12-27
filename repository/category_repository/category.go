package category_repository

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

func (d *Dao) Count(whr *po.YuCategory) int {
	qb := &yutil.SqlBuffer{}
	qb.Add("select count(*) from yu_category where 1 = 1")
	qb.AddWhrGtZero(" and id = ?", whr.Id)
	qb.AddWhrGtZero(" and is_delete = ?", whr.IsDelete)
	qb.AddWhrNeqEmpty(" and name = ?", whr.Name)
	sql, params := qb.Build()
	var value int
	if err := common.DB.Get(&value, sql, params...); err != nil {
		d.Error(err, sql, params)
		return 0
	}
	return value
}

func (d *Dao) QueryPage(whr *po.YuCategory, page *common.Page) []*po.YuCategory {
	qb := &yutil.SqlBuffer{}
	qb.Add("select * from yu_category where 1 = 1")
	qb.AddWhrGtZero(" and id = ?", whr.Id)
	qb.AddWhrGtZero(" and is_delete = ?", whr.IsDelete)
	qb.AddWhrNeqEmpty(" and name = ?", whr.Name)
	qb.Add(" limit ?,?", page.StartRow, page.PageSize)
	sql, params := qb.Build()
	var values []*po.YuCategory
	if err := common.DB.Select(&values, sql, params...); err != nil {
		d.Error(err, sql, params)
		return nil
	}
	return values
}

func (d *Dao) Get(whr *po.YuCategory) *po.YuCategory {
	qb := &yutil.SqlBuffer{}
	qb.Add("select * from yu_category where 1 = 1")
	qb.AddWhrGtZero(" and id = ?", whr.Id)
	qb.AddWhrNeqEmpty(" and name = ?", whr.Name)
	qb.Add(" limit 1")
	sql, params := qb.Build()
	var values []*po.YuCategory
	if err := common.DB.Select(&values, sql, params...); err != nil {
		d.Error(err, sql, params)
		return nil
	}
	if len(values) == 0 {
		return nil
	}
	return values[0]
}

func (d *Dao) GetByIds(ids ...int) []*po.YuCategory {
	qb := &yutil.SqlBuffer{}
	qb.Add("select * from yu_category where 1 = 1")
	qb.AddWhrIn(" and id in (%s)", yutil.ConvertIntToInterface(ids)...)
	sql, params := qb.Build()
	var values []*po.YuCategory
	if err := common.DB.Select(&values, sql, params...); err != nil {
		d.Error(err, sql, params)
		return nil
	}
	return values
}

func (d *Dao) Add(whr *po.YuCategory) int {
	qb := &yutil.SqlBuffer{}
	qb.Add("insert into yu_category(name,color) values(?,?)", whr.Name, whr.Color)
	sql, params := qb.Build()
	rs, err := common.DB.Exec(sql, params...)
	if err != nil {
		d.Error(err, sql, params)
		return 0
	}
	id, err := rs.LastInsertId()
	if err != nil {
		d.Error(err, sql, params)
		return 0
	}
	return int(id)
}

func (d *Dao) Update(whr *po.YuCategory) bool {
	qb := &yutil.SqlBuffer{}
	qb.Add("update yu_category set updated_at = now()")
	qb.AddWhrGtZero(",is_delete = ?", whr.IsDelete)
	qb.AddWhrNeqEmpty(",name = ?", whr.Name)
	qb.AddWhrNeqEmpty(",color = ?", whr.Color)
	qb.Add(" where id = ?", whr.Id)
	sql, params := qb.Build()
	_, err := common.DB.Exec(sql, params...)
	return !d.Error(err, sql, params)
}

func (d *Dao) Delete(id int) bool {
	qb := &yutil.SqlBuffer{}
	qb.Add("delete from yu_category where id = ?", id)
	sql, params := qb.Build()
	_, err := common.DB.Exec(sql, params...)
	return !d.Error(err, sql, params)
}
