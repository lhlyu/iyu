package repository

import (
	"github.com/lhlyu/iyu/common"
	"github.com/lhlyu/iyu/errcode"
	"github.com/lhlyu/iyu/repository/po"
)

/**
添加一个nail 【事务】
1. 判断是否已经存在
2. 如果已被删除，重新改为启用
3. 如果不存在，添加
*/
func (*dao) AddNailOne(nail *model.YuNail) *repositoryError {
	tx, err := common.DB.Beginx()
	if err != nil {
		return NewRepositoryError("AddNailOne", "", errcode.ERROR, err)
	}
	defer tx.Commit()
	newNails := []*model.YuNail{}
	sql := "select * from yu_nail where name = ? limit 1"
	if err = tx.Select(&newNails, sql, nail.Name); err != nil {
		return NewRepositoryError("AddNailOne", sql, errcode.ERROR, err)
	}
	if len(newNails) == 0 {
		// 不存在，添加
		sql = "insert into yu_nail(name,color) values(?,?)"
		if _, err = tx.Exec(sql, nail.Name, nail.Color); err != nil {
			rollerr := tx.Rollback()
			return NewRepositoryError("AddNailOne", sql, errcode.ERROR, err, rollerr)
		}
		return nil
	}
	// 存在
	newNail := newNails[0]
	if newNail.IsDelete == 1 {
		sql = "update yu_nail set color = ?,is_delete = 0,updated_at = now() where name = ?"
		if _, err = tx.Exec(sql, nail.Color, nail.Name); err != nil {
			rollerr := tx.Rollback()
			return NewRepositoryError("AddNailOne", sql, errcode.ERROR, err, rollerr)
		}
		return nil
	}
	return NewRepositoryError("AddNailOne", "", errcode.EXISTS_DATA)
}

/**
修改一个nail 【事务】
1. 判断名字是否存在
           - 存在 -> 是否已删除
                     - 已删除 -> 改成 未删除 -> 原删除
                     - 未删除 -> 返回错误 存在相同名字
           - 不存在 -> 修改
ps: sqlx的Get方法，不存在就报错，所以还是采用Select放到切片
*/
func (*dao) UpdateNailOne(nail *model.YuNail) *repositoryError {
	tx, err := common.DB.Beginx()
	if err != nil {
		return NewRepositoryError("UpdateNailOne", "", errcode.ERROR, err)
	}
	defer tx.Commit()
	newNails := []*model.YuNail{}
	sql := "select * from yu_nail where name = ? limit 1"
	if err = tx.Select(&newNails, sql, nail.Name); err != nil {
		return NewRepositoryError("UpdateNailOne", sql, errcode.ERROR, err)
	}
	if len(newNails) == 0 {
		// 不存在，修改
		sql = "update yu_nail set name = ?,color = ?,is_delete = 0,updated_at = now() where id = ?"
		if _, err = tx.Exec(sql, nail.Name, nail.Color, nail.Id); err != nil {
			rollerr := tx.Rollback()
			return NewRepositoryError("UpdateNailOne", sql, errcode.ERROR, err, rollerr)
		}
		return nil
	}
	newNail := newNails[0]
	if newNail.IsDelete == 1 {
		sql = "update yu_nail set color = ?,is_delete = 0,updated_at = now() where name = ?"
		if _, err = tx.Exec(sql, nail.Color, nail.Name); err != nil {
			rollerr := tx.Rollback()
			return NewRepositoryError("UpdateNailOne", sql, errcode.ERROR, err, rollerr)
		}
		sql = "update yu_nail set is_delete = 1,updated_at = now() where id = ?"
		if _, err = tx.Exec(sql, nail.Id); err != nil {
			rollerr := tx.Rollback()
			return NewRepositoryError("UpdateNailOne", sql, errcode.ERROR, err, rollerr)
		}
		return nil
	}
	return NewRepositoryError("UpdateNailOne", "", errcode.EXISTS_DATA)
}

// 查询所有nail
func (*dao) QueryNail() ([]*model.YuNail, *repositoryError) {
	sql := "select * from yu_nail where is_delete = 0 order by updated_at desc,created_at desc"
	nails := []*model.YuNail{}
	if err := common.DB.Select(&nails, sql); err != nil {
		return nil, NewRepositoryError("QueryNail", sql, errcode.ERROR, err)
	}
	if len(nails) == 0 {
		return nil, NewRepositoryError("QueryNail", sql, errcode.EMPTY_DATA, nil)
	}
	return nails, nil
}
