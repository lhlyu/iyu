package tag_repository

import (
	"github.com/lhlyu/iyu/common"
	"github.com/lhlyu/iyu/repository/po"
	"github.com/lhlyu/yutil"
)

type Dao struct {
	common.BaseDao
}

func NewDao(traceId string) *Dao {
	dao := &Dao{}
	dao.SetTraceId(traceId)
	return dao
}

func (d *Dao) Count(whr *po.YuTag) int {
	qb := &yutil.SqlBuffer{}
	qb.Add("select count(*) from yu_tag where 1 = 1")
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

func (d *Dao) QueryPage(whr *po.YuTag, page *common.Page) []*po.YuTag {
	qb := &yutil.SqlBuffer{}
	qb.Add("select * from yu_tag where 1 = 1")
	qb.AddWhrGtZero(" and id = ?", whr.Id)
	qb.AddWhrGtZero(" and is_delete = ?", whr.IsDelete)
	qb.AddWhrNeqEmpty(" and name = ?", whr.Name)
	qb.Add(" limit ?,?", page.StartRow, page.PageSize)
	sql, params := qb.Build()
	var values []*po.YuTag
	if err := common.DB.Select(&values, sql, params...); err != nil {
		d.Error(err, sql, params)
		return nil
	}
	return values
}

func (d *Dao) Get(whr *po.YuTag) *po.YuTag {
	qb := &yutil.SqlBuffer{}
	qb.Add("select * from yu_tag where 1 = 1")
	qb.AddWhrGtZero(" and id = ?", whr.Id)
	qb.AddWhrNeqEmpty(" and name = ?", whr.Name)
	qb.Add(" limit 1")
	sql, params := qb.Build()
	var values []*po.YuTag
	if err := common.DB.Select(&values, sql, params...); err != nil {
		d.Error(err, sql, params)
		return nil
	}
	if len(values) == 0 {
		return nil
	}
	return values[0]
}

func (d *Dao) GetByIds(ids ...int) []*po.YuTag {
	qb := &yutil.SqlBuffer{}
	qb.Add("select * from yu_tag where 1 = 1")
	qb.AddWhrIn(" and id in (%s)", yutil.ConvertIntToInterface(ids)...)
	sql, params := qb.Build()
	var values []*po.YuTag
	if err := common.DB.Select(&values, sql, params...); err != nil {
		d.Error(err, sql, params)
		return nil
	}
	return values
}

func (d *Dao) Add(whr *po.YuTag) int {
	qb := &yutil.SqlBuffer{}
	qb.Add("insert into yu_tag(name) values(?,?)", whr.Name)
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

func (d *Dao) Update(whr *po.YuTag) bool {
	qb := &yutil.SqlBuffer{}
	qb.Add("update yu_tag set updated_at = now()")
	qb.AddWhrGtZero(",is_delete = ?", whr.IsDelete)
	qb.AddWhrNeqEmpty(",name = ?", whr.Name)
	qb.Add(" where id = ?", whr.Id)
	sql, params := qb.Build()
	_, err := common.DB.Exec(sql, params...)
	return !d.Error(err, sql, params)
}

func (d *Dao) Delete(id int) bool {
	qb := &yutil.SqlBuffer{}
	qb.Add("delete from yu_tag where id = ?", id)
	sql, params := qb.Build()
	_, err := common.DB.Exec(sql, params...)
	return !d.Error(err, sql, params)
}
