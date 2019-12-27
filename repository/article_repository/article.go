package article_repository

import (
	"github.com/lhlyu/iyu/common"
	"github.com/lhlyu/iyu/controller/dto"
	"github.com/lhlyu/iyu/repository/po"
	"github.com/lhlyu/yutil"
)

type Dao struct {
	common.BaseDao
}

func NewDao(tracker *common.Tracker) *Dao {
	dao := &Dao{}
	dao.SetTracker(tracker)
	return dao
}

func (d *Dao) Count(whr *dto.ArticleDto) int {
	qb := &yutil.SqlBuffer{}
	qb.Add("select count(distinct a.id) from yu_article a left join yu_article_tag t on a.id = t.article_id where 1 = 1")
	qb.AddWhrGtZero(" and a.id = ?", whr.Id)
	qb.AddWhrGtZero(" and a.is_delete = ?", whr.IsDelete)
	qb.AddWhrGtZero(" and a.category_id = ?", whr.CategoryId)
	qb.AddWhrGtZero(" and a.kind = ?", whr.Kind)
	qb.AddWhrGtZero(" and t.tag_id = ?", whr.TagId)
	qb.AddWhrNeqEmpty(" and a.code = ?", whr.Code)
	if whr.KeyWord != "" {
		qb.AddWhr(" and title like ?", "%"+whr.KeyWord+"%")
	}
	sql, params := qb.Build()
	var value int
	if err := common.DB.Get(&value, sql, params...); err != nil {
		d.Error(err, sql, params)
		return 0
	}
	return value
}

func (d *Dao) QueryPage(whr *dto.ArticleDto, page *common.Page) []*po.YuArticle {
	qb := &yutil.SqlBuffer{}
	qb.Add("select a.* from yu_article a left join yu_article_tag t on a.id = t.article_id where 1 = 1")
	qb.AddWhrGtZero(" and a.id = ?", whr.Id)
	qb.AddWhrGtZero(" and a.is_delete = ?", whr.IsDelete)
	qb.AddWhrGtZero(" and a.category_id = ?", whr.CategoryId)
	qb.AddWhrGtZero(" and a.kind = ?", whr.Kind)
	qb.AddWhrGtZero(" and t.tag_id = ?", whr.TagId)
	qb.AddWhrNeqEmpty(" and a.code = ?", whr.Code)
	if whr.KeyWord != "" {
		qb.AddWhr(" and title like ?", "%"+whr.KeyWord+"%")
	}
	qb.Add(" group by a.id order by a.is_top desc,sort_num,created_at desc limit ?,?", page.StartRow, page.PageSize)
	sql, params := qb.Build()
	var values []*po.YuArticle
	if err := common.DB.Select(&values, sql, params...); err != nil {
		d.Error(err, sql, params)
		return nil
	}
	return values
}

func (d *Dao) QueryTagsByArticleIds(ids ...int) []*po.YuArticleTag {
	qb := &yutil.SqlBuffer{}
	qb.AddWhrIn("select * from yu_article_tag where article_id in (%s)", yutil.ConvertIntToInterface(ids)...)
	sql, params := qb.Build()
	var values []*po.YuArticleTag
	if err := common.DB.Select(&values, sql, params...); err != nil {
		d.Error(err, sql, params)
		return nil
	}
	return values
}

func (d *Dao) QueryArticleTimeline() []*po.YuArticle {
	sql := "select code,title,created_at from yu_article where is_delete = 1 order by created_at desc"
	var values []*po.YuArticle
	if err := common.DB.Select(&values, sql); err != nil {
		d.Error(err, sql)
		return nil
	}
	return values
}

func (d *Dao) GetCodeCount(code string) int {
	sql := "select count(*) from yu_article where code = ?"
	var value int
	if err := common.DB.Get(&value, sql, code); err != nil {
		d.Error(err, sql, code)
		return 0
	}
	return value
}

func (d *Dao) Add(whr *po.YuArticle, tags []int) int {
	tx, err := common.DB.Beginx()
	if err != nil {
		d.Error(err)
		return 0
	}
	sql := "insert into explore.yu_article (code,user_id,wrapper,title,summary,content,is_top,category_id,kind,sort_num,cmnt_status,is_delete,created_at,updated_at) values(?,?,?,?,?,?,?,?,?,?,?,?,now(),now())"
	rs, err := tx.Exec(sql, whr.Code, whr.UserId, whr.Wrapper, whr.Title, whr.Summary, whr.Content, whr.IsTop, whr.CategoryId, whr.Kind, whr.SortNum, whr.CmntStatus, whr.IsDelete)
	if err != nil {
		d.Error(err, sql, yutil.JsonObjToStr(whr))
		tx.Rollback()
		return 0
	}
	id, err := rs.LastInsertId()
	if err != nil {
		d.Error(err, sql, yutil.JsonObjToStr(whr))
		tx.Rollback()
		return 0
	}
	if len(tags) > 0 {
		var items [][]interface{}
		for _, v := range tags {
			var item []interface{}
			item = append(item, id, v)
			items = append(items, item)
		}
		str, params := yutil.CreateQuestionMarksForBatch(items...)
		if str == "" {
			return int(id)
		}
		sql = "insert into yu_article_tag(article_id,tag_id) values" + str
		_, err := tx.Exec(sql, params...)
		if err != nil {
			d.Error(err, sql, id, tags)
			tx.Rollback()
			return 0
		}
	}
	if err := tx.Commit(); err != nil {
		d.Error(err)
		tx.Rollback()
		return 0
	}
	return int(id)
}
