package repository

import (
	"fmt"
	"github.com/lhlyu/iyu/common"
	"github.com/lhlyu/iyu/controller/vo"
	"github.com/lhlyu/iyu/repository/po"
)

func (d *dao) QueryCategory(id ...int) []*po.YuCategory {
	sql := "SELECT * FROM yu_category"
	var params []interface{}
	if len(id) > 0 {
		marks := d.createQuestionMarks(len(id))
		params = d.intConvertToInterface(id)
		sql += fmt.Sprintf(" where id in (%s)", marks)
	}
	sql += " ORDER BY is_delete"
	var values []*po.YuCategory
	if err := common.DB.Select(&values, sql, params...); err != nil {
		common.Ylog.Debug(err)
		return nil
	}
	return values
}

// get category by name
func (d *dao) GetCategoryByName(name string) *po.YuCategory {
	sql := "SELECT * FROM yu_category WHERE `name` = ? limit 1"
	value := &po.YuCategory{}
	if err := common.DB.Get(value, sql, name); err != nil {
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
func (d *dao) UpdateCategory(param *po.YuCategory) error {
	sql := "UPDATE yu_category SET is_delete=?,`name` = ?,updated_at = NOW() WHERE id = ?"
	if _, err := common.DB.Exec(sql, param.IsDelete, param.Name, param.Id); err != nil {
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
func (d *dao) InsertCategory(param *vo.CategoryVo) (int, error) {
	sql := "INSERT INTO yu_category(`name`) VALUES(?)"
	result, err := common.DB.Exec(sql, param.Name)
	if err != nil {
		common.Ylog.Debug(err)
		return 0, err
	}
	id, _ := result.LastInsertId()
	return int(id), nil
}
