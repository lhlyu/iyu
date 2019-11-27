package repository

import (
	"github.com/lhlyu/iyu/repository/po"
)

func (d *dao) InsertArticle(article *po.YuArticle, articleTag *po.YuArticleTag) error {
	//sql1 := "INSERT INTO yu_article(user_id,wraper,title,content,is_top,category_id,nail_id,kind,is_delete,created_at,updated_at) VALUES(?,?,?,?,?,?,?,?,?,NOW(),NOW());"
	//sql2 := "INSERT INTO yu_article_tag(article_id,tag_id)"
	return nil
}
