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

func NewDao(traceId string) *Dao {
	dao := &Dao{}
	dao.SetTraceId(traceId)
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
