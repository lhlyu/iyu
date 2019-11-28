package repository

import (
	"github.com/lhlyu/iyu/common"
	"github.com/lhlyu/iyu/repository/po"
)

func (d *dao) InsertArticle(article *po.YuArticle, articleTags []int) error {
	sql1 := "INSERT INTO yu_article(user_id,wraper,title,content,is_top,category_id,nail_id,kind,is_delete,created_at,updated_at) VALUES(?,?,?,?,?,?,?,?,?,NOW(),NOW());"
	sql2 := "INSERT INTO yu_article_tag(article_id,tag_id)"
	tx, _ := common.DB.Beginx()
	defer tx.Commit()
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
	sql := "select * from yu_article_tag where article_id = ?"
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
	sql1 := "INSERT INTO yu_article(user_id,wraper,title,content,is_top,category_id,nail_id,kind,is_delete,created_at,updated_at) VALUES(?,?,?,?,?,?,?,?,?,NOW(),NOW());"
	sql2 := "INSERT INTO yu_article_tag(article_id,tag_id)"
	tx, _ := common.DB.Beginx()
	defer tx.Commit()
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
	return nil
}
