package repository

import (
	"fmt"
	"github.com/lhlyu/iyu/common"
	"github.com/lhlyu/iyu/controller/vo"
	"github.com/lhlyu/iyu/repository/po"
)

func (d *dao) QueryTag(id ...int) []*po.YuTag {
	sql := "SELECT * FROM yu_tag"
	var params []interface{}
	if len(id) > 0 {
		marks := d.createQuestionMarks(len(id))
		params = d.intConvertToInterface(id)
		sql += fmt.Sprintf(" where id in (%s)", marks)
	}
	sql += " ORDER BY is_delete"
	var values []*po.YuTag
	if err := common.DB.Select(&values, sql, params...); err != nil {
		d.Error(err)
		return nil
	}
	return values
}

// get tag by name
func (d *dao) GetTagByName(name string) *po.YuTag {
	sql := "SELECT * FROM yu_tag WHERE `name` = ? limit 1"
	value := &po.YuTag{}
	if err := common.DB.Get(value, sql, name); err != nil {
		d.Error(err)
		return nil
	}
	return value
}

func (d *dao) GetTagById(id int) *po.YuTag {
	sql := "SELECT * FROM yu_tag WHERE id = ? limit 1"
	value := &po.YuTag{}
	if err := common.DB.Get(value, sql, id); err != nil {
		d.Error(err)
		return nil
	}
	return value
}

// update tag
func (d *dao) UpdateTag(param *po.YuTag) error {
	sql := "UPDATE yu_tag SET is_delete=?,`name` = ?,updated_at = NOW() WHERE id = ?"
	if _, err := common.DB.Exec(sql, param.IsDelete, param.Name, param.Id); err != nil {
		d.Error(err)
		return err
	}
	return nil
}

// delete by id
func (d *dao) DeleteTagById(id int) error {
	sql := "DELETE FROM yu_tag WHERE id = ?"
	if _, err := common.DB.Exec(sql, id); err != nil {
		d.Error(err)
		return err
	}
	return nil
}

// add tag
func (d *dao) InsertTag(param *vo.TagVo) (int, error) {
	sql := "INSERT INTO yu_tag(`name`) VALUES(?)"
	result, err := common.DB.Exec(sql, param.Name)
	if err != nil {
		d.Error(err)
		return 0, err
	}
	id, _ := result.LastInsertId()
	return int(id), nil
}
