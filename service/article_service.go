package service

import (
	"github.com/lhlyu/iyu/common"
	"github.com/lhlyu/iyu/controller/dto"
	"github.com/lhlyu/iyu/result"
	"github.com/lhlyu/iyu/service/vo"
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

// article paging query
func (s *ArticleService) QueryPage(params *dto.Article, page *common.Page) ([]*vo.Category, *result.R) {

	return nil, nil
}
