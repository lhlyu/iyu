package repository

import (
	"fmt"
	"github.com/lhlyu/iyu/common"
	"github.com/lhlyu/iyu/controller/vo"
	"github.com/lhlyu/iyu/repository/po"
)

// 查询数量
func (d *dao) GetArticleCount(param *vo.ArticleParam) (int, error) {
	sql := `SELECT COUNT(DISTINCT a.id) FROM yu_article a LEFT JOIN yu_article_tag t ON t.article_id = a.id WHERE 1=1`
	var params []interface{}
	if param.CategoryId > 0 {
		sql += " AND category_id = ?"
		params = append(params, param.CategoryId)
	}
	if param.Kind > 0 {
		sql += " AND kind = ?"
		params = append(params, param.Kind)
	}
	if param.IsDelete > 0 {
		sql += " AND a.is_delete = ?"
		params = append(params, param.IsDelete)
	}
	if param.TagId > 0 {
		sql += " AND t.tag_id = ?"
		params = append(params, param.TagId)
	}
	if param.KeyWord != "" {
		sql += " AND title like ?"
		word := "%" + param.KeyWord + "%"
		params = append(params, word)
	}
	var total int
	if err := common.DB.Get(&total, sql, params...); err != nil {
		d.Error(err)
		return 0, err
	}
	return total, nil
}

// 查询
func (d *dao) QueryArticlePage(param *vo.ArticleParam, page *common.Page) ([]int, error) {
	sql := `SELECT DISTINCT a.id FROM yu_article a LEFT JOIN yu_article_tag t ON t.article_id = a.id WHERE 1=1`
	var params []interface{}
	if param.CategoryId > 0 {
		sql += " AND category_id = ?"
		params = append(params, param.CategoryId)
	}
	if param.Kind > 0 {
		sql += " AND kind = ?"
		params = append(params, param.Kind)
	}
	if param.IsDelete > 0 {
		sql += " AND a.is_delete = ?"
		params = append(params, param.IsDelete)
	}
	if param.TagId > 0 {
		sql += " AND t.tag_id = ?"
		params = append(params, param.TagId)
	}
	if param.KeyWord != "" {
		sql += " AND title like ?"
		word := "%" + param.KeyWord + "%"
		params = append(params, word)
	}
	sql += " ORDER BY is_top DESC,a.created_at DESC LIMIT ?,?"
	params = append(params, page.StartRow, page.PageSize)
	var result []int
	if err := common.DB.Select(&result, sql, params...); err != nil {
		d.Error(err)
		return nil, err
	}
	return result, nil
}

// 插入
func (d *dao) InsertArticle(article *vo.ArticleVo) (int, error) {
	sql1 := "INSERT INTO yu_article(user_id,wraper,title,content,is_top,category_id,nail_id,kind,is_open,is_delete,created_at,updated_at) VALUES(?,?,?,?,?,?,?,?,?,NOW(),NOW());"
	sql2 := "INSERT INTO yu_article_tag(article_id,tag_id)"
	tx, _ := common.DB.Beginx()
	rs, err := tx.Exec(sql1, article.UserId, article.Wraper, article.Title, article.Content, article.IsTop, article.CategoryId, article.NailId, article.Kind, article.IsOpen, article.IsDelete)
	if err != nil {
		d.Error(err)
		tx.Rollback()
		return 0, nil
	}
	id, err := rs.LastInsertId()
	if err != nil {
		d.Error(err)
		tx.Rollback()
		return 0, nil
	}
	if len(article.TagArr) == 0 {
		return 0, nil
	}
	article.Id = int(id)
	var valueArr [][]interface{}
	for _, v := range article.TagArr {
		var values []interface{}
		values = append(values, article.Id, v)
		valueArr = append(valueArr, values)
	}
	batchSql, params := d.createQuestionMarksForBatch(valueArr...)
	sql2 += batchSql
	_, err = tx.Exec(sql2, params...)
	if err != nil {
		d.Error(err)
		tx.Rollback()
		return 0, nil
	}
	if err = tx.Commit(); err != nil {
		d.Error(err)
		tx.Rollback()
		return 0, nil
	}
	return int(id), nil
}

// 获取单篇
func (d *dao) GetArticleById(id int) (*po.YuArticle, error) {
	sql := "select * from yu_article where id = ?"
	article := &po.YuArticle{}
	if err := common.DB.Get(article, sql, id); err != nil {
		d.Error(err)
		return nil, err
	}
	return article, nil
}

// 获取标签
func (d *dao) GetArticleTags(ids ...int) ([]*po.YuArticleTagV2, error) {
	d.Info(ids)
	sql := "SELECT article_id,GROUP_CONCAT(tag_id) as tags FROM yu_article_tag WHERE is_delete = 1"
	if len(ids) > 0 {
		sql += " AND article_id IN (%s)"
		marks := d.createQuestionMarks(len(ids))
		sql = fmt.Sprintf(sql, marks)
	}
	sql += " GROUP BY article_id"
	params := d.intConvertToInterface(ids)
	var articleTags []*po.YuArticleTagV2
	if err := common.DB.Select(&articleTags, sql, params...); err != nil {
		d.Error(err)
		return nil, err
	}
	return articleTags, nil
}

// 获取统计数据
func (d *dao) GetArticleStat(ids ...int) ([]*po.Stat, error) {
	sql := "SELECT business_id,`action`,COUNT(`action`) number FROM yu_record WHERE business_kind = 1"
	if len(ids) > 0 {
		sql += " AND business_id IN (%v)"
		marks := d.createQuestionMarks(len(ids))
		sql = fmt.Sprintf(sql, marks)
	}
	sql += " GROUP BY business_id,`action` ORDER BY business_id,`action`"
	params := d.intConvertToInterface(ids)
	var stats []*po.Stat
	if err := common.DB.Select(&stats, sql, params...); err != nil {
		d.Error(err)
		return nil, err
	}
	return stats, nil
}

// 更新
func (d *dao) UpdateArticle(article *po.YuArticle, articleTags []int) error {
	sql := "UPDATE yu_article SET user_id = ?,wraper = ?,title = ?,content = ?,is_top = ?,category_id = ?,nail_id = ?,kind = ?,is_open = ?,is_delete = ?,updated_at = NOW() WHERE id = ?"
	tx, _ := common.DB.Beginx()
	if _, err := tx.Exec(sql, article.UserId, article.Wraper, article.Title, article.Content, article.IsTop, article.CategoryId, article.NailId, article.Kind, article.IsOpen, article.IsDelete, article.Id); err != nil {
		d.Error(err)
		tx.Rollback()
		return nil
	}
	if len(articleTags) > 0 {
		sql = "UPDATE yu_article_tag SET is_delete = 2,updated_at = NOW() WHERE article_id = ?"
		if _, err := tx.Exec(sql, article.Id); err != nil {
			d.Error(err)
			tx.Rollback()
			return nil
		}
		sql = "INSERT INTO yu_article_tag(article_id,tag_id)"
		var valueArr [][]interface{}
		for _, v := range articleTags {
			var values []interface{}
			values = append(values, article.Id, v)
			valueArr = append(valueArr, values)
		}
		batchSql, params := d.createQuestionMarksForBatch(valueArr...)
		sql += batchSql
		if _, err := tx.Exec(sql, params...); err != nil {
			d.Error(err)
			tx.Rollback()
			return nil
		}
	}
	if err := tx.Commit(); err != nil {
		d.Error(err)
		tx.Rollback()
		return nil
	}
	return nil
}

// 删除
func (d *dao) DeleteArticle(real bool, ids ...int) error {
	sql := "update yu_article set is_delete = 2 where id in (%s)"
	if real {
		sql = "delete from yu_article where id in (%s)"
	}
	marks := d.createQuestionMarks(len(ids))
	params := d.intConvertToInterface(ids)
	sql = fmt.Sprintf(sql, marks)
	_, err := common.DB.Exec(sql, params...)
	if err != nil {
		d.Error(err)
		return err
	}
	return nil
}

func (d *dao) QueryArticle(ids ...int) ([]*po.YuArticle, error) {
	sql := `SELECT * FROM yu_article WHERE 1=1 `
	var params []interface{}
	if len(ids) > 0 {
		sql += " and id in (%s)"
		marks := d.createQuestionMarks(len(ids))
		params = d.intConvertToInterface(ids)
		sql = fmt.Sprintf(sql, marks)
	}
	var result []*po.YuArticle
	if err := common.DB.Select(&result, sql, params...); err != nil {
		d.Error(err)
		return nil, err
	}
	return result, nil
}
