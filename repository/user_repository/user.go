package user_repository

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

func (d *Dao) GetByIds(ids ...int) []*po.YuUser {
	qb := &yutil.SqlBuffer{}
	qb.Add("select * from yu_user where 1 = 1")
	qb.AddWhrIn(" and id in (%s)", yutil.ConvertIntToInterface(ids)...)
	sql, params := qb.Build()
	var values []*po.YuUser
	if err := common.DB.Select(&values, sql, params...); err != nil {
		d.Error(err, sql, params)
		return nil
	}
	return values
}
