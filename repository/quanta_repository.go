package repository

import (
	"fmt"
	"github.com/lhlyu/iyu/common"
	"github.com/lhlyu/iyu/controller/vo"
	"github.com/lhlyu/iyu/repository/po"
)

func (d *dao) QueryQuanta(id ...int) []*po.YuQuanta {
	sql := "SELECT * FROM yu_quanta"
	var params []interface{}
	if len(id) > 0 {
		marks := d.createQuestionMarks(len(id))
		params = d.intConvertToInterface(id)
		sql += fmt.Sprintf(" where id in (%s)", marks)
	}
	sql += " ORDER BY is_enable,`key`"
	var values []*po.YuQuanta
	if err := common.DB.Select(&values, sql, params...); err != nil {
		d.Error(err)
		return nil
	}
	return values
}

func (d *dao) QueryQuantaCount() int {
	sql := "SELECT count(*) FROM yu_quanta ORDER BY is_enable,`key`"
	var value int
	if err := common.DB.Get(&value, sql); err != nil {
		d.Error(err)
		return 0
	}
	return value
}

func (d *dao) QueryQuantaPage(page *common.Page) []*po.YuQuanta {
	sql := "SELECT * FROM yu_quanta ORDER BY is_enable,`key` limit ?,?"
	var values []*po.YuQuanta
	if err := common.DB.Select(&values, sql, page.StartRow, page.PageSize); err != nil {
		d.Error(err)
		return nil
	}
	return values
}

// get quanta by key
func (d *dao) GetQuantaByKey(key ...string) []*po.YuQuanta {
	if len(key) == 0 {
		return nil
	}
	sql := "SELECT * FROM yu_quanta WHERE `key`in (%s) limit 1"
	marks := d.createQuestionMarks(len(key))
	sql = fmt.Sprintf(sql, marks)
	params := d.strConvertToInterface(key)
	var values []*po.YuQuanta
	if err := common.DB.Select(values, sql, params...); err != nil {
		d.Error(err)
		return nil
	}
	return values
}

func (d *dao) GetQuantaById(id int) *po.YuQuanta {
	sql := "SELECT * FROM yu_quanta WHERE id = ? limit 1"
	value := &po.YuQuanta{}
	if err := common.DB.Get(value, sql, id); err != nil {
		d.Error(err)
		return nil
	}
	return value
}

// update quanta
func (d *dao) UpdateQuanta(p *po.YuQuanta) error {
	sql := "UPDATE yu_quanta SET is_enable=?,`key` = ?,`value` = ?,`desc` = ?,updated_at = NOW() WHERE id = ?"
	if _, err := common.DB.Exec(sql, p.IsEnable, p.Key, p.Value, p.Desc, p.Id); err != nil {
		d.Error(err)
		return err
	}
	return nil
}

// delete by id
func (d *dao) DeleteQuantaById(id int) error {
	sql := "DELETE FROM yu_quanta WHERE id = ?"
	if _, err := common.DB.Exec(sql, id); err != nil {
		d.Error(err)
		return err
	}
	return nil
}

// add quanta
func (d *dao) InsertQuanta(p *vo.QuantaVo) (int, error) {
	sql := "INSERT INTO yu_quanta(`key`,`value`,`desc`,is_enable) VALUES(?,?,?,?)"
	result, err := common.DB.Exec(sql, p.Key, p.Value, p.Desc, p.IsEnable)
	if err != nil {
		d.Error(err)
		return 0, err
	}
	id, _ := result.LastInsertId()
	return int(id), nil
}
