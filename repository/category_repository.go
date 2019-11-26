package repository

import (
	"github.com/lhlyu/iyu/common"
	"github.com/lhlyu/iyu/repository/po"
)

// get all categorys
func (d *dao) GetCategoryAll() []*po.YuCategory {
	sql := "SELECT * FROM yu_category ORDER BY is_delete,updated_at DESC,created_at DESC"
	var values []*po.YuCategory
	if err := common.DB.Select(&values, sql); err != nil {
		common.Ylog.Debug(err)
		return nil
	}
	return values
}

// get category by name
func (d *dao) GetCategoryByName(id int, name string) *po.YuCategory {
	sql := "SELECT * FROM yu_category WHERE `name` = ? and id <> ? limit 1"
	value := &po.YuCategory{}
	if err := common.DB.Get(value, sql, name, id); err != nil {
		common.Ylog.Debug(err)
		return nil
	}
	return value
}

func (d *dao) GetCategoryById(id int) *po.YuCategory {
	sql := "SELECT * FROM yu_category WHERE id = ? limit 1"
	value := &po.YuCategory{}
	if err := common.DB.Get(value, sql, id); err != nil {
		common.Ylog.Debug(err)
		return nil
	}
	return value
}

// update category
func (d *dao) UpdateCategory(id, status int, name string) error {
	sql := "UPDATE yu_category SET is_delete=?,`name` = ?,updated_at = NOW() WHERE id = ?"
	if _, err := common.DB.Exec(sql, status, name, id); err != nil {
		common.Ylog.Debug(err)
		return err
	}
	return nil
}

// delete by id
func (d *dao) DeleteCategoryById(id int) error {
	sql := "DELETE FROM yu_category WHERE id = ?"
	if _, err := common.DB.Exec(sql, id); err != nil {
		common.Ylog.Debug(err)
		return err
	}
	return nil
}

// add category
func (d *dao) InsertCategory(name string) error {
	sql := "INSERT INTO yu_category(`name`) VALUES(?)"
	if _, err := common.DB.Exec(sql, name); err != nil {
		common.Ylog.Debug(err)
		return err
	}
	return nil
}
