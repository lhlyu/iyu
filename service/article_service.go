package service

import (
	"github.com/lhlyu/iyu/controller/dto"
	"github.com/lhlyu/iyu/trace"
)

type ArticleService struct {
	BaseService
}

func NewArticleService(tracker trace.ITracker) *ArticleService {
	return &ArticleService{
		BaseService: NewBaseService(tracker),
	}
}

func (s *ArticleService) UpdateArticle(req *dto.Article) {

}
