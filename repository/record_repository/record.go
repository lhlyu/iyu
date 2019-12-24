package record_repository

import (
	"github.com/lhlyu/iyu/common"
	"github.com/lhlyu/iyu/controller/dto"
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

func (d *Dao) Count(whr *dto.RecordDto) int {
	qb := &yutil.SqlBuffer{}
	qb.Add("select count(*) from yu_record where 1 = 1")
	qb.AddWhrGtZero(" and user_id = ?", whr.UserId)
	qb.AddWhrGtZero(" and business_id = ?", whr.BusinessId)
	qb.AddWhrGtZero(" and business_kind = ?", whr.BusinessKind)
	qb.AddWhrNeqEmpty(" and content like ?", "%"+whr.Content+"%")
	qb.AddWhrNeqEmpty(" and agent like ?", "%"+whr.Agent+"%")
	qb.AddWhrNeqEmpty(" and ip like ?", "%"+whr.Ip+"%")
	qb.AddWhrNeqEmpty(" and created_at > ?", whr.CreatedAt)
	sql, params := qb.Build()
	var value int
	if err := common.DB.Get(&value, sql, params...); err != nil {
		d.Error(err, sql, params)
		return 0
	}
	return value
}

func (d *Dao) QueryPage(whr *dto.RecordDto, page *common.Page) []*po.YuRecord {
	qb := &yutil.SqlBuffer{}
	qb.Add("select * from yu_record where 1 = 1")
	qb.AddWhrGtZero(" and user_id = ?", whr.UserId)
	qb.AddWhrGtZero(" and business_id = ?", whr.BusinessId)
	qb.AddWhrGtZero(" and business_kind = ?", whr.BusinessKind)
	qb.AddWhrNeqEmpty(" and content like ?", "%"+whr.Content+"%")
	qb.AddWhrNeqEmpty(" and agent like ?", "%"+whr.Agent+"%")
	qb.AddWhrNeqEmpty(" and ip like ?", "%"+whr.Ip+"%")
	qb.AddWhrNeqEmpty(" and created_at >= ?", whr.CreatedAt)
	qb.Add(" order by create_at desc limit ?,?", page.StartRow, page.PageSize)
	sql, params := qb.Build()
	var values []*po.YuRecord
	if err := common.DB.Select(&values, sql, params...); err != nil {
		d.Error(err, sql, params)
		return nil
	}
	return values
}

func (d *Dao) Add(whr *po.YuRecord) bool {
	sql := "insert into yu_record(user_id,business_id,content,business_kind,agent,ip,created_at,updated_at) values(?,?,?,?,?,?,now(),now())"
	_, err := common.DB.Exec(sql, whr.UserId, whr.BusinessId, whr.Content, whr.BusinessKind, whr.Agent, whr.Ip)
	if err != nil {
		d.Error(err, sql, yutil.JsonObjToStr(whr))
		return false
	}
	return true
}
