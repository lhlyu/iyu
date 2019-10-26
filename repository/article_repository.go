package repository

import (
	"fmt"
	"github.com/lhlyu/iyu/common"
	"github.com/lhlyu/iyu/errcode"
	"github.com/lhlyu/iyu/repository/po"
	"strconv"
)

// add a article
func (*dao) AddArticleOne(article *po.YuArticle, labels []int) *repositoryError {
	tx, err := common.DB.Beginx()
	if err != nil {
		return NewRepositoryError("AddArticleOne", "", errcode.ERROR, err)
	}
	defer tx.Commit()
	sql := "INSERT INTO yu_article(user_id,bg,title,content,is_top,category_id,nail_id,kind) VALUE(?,?,?,?,?,?,?,?)"
	result, err := tx.Exec(sql, article.UserId, article.Bg, article.Title, article.Content, article.IsTop, article.CategoryId, article.NailId, article.Kind)
	if err != nil {
		tx.Rollback()
		return NewRepositoryError("AddArticleOne", sql, errcode.ERROR, err)
	}
	if len(labels) == 0 {
		return nil
	}
	article_id, _ := result.LastInsertId()
	sql = "INSERT INTO yu_article_label(article_id,label_id) VALUES(" + strconv.FormatInt(article_id, 10) + ",?)"
	params := []interface{}{labels[0]}
	for _, v := range labels[1:] {
		sql += fmt.Sprintf(",(%d,?)", article_id)
		params = append(params, v)
	}
	_, err = tx.Exec(sql, params...)
	if err != nil {
		tx.Rollback()
		return NewRepositoryError("AddArticleOne", sql, errcode.ERROR, params, err)
	}
	return nil
}

// update a article
func (*dao) UpdateArticleOne(article *po.YuArticle) *repositoryError {
	tx, err := common.DB.Beginx()
	if err != nil {
		return NewRepositoryError("UpdateUserOne", "", errcode.ERROR, err)
	}
	defer tx.Commit()
	newArticles := []*po.YuArticle{}
	sql := "select * from yu_article where id = ? limit 1"
	if err = tx.Select(&newArticles, sql, article.Id); err != nil {
		return NewRepositoryError("UpdateUserOne", sql, errcode.ERROR, err)
	}
	if len(newArticles) == 0 {
		// 不存在
		return NewRepositoryError("UpdateUserOne", sql, errcode.NO_EXISTS_DATA)
	}
	newArticle := newArticles[0]
	if article.Title != "" && article.Title != newArticle.Title {
		newArticle.Title = article.Title
	}
	if article.Content != "" && article.Content != newArticle.Content {
		newArticle.Content = article.Content
	}
	if article.Bg != "" && article.Bg != newArticle.Bg {
		newArticle.Bg = article.Bg
	}
	if article.IsTop > 0 && article.IsTop != newArticle.IsTop {
		newArticle.IsTop = article.IsTop
	}
	if article.CategoryId > 0 && article.CategoryId != newArticle.CategoryId {
		newArticle.CategoryId = article.CategoryId
	}
	if article.NailId > 0 && article.NailId != newArticle.NailId {
		newArticle.NailId = article.NailId
	}
	if article.Kind > 0 && article.Kind != newArticle.Kind {
		newArticle.Kind = article.Kind
	}
	if article.CommentsNumber > 0 && article.CommentsNumber != newArticle.CommentsNumber {
		newArticle.CommentsNumber = article.CommentsNumber
	}
	if article.View > 0 && article.View != newArticle.View {
		newArticle.View = article.View
	}
	if article.Like > 0 && article.Like != newArticle.Like {
		newArticle.Like = article.Like
	}
	if article.Unlike > 0 && article.Unlike != newArticle.Unlike {
		newArticle.Unlike = article.Unlike
	}
	if article.IsDelete > 0 && article.IsDelete != newArticle.IsDelete {
		newArticle.IsDelete = article.IsDelete
	}
	sql = "UPDATE yu_article SET user_id = ?, `like` = ?, unlike = ?, view = ?, comments_number = ?, bg = ?, title = ?, content = ?, is_top = ?, category_id = ?, nail_id = ?, kind = ?, is_delete = ?, updated_at = NOW() WHERE id = ?;"
	if _, err = tx.Exec(sql, newArticle.UserId, newArticle.Like, newArticle.Unlike, newArticle.View, newArticle.CommentsNumber, newArticle.Bg, newArticle.Title, newArticle.Content, newArticle.IsTop,
		newArticle.CategoryId, newArticle.NailId, newArticle.Kind, newArticle.IsDelete, newArticle.Id); err != nil {
		rollerr := tx.Rollback()
		return NewRepositoryError("UpdateUserOne", sql, errcode.ERROR, err, rollerr)
	}
	return nil
}

// conditional query article
func (*dao) QueryArticle(article *po.YuArticle, page *common.Page) ([]*po.YuArticle, *repositoryError) {
	countSql := "select count(*) from yu_article where 1 = 1 "
	var params []interface{}
	whstr := ""
	if article == nil {
		article = &po.YuArticle{}
	}
	if article.Id > 0 {
		whstr += "and id = ? "
		params = append(params, article.Id)
	}
	if article.Kind > 0 {
		whstr += "and kind = ? "
		params = append(params, article.Kind)
	}
	if article.NailId > 0 {
		whstr += "and nail = ? "
		params = append(params, article.NailId)
	}
	if article.CategoryId > 0 {
		whstr += "and category_id = ? "
		params = append(params, article.CategoryId)
	}
	if article.IsTop > 0 {
		whstr += "and is_top = ? "
		params = append(params, article.IsTop)
	}
	if article.Title != "" {
		whstr += "and title like ? "
		params = append(params, "%"+article.Title+"%")
	}
	var total int
	if err := common.DB.Get(&total, countSql+whstr, params...); err != nil {
		return nil, NewRepositoryError("QueryArticle", countSql+whstr, errcode.ERROR, params, err)
	}
	if total == 0 {
		return nil, NewRepositoryError("QueryArticle", "", errcode.NO_EXISTS_DATA)
	}
	page.SetTotal(total)
	sql := "select * from yu_article where 1 = 1 " + whstr + "order by is_top desc,updated_at desc,created_at desc limit ?,?"
	params = append(params, page.StartRow, page.PageSize)
	newArticles := []*po.YuArticle{}
	if err := common.DB.Select(&newArticles, sql, params...); err != nil {
		return nil, NewRepositoryError("QueryArticle", sql, errcode.ERROR, params, err)
	}
	if len(newArticles) == 0 {
		// 不存在
		return nil, NewRepositoryError("QueryArticle", "", errcode.NO_EXISTS_DATA)
	}
	return newArticles, nil
}

// get a article
func (*dao) GetArticleOne(article *po.YuArticle) (*po.YuArticle, *repositoryError) {
	sql := "select * from yu_article where id = ?"
	newArticles := []*po.YuArticle{}
	if err := common.DB.Select(&newArticles, sql, article.Id); err != nil {
		return nil, NewRepositoryError("GetUser", sql, errcode.ERROR, err)
	}
	if len(newArticles) == 0 {
		// 不存在
		return nil, NewRepositoryError("GetUser", sql, errcode.NO_EXISTS_DATA)
	}
	return newArticles[0], nil
}
