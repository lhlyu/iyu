package dao

import (
	"github.com/lhlyu/iyu/trace"
)

type ArticleDao struct {
	BaseDao
}

func NewArticleDao(tracker trace.ITracker) ArticleDao {
	return ArticleDao{
		BaseDao: NewBaseDao(tracker),
	}
}
