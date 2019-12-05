package repository

import (
	"fmt"
	"github.com/lhlyu/iyu/common"
	"github.com/lhlyu/iyu/controller/vo"
	"github.com/lhlyu/iyu/repository/po"
)

func (d *dao) QueryNail(id ...int) []*po.YuNail {
	sql := "SELECT * FROM yu_nail"
	var params []interface{}
	if len(id) > 0 {
		marks := d.createQuestionMarks(len(id))
		params = d.intConvertToInterface(id)
		sql += fmt.Sprintf(" where id in (%s)", marks)
	}
	sql += " ORDER BY is_delete"
	var values []*po.YuNail
	if err := common.DB.Select(&values, sql, params...); err != nil {
		common.Ylog.Debug(err)
		return nil
	}
	return values
}

// get nail by name
func (d *dao) GetNailByName(name string) *po.YuNail {
	sql := "SELECT * FROM yu_nail WHERE `name` = ? limit 1"
	value := &po.YuNail{}
	if err := common.DB.Get(value, sql, name); err != nil {
		common.Ylog.Debug(err)
		return nil
	}
	return value
}

func (d *dao) GetNailById(id int) *po.YuNail {
	sql := "SELECT * FROM yu_nail WHERE id = ? limit 1"
	value := &po.YuNail{}
	if err := common.DB.Get(value, sql, id); err != nil {
		common.Ylog.Debug(err)
		return nil
	}
	return value
}

// update nail
func (d *dao) UpdateNail(param *po.YuNail) error {
	sql := "UPDATE yu_nail SET is_delete=?,`name` = ?,color = ?,updated_at = NOW() WHERE id = ?"
	if _, err := common.DB.Exec(sql, param.IsDelete, param.Name, param.Color, param.Id); err != nil {
		common.Ylog.Debug(err)
		return err
	}
	return nil
}

// delete by id
func (d *dao) DeleteNailById(id int) error {
	sql := "DELETE FROM yu_nail WHERE id = ?"
	if _, err := common.DB.Exec(sql, id); err != nil {
		common.Ylog.Debug(err)
		return err
	}
	return nil
}

// add nail
func (d *dao) InsertNail(param *vo.NailVo) (int, error) {
	sql := "INSERT INTO yu_nail(`name`,color) VALUES(?)"
	result, err := common.DB.Exec(sql, param.Name, param.Color)
	if err != nil {
		common.Ylog.Debug(err)
		return 0, err
	}
	id, _ := result.LastInsertId()
	return int(id), nil
}
