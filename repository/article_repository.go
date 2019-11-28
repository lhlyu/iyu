package repository

import (
	"github.com/lhlyu/iyu/common"
	"github.com/lhlyu/iyu/controller/vo"
	"github.com/lhlyu/iyu/repository/po"
)

func (d *dao) QueryArticleCount(param *vo.ArticleParam) (int, error) {
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
		sql += " AND is_delete = ?"
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
		common.Ylog.Debug(err)
		return 0, err
	}
	return total, nil
}

func (d *dao) QueryArticles(param *vo.ArticleParam) ([]int, error) {
	sql := `SELECT a.id FROM yu_article a LEFT JOIN yu_article_tag t ON t.article_id = a.id WHERE 1=1`
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
		sql += " AND is_delete = ?"
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
	sql += " GROUP BY a.id ORDER BY is_top DESC,a.created_at DESC"
	return nil, nil
}

func (d *dao) InsertArticle(article *po.YuArticle, articleTags []int) error {
	sql1 := "INSERT INTO yu_article(user_id,wraper,title,content,is_top,category_id,nail_id,kind,is_delete,created_at,updated_at) VALUES(?,?,?,?,?,?,?,?,?,NOW(),NOW());"
	sql2 := "INSERT INTO yu_article_tag(article_id,tag_id)"
	tx, _ := common.DB.Beginx()
	rs, err := tx.Exec(sql1, article.UserId, article.Wraper, article.Title, article.Content, article.IsTop, article.CategoryId, article.NailId, article.Kind, article.IsDelete, article.CreatedAt, article.UpdatedAt)
	if err != nil {
		common.Ylog.Debug(err)
		tx.Rollback()
		return nil
	}
	id, err := rs.LastInsertId()
	if err != nil {
		common.Ylog.Debug(err)
		tx.Rollback()
		return nil
	}
	if len(articleTags) == 0 {
		return nil
	}
	article.Id = int(id)
	var ids []interface{}
	var tags []interface{}
	for _, v := range articleTags {
		ids = append(ids, id)
		tags = append(tags, v)
	}
	batchSql, params := d.createQuestionMarksForBatch(ids, tags)
	sql2 += batchSql
	_, err = tx.Exec(sql2, params...)
	if err != nil {
		common.Ylog.Debug(err)
		tx.Rollback()
		return nil
	}
	if err = tx.Commit(); err != nil {
		common.Ylog.Debug(err)
		tx.Rollback()
		return nil
	}
	return nil
}

func (d *dao) GetArticle(id int) (*po.YuArticle, error) {
	sql := "select * from yu_article where id = ?"
	article := &po.YuArticle{}
	if err := common.DB.Get(article, sql, id); err != nil {
		common.Ylog.Debug(err)
		return nil, err
	}
	return article, nil
}

func (d *dao) GetArticleTags(id int) ([]*po.YuArticleTag, error) {
	sql := "select * from yu_article_tag where article_id = ? AND is_delete = 1"
	var articleTags []*po.YuArticleTag
	if err := common.DB.Select(&articleTags, sql, id); err != nil {
		common.Ylog.Debug(err)
		return nil, err
	}
	return articleTags, nil
}

func (d *dao) GetArticleStat(id int) ([]*po.Stat, error) {
	sql := "SELECT `action`,COUNT(`action`) number FROM yu_record  where business_id = ? and business_kind = 1 GROUP BY `action`"
	var stats []*po.Stat
	if err := common.DB.Select(&stats, sql, id); err != nil {
		common.Ylog.Debug(err)
		return nil, err
	}
	return stats, nil
}

// 更新
func (d *dao) UpdateArticle(article *po.YuArticle, articleTags []int) error {
	sql := "UPDATE yu_article SET user_id = ?,wraper = ?,title = ?,content = ?,is_top = ?,category_id = ?,nail_id = ?,kind = ?,is_delete = ?,updated_at = NOW() WHERE id = ?"
	tx, _ := common.DB.Beginx()
	if _, err := tx.Exec(sql, article.UserId, article.Wraper, article.Title, article.Content, article.IsTop, article.CategoryId, article.NailId, article.Kind, article.IsDelete, article.Id); err != nil {
		common.Ylog.Debug(err)
		tx.Rollback()
		return nil
	}
	if len(articleTags) > 0 {
		sql = "UPDATE yu_article_tag SET is_delete = 2,updated_at = NOW() WHERE article_id = ?"
		if _, err := tx.Exec(sql, article.Id); err != nil {
			common.Ylog.Debug(err)
			tx.Rollback()
			return nil
		}
		sql = "INSERT INTO yu_article_tag(article_id,tag_id)"
		var ids []interface{}
		var tags []interface{}
		for _, v := range articleTags {
			ids = append(ids, article.Id)
			tags = append(tags, v)
		}
		batchSql, params := d.createQuestionMarksForBatch(ids, tags)
		sql += batchSql
		if _, err := tx.Exec(sql, params...); err != nil {
			common.Ylog.Debug(err)
			tx.Rollback()
			return nil
		}
	}
	if err := tx.Commit(); err != nil {
		common.Ylog.Debug(err)
		tx.Rollback()
		return nil
	}
	return nil
}
