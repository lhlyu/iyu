package repository

import (
	"github.com/lhlyu/iyu/common"
	"github.com/lhlyu/iyu/errcode"
	"github.com/lhlyu/iyu/repository/po"
)

/**
添加一个label 【事务】
1. 判断是否已经存在
2. 如果已被删除，重新改为启用
3. 如果不存在，添加
*/
func (*dao) AddLabelOne(label *model.YuLabel) *repositoryError {
	tx, err := common.DB.Beginx()
	if err != nil {
		return NewRepositoryError("AddLabelOne", "", errcode.ERROR, err)
	}
	defer tx.Commit()
	newLabels := []*model.YuLabel{}
	sql := "select * from yu_label where name = ? limit 1"
	if err = tx.Select(&newLabels, sql, label.Name); err != nil {
		return NewRepositoryError("AddLabelOne", sql, errcode.ERROR, err)
	}
	if len(newLabels) == 0 {
		// 不存在，添加
		sql = "insert into yu_label(name) values(?)"
		if _, err = tx.Exec(sql, label.Name); err != nil {
			rollerr := tx.Rollback()
			return NewRepositoryError("AddLabelOne", sql, errcode.ERROR, err, rollerr)
		}
		return nil
	}
	// 存在
	newLabel := newLabels[0]
	if newLabel.IsDelete == 1 {
		sql = "update yu_label set is_delete = 0,updated_at = now() where name = ?"
		if _, err = tx.Exec(sql, label.Name); err != nil {
			rollerr := tx.Rollback()
			return NewRepositoryError("AddLabelOne", sql, errcode.ERROR, err, rollerr)
		}
		return nil
	}
	return NewRepositoryError("AddLabelOne", "", errcode.EXISTS_DATA)
}

/**
修改一个label 【事务】
1. 判断名字是否存在
           - 存在 -> 是否已删除
                     - 已删除 -> 改成 未删除 -> 原删除
                     - 未删除 -> 返回错误 存在相同名字
           - 不存在 -> 修改
ps: sqlx的Get方法，不存在就报错，所以还是采用Select放到切片
*/
func (*dao) UpdateLabelOne(label *model.YuLabel) *repositoryError {
	tx, err := common.DB.Beginx()
	if err != nil {
		return NewRepositoryError("UpdateLabelOne", "", errcode.ERROR, err)
	}
	defer tx.Commit()
	newLabels := []*model.YuLabel{}
	sql := "select * from yu_label where name = ? limit 1"
	if err = tx.Select(&newLabels, sql, label.Name); err != nil {
		return NewRepositoryError("UpdateLabelOne", sql, errcode.ERROR, err)
	}
	if len(newLabels) == 0 {
		// 不存在，修改
		sql = "update yu_label set name = ?,is_delete = 0,updated_at = now() where id = ?"
		if _, err = tx.Exec(sql, label.Name, label.Id); err != nil {
			rollerr := tx.Rollback()
			return NewRepositoryError("UpdateLabelOne", sql, errcode.ERROR, err, rollerr)
		}
		return nil
	}
	newLabel := newLabels[0]
	if newLabel.IsDelete == 1 {
		sql = "update yu_label set is_delete = 0,updated_at = now() where name = ?"
		if _, err = tx.Exec(sql, label.Name); err != nil {
			rollerr := tx.Rollback()
			return NewRepositoryError("UpdateLabelOne", sql, errcode.ERROR, err, rollerr)
		}
		sql = "update yu_label set is_delete = 1,updated_at = now() where id = ?"
		if _, err = tx.Exec(sql, label.Id); err != nil {
			rollerr := tx.Rollback()
			return NewRepositoryError("UpdateLabelOne", sql, errcode.ERROR, err, rollerr)
		}
		return nil
	}
	return NewRepositoryError("UpdateLabelOne", "", errcode.EXISTS_DATA)
}

// 查询所有label
func (*dao) QueryLabel() ([]*model.YuLabel, *repositoryError) {
	sql := "select * from yu_label where is_delete = 0 order by updated_at desc,created_at desc"
	labels := []*model.YuLabel{}
	if err := common.DB.Select(&labels, sql); err != nil {
		return nil, NewRepositoryError("QueryLabel", sql, errcode.ERROR, err)
	}
	if len(labels) == 0 {
		return nil, NewRepositoryError("QueryLabel", sql, errcode.EMPTY_DATA, nil)
	}
	return labels, nil
}
