package repository

import (
	"github.com/lhlyu/iyu/common"
	"github.com/lhlyu/iyu/errcode"
	"github.com/lhlyu/iyu/repository/po"
)

// 查询全部
func (*dao) QueryQuanta() ([]*po.YuQuanta, *repositoryError) {
	sql := "select * from yu_quanta"
	quantas := []*po.YuQuanta{}
	if err := common.DB.Select(&quantas, sql); err != nil {
		return nil, NewRepositoryError("QueryQuanta", sql, errcode.ERROR, err)
	}
	if len(quantas) == 0 {
		return nil, NewRepositoryError("QueryQuanta", sql, errcode.EMPTY_DATA, nil)
	}
	return quantas, nil
}

// 修改
func (*dao) UpdateQuantaOne(quanta *po.YuQuanta) *repositoryError {
	tx, err := common.DB.Beginx()
	if err != nil {
		return NewRepositoryError("UpdateQuantaOne", "", errcode.ERROR, err)
	}
	defer tx.Commit()
	newQuantas := []*po.YuQuanta{}
	sql := "select * from yu_quanta where id = ? limit 1"
	if err = tx.Select(&newQuantas, sql, quanta.Id); err != nil {
		return NewRepositoryError("UpdateQuantaOne", sql, errcode.ERROR, err)
	}
	if len(newQuantas) == 0 {
		// 不存在
		return NewRepositoryError("UpdateQuantaOne", sql, errcode.NO_EXISTS_DATA, err)
	}
	newQuanta := newQuantas[0]
	if quanta.Desc != "" && quanta.Desc != newQuanta.Desc {
		newQuanta.Desc = quanta.Desc
	}
	if quanta.Value != "" && quanta.Value != newQuanta.Value {
		newQuanta.Value = quanta.Value
	}
	if quanta.IsEnable > 0 && quanta.IsEnable != newQuanta.IsEnable {
		newQuanta.IsEnable = quanta.IsEnable
	}
	sql = "update yu_quanta set `value` = ?,`desc` = ?,is_enable = ?,updated_at = now() where id = ?"
	if _, err = tx.Exec(sql, newQuanta.Value, newQuanta.Desc, newQuanta.IsEnable, newQuanta.Id); err != nil {
		rollerr := tx.Rollback()
		return NewRepositoryError("UpdateQuantaOne", sql, errcode.ERROR, err, rollerr)
	}
	return nil

}
