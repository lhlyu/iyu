package repository

import (
	"github.com/lhlyu/iyu/common"
	"github.com/lhlyu/iyu/controller/vo"
)

func (d *dao) InsertRecord(param []*vo.RecordParam) error {
	sql := "INSERT INTO yu_record(user_id,business_id,business_kind,`action`,ip)"
	var valueArr [][]interface{}
	for _, v := range param {
		var values []interface{}
		values = append(values, v.UserId, v.BusinessId, v.BusinessKind, v.Action, v.Ip)
		valueArr = append(valueArr, values)
	}
	batchSql, params := d.createQuestionMarksForBatch(valueArr...)
	sql += batchSql
	if _, err := common.DB.Exec(sql, params...); err != nil {
		common.Ylog.Debug(err)
		return err
	}
	return nil
}
