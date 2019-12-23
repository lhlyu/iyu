package article_service

import (
	"github.com/lhlyu/iyu/common"
	"github.com/lhlyu/iyu/controller/dto"
	"github.com/lhlyu/iyu/errcode"
	"github.com/lhlyu/iyu/service/vo"
)

func (s *Service) QueryHomeArticlePage(param *dto.ArticleDto) *errcode.ErrCode {
	page, datas := s.QueryArticlePage(param)
	if datas == nil {
		return errcode.EmptyData
	}
	var items []*vo.ArticleData
	for _, v := range datas {
		item := v.ArticleData
		items = append(items, &item)
	}
	return errcode.Success.WithPage(page, items)
}

func (s *Service) GetAbout() *errcode.ErrCode {
	page := common.NewPageOne()
	param := &dto.ArticleDto{
		Page:     page,
		IsDelete: 1,
		Kind:     2,
	}
	_, data := s.QueryArticlePage(param)
	if len(data) == 0 {
		return errcode.NoExsistData
	}
	return errcode.Success.WithData(data[0].ArticleData)
}
