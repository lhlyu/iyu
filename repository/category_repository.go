package repository

import (
	"github.com/lhlyu/iyu/common"
	"github.com/lhlyu/iyu/errcode"
	"github.com/lhlyu/iyu/repository/po"
)

/**
添加一个category 【事务】
1. 判断是否已经存在
2. 如果已被删除，重新改为启用
3. 如果不存在，添加
*/
func (*dao) AddCategoryOne(category *model.YuCategory) *repositoryError {
	tx, err := common.DB.Beginx()
	if err != nil {
		return NewRepositoryError("AddCategoryOne", "", errcode.ERROR, err)
	}
	defer tx.Commit()
	newCategorys := []*model.YuCategory{}
	sql := "select * from yu_category where name = ? limit 1"
	if err = tx.Select(&newCategorys, sql, category.Name); err != nil {
		return NewRepositoryError("AddCategoryOne", sql, errcode.ERROR, err)
	}
	if len(newCategorys) == 0 {
		// 不存在，添加
		sql = "insert into yu_category(name) values(?)"
		if _, err = tx.Exec(sql, category.Name); err != nil {
			rollerr := tx.Rollback()
			return NewRepositoryError("AddCategoryOne", sql, errcode.ERROR, err, rollerr)
		}
		return nil
	}
	// 存在
	newCategory := newCategorys[0]
	if newCategory.IsDelete == 1 {
		sql = "update yu_category set is_delete = 0,updated_at = now() where name = ?"
		if _, err = tx.Exec(sql, category.Name); err != nil {
			rollerr := tx.Rollback()
			return NewRepositoryError("AddCategoryOne", sql, errcode.ERROR, err, rollerr)
		}
		return nil
	}
	return NewRepositoryError("AddCategoryOne", "", errcode.EXISTS_DATA)
}

/**
修改一个category 【事务】
1. 判断名字是否存在
           - 存在 -> 是否已删除
                     - 已删除 -> 改成 未删除 -> 原删除
                     - 未删除 -> 返回错误 存在相同名字
           - 不存在 -> 修改
ps: sqlx的Get方法，不存在就报错，所以还是采用Select放到切片
*/
func (*dao) UpdateCategoryOne(category *model.YuCategory) *repositoryError {
	tx, err := common.DB.Beginx()
	if err != nil {
		return NewRepositoryError("UpdateCategoryOne", "", errcode.ERROR, err)
	}
	defer tx.Commit()
	newCategorys := []*model.YuCategory{}
	sql := "select * from yu_category where name = ? limit 1"
	if err = tx.Select(&newCategorys, sql, category.Name); err != nil {
		return NewRepositoryError("UpdateCategoryOne", sql, errcode.ERROR, err)
	}
	if len(newCategorys) == 0 {
		// 不存在，修改
		sql = "update yu_category set name = ?,is_delete = 0,updated_at = now() where id = ?"
		if _, err = tx.Exec(sql, category.Name, category.Id); err != nil {
			rollerr := tx.Rollback()
			return NewRepositoryError("UpdateCategoryOne", sql, errcode.ERROR, err, rollerr)
		}
		return nil
	}
	newCategory := newCategorys[0]
	if newCategory.IsDelete == 1 {
		sql = "update yu_category set is_delete = 0,updated_at = now() where name = ?"
		if _, err = tx.Exec(sql, category.Name); err != nil {
			rollerr := tx.Rollback()
			return NewRepositoryError("UpdateCategoryOne", sql, errcode.ERROR, err, rollerr)
		}
		sql = "update yu_category set is_delete = 1,updated_at = now() where id = ?"
		if _, err = tx.Exec(sql, category.Id); err != nil {
			rollerr := tx.Rollback()
			return NewRepositoryError("UpdateCategoryOne", sql, errcode.ERROR, err, rollerr)
		}
		return nil
	}
	return NewRepositoryError("UpdateCategoryOne", "", errcode.EXISTS_DATA)
}

// 查询所有category
func (*dao) QueryCategory() ([]*model.YuCategory, *repositoryError) {
	sql := "select * from yu_category where is_delete = 0 order by updated_at desc,created_at desc"
	categorys := []*model.YuCategory{}
	if err := common.DB.Select(&categorys, sql); err != nil {
		return nil, NewRepositoryError("QueryCategory", sql, errcode.ERROR, err)
	}
	if len(categorys) == 0 {
		return nil, NewRepositoryError("QueryCategory", sql, errcode.EMPTY_DATA, nil)
	}
	return categorys, nil
}
