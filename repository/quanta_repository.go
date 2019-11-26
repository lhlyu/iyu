package repository

import (
	"github.com/lhlyu/iyu/common"
	"github.com/lhlyu/iyu/repository/po"
)

// get all quantas
func (d *dao) GetQuantaAll() []*po.YuQuanta {
	sql := "SELECT * FROM yu_quanta ORDER BY is_enable,`key`,updated_at DESC,created_at DESC"
	var values []*po.YuQuanta
	if err := common.DB.Select(&values, sql); err != nil {
		common.Ylog.Debug(err)
		return nil
	}
	return values
}

// get quanta by key
func (d *dao) GetQuantaByKey(id int, key string) *po.YuQuanta {
	sql := "SELECT * FROM yu_quanta WHERE `key` = ? and id <> ? limit 1"
	value := &po.YuQuanta{}
	if err := common.DB.Get(value, sql, key, id); err != nil {
		common.Ylog.Debug(err)
		return nil
	}
	return value
}

func (d *dao) GetQuantaById(id int) *po.YuQuanta {
	sql := "SELECT * FROM yu_quanta WHERE id = ? limit 1"
	value := &po.YuQuanta{}
	if err := common.DB.Get(value, sql, id); err != nil {
		common.Ylog.Debug(err)
		return nil
	}
	return value
}

// update quanta
func (d *dao) UpdateQuanta(p *po.YuQuanta) error {
	sql := "UPDATE yu_quanta SET is_enable=?,`key` = ?,`value` = ?,`desc` = ?,updated_at = NOW() WHERE id = ?"
	if _, err := common.DB.Exec(sql, p.IsEnable, p.Key, p.Value, p.Desc, p.Id); err != nil {
		common.Ylog.Debug(err)
		return err
	}
	return nil
}

// delete by id
func (d *dao) DeleteQuantaById(id int) error {
	sql := "DELETE FROM yu_quanta WHERE id = ?"
	if _, err := common.DB.Exec(sql, id); err != nil {
		common.Ylog.Debug(err)
		return err
	}
	return nil
}

// add quanta
func (d *dao) InsertQuanta(p *po.YuQuanta) error {
	sql := "INSERT INTO yu_quanta(`key`,`value`,`desc`,is_enable) VALUES(?,?,?,?)"
	if _, err := common.DB.Exec(sql, p.Key, p.Value, p.Desc, p.IsEnable); err != nil {
		common.Ylog.Debug(err)
		return err
	}
	return nil
}
