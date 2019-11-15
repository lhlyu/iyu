package service

import (
	"github.com/lhlyu/iyu/errcode"
	"github.com/lhlyu/iyu/service/bo"
)

type articleService struct {
}

func NewArticleService() *articleService {
	return &articleService{}
}

func (*articleService) GetArticles(param *bo.ArticleParam) *errcode.ErrCode {

	return nil
}
