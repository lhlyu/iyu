package repository

import (
	"github.com/lhlyu/iyu/common"
	"github.com/lhlyu/iyu/repository/po"
)

// get all nails
func (d *dao) GetNailAll() []*po.YuNail {
	sql := "SELECT * FROM yu_nail ORDER BY is_delete,updated_at DESC,created_at DESC"
	var values []*po.YuNail
	if err := common.DB.Select(&values, sql); err != nil {
		common.Ylog.Debug(err)
		return nil
	}
	return values
}

// get nail by name
func (d *dao) GetNailByName(id int, name string) *po.YuNail {
	sql := "SELECT * FROM yu_nail WHERE `name` = ? and id <> ? limit 1"
	value := &po.YuNail{}
	if err := common.DB.Get(value, sql, name, id); err != nil {
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
func (d *dao) UpdateNail(id, status int, name, color string) error {
	sql := "UPDATE yu_nail SET is_delete=?,`name` = ?,color = ?,updated_at = NOW() WHERE id = ?"
	if _, err := common.DB.Exec(sql, status, name, color, id); err != nil {
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
func (d *dao) InsertNail(name, color string) error {
	sql := "INSERT INTO yu_nail(`name`,color) VALUES(?)"
	if _, err := common.DB.Exec(sql, name, color); err != nil {
		common.Ylog.Debug(err)
		return err
	}
	return nil
}
