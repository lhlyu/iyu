package service

import (
	"github.com/kataras/iris/v12"
	"github.com/lhlyu/iyu/common"
	"github.com/lhlyu/iyu/dao/po"
	"github.com/lhlyu/iyu/result"
	"github.com/lhlyu/iyu/trace"
)

type ArticleService struct {
	trace.BaseTracker
}

func NewArticleService(ctx iris.Context) *ArticleService {
	tracker := ctx.Values().Get(trace.TRACKER).(*trace.Tracker)
	return &ArticleService{
		BaseTracker: trace.NewBaseTracker(tracker),
	}
}

func (s *ArticleService) GetArticleByCode(code string) *result.R {
	article := &po.Article{}
	e := common.DB.Where("code = ?", code).First(article).Error
	if e != nil {
		s.Error(e, code)
		return result.NoExsistData
	}
	return result.Success.WithData(article)
}

func (s *ArticleService) QueryArticles() *result.R {
	article := &po.Article{}
	var items []*po.Article
	page := common.NewPage(1, 10)
	if e := article.Query(&items, page, nil, ""); e != nil {
		s.Error(e)
		return result.EmptyData
	}
	return result.Success.WithPage(page, items)
}

func (s *ArticleService) AddArticle() *result.R {
	article := &po.Article{}
	if e := article.Add(); e != nil {
		s.Error(e)
		return result.InsertError
	}
	return result.Success
}

func (s *ArticleService) UpdateArticle() *result.R {
	article := &po.Article{}
	if e := article.Update(nil); e != nil {
		s.Error(e)
		return result.UpdateError
	}
	return result.Success
}
