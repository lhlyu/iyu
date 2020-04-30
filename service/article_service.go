package service

import (
	"github.com/lhlyu/iyu/controller/dto"
	"github.com/lhlyu/iyu/trace"
)

type ArticleService struct {
	trace.BaseTracker
}

func NewArticleService(tracker trace.ITracker) *ArticleService {
	return &ArticleService{
		BaseTracker: trace.NewBaseTracker(tracker),
	}
}

func (s *ArticleService) UpdateArticle(req *dto.Article) {

}
