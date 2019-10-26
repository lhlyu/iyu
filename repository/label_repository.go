package repository

import (
	"fmt"
	"github.com/lhlyu/iyu/common"
	"github.com/lhlyu/iyu/errcode"
	"github.com/lhlyu/iyu/repository/po"
	"strconv"
)

/**
添加一个label 【事务】
1. 判断是否已经存在
2. 如果已被删除，重新改为启用
3. 如果不存在，添加
*/
func (*dao) AddLabelOne(label *po.YuLabel) *repositoryError {
	tx, err := common.DB.Beginx()
	if err != nil {
		return NewRepositoryError("AddLabelOne", "", errcode.ERROR, err)
	}
	defer tx.Commit()
	newLabels := []*po.YuLabel{}
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
	if newLabel.IsDelete == 2 {
		sql = "update yu_label set is_delete = 1,updated_at = now() where name = ?"
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
func (*dao) UpdateLabelOne(label *po.YuLabel) *repositoryError {
	tx, err := common.DB.Beginx()
	if err != nil {
		return NewRepositoryError("UpdateLabelOne", "", errcode.ERROR, err)
	}
	defer tx.Commit()
	newLabels := []*po.YuLabel{}
	sql := "select * from yu_label where name = ? limit 1"
	if err = tx.Select(&newLabels, sql, label.Name); err != nil {
		return NewRepositoryError("UpdateLabelOne", sql, errcode.ERROR, err)
	}
	if len(newLabels) == 0 {
		// 不存在，修改
		sql = "update yu_label set name = ?,is_delete = 1,updated_at = now() where id = ?"
		if _, err = tx.Exec(sql, label.Name, label.Id); err != nil {
			rollerr := tx.Rollback()
			return NewRepositoryError("UpdateLabelOne", sql, errcode.ERROR, err, rollerr)
		}
		return nil
	}
	newLabel := newLabels[0]
	if newLabel.IsDelete == 2 {
		sql = "update yu_label set is_delete = 1,updated_at = now() where name = ?"
		if _, err = tx.Exec(sql, label.Name); err != nil {
			rollerr := tx.Rollback()
			return NewRepositoryError("UpdateLabelOne", sql, errcode.ERROR, err, rollerr)
		}
		sql = "update yu_label set is_delete = 2,updated_at = now() where id = ?"
		if _, err = tx.Exec(sql, label.Id); err != nil {
			rollerr := tx.Rollback()
			return NewRepositoryError("UpdateLabelOne", sql, errcode.ERROR, err, rollerr)
		}
		return nil
	}
	return NewRepositoryError("UpdateLabelOne", "", errcode.EXISTS_DATA)
}

// 查询所有label
func (*dao) QueryLabel() ([]*po.YuLabel, *repositoryError) {
	sql := "select * from yu_label where is_delete = 1 order by updated_at desc,created_at desc"
	labels := []*po.YuLabel{}
	if err := common.DB.Select(&labels, sql); err != nil {
		return nil, NewRepositoryError("QueryLabel", sql, errcode.ERROR, err)
	}
	if len(labels) == 0 {
		return nil, NewRepositoryError("QueryLabel", sql, errcode.EMPTY_DATA, nil)
	}
	return labels, nil
}

// query all labels by article id
func (*dao) QueryLabelByArticleId(articleId int) ([]*po.YuArticleLabel, *repositoryError) {
	sql := "select * from yu_article_label where article_id = ? order by updated_at desc,created_at desc"
	labels := []*po.YuArticleLabel{}
	if err := common.DB.Select(&labels, sql, articleId); err != nil {
		return nil, NewRepositoryError("QueryLabelByArticleId", sql, errcode.ERROR, err)
	}
	if len(labels) == 0 {
		return nil, NewRepositoryError("QueryLabelByArticleId", sql, errcode.EMPTY_DATA, nil)
	}
	return labels, nil
}

// update article's labels
func (d *dao) UpdateArticleLabel(articleId int, labels []int) *repositoryError {
	articleArticles, err := d.QueryLabelByArticleId(articleId)
	if err != nil {
		return err
	}
	// add  update  delete
	tx, e := common.DB.Beginx()
	if e != nil {
		return NewRepositoryError("UpdateArticleLabel", "", errcode.ERROR, e)
	}
	defer tx.Commit()
	sql := "update yu_article_label set is_delete = 2 where article_id = ?"
	if _, e := tx.Exec(sql, articleId); err != nil {
		tx.Rollback()
		return NewRepositoryError("UpdateArticleLabel", sql, errcode.ERROR, e)
	}
	labelMap := make(map[int]bool)
	for _, v := range articleArticles {
		labelMap[v.LabelId] = true
	}
	var newLabels []int
	for _, v := range labels {
		if _, has := labelMap[v]; !has {
			newLabels = append(newLabels, v)
		}
	}
	if len(newLabels) > 0 {
		sql = "INSERT INTO yu_article_label(article_id,label_id) VALUES(" + strconv.Itoa(articleId) + ",?)"
		params := []interface{}{newLabels[0]}
		for _, v := range newLabels[1:] {
			sql += fmt.Sprintf(",(%d,?)", articleId)
			params = append(params, v)
		}
		_, e = tx.Exec(sql, params...)
		if e != nil {
			tx.Rollback()
			return NewRepositoryError("UpdateArticleLabel", sql, errcode.ERROR, params, e)
		}
	}
	sql = fmt.Sprintf("update yu_article_label set is_delete = 1 where is_delete = 2 and  article_id in (%s)", d.createQuestionMarks(len(labels)))
	_, e = tx.Exec(sql, d.intConvertToInterface(labels)...)
	if e != nil {
		tx.Rollback()
		return NewRepositoryError("UpdateArticleLabel", sql, errcode.ERROR, labels, e)
	}
	return nil
}
