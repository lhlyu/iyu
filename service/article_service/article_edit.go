package article_service

import (
	"github.com/lhlyu/iyu/controller/dto"
	"github.com/lhlyu/iyu/errcode"
	"github.com/lhlyu/iyu/repository/po"
	"github.com/lhlyu/yutil"
)

func (s *Service) AddArticle(param *dto.ArticleEditDto) *errcode.ErrCode {
	code := s.getUniqueCode()
	whr := &po.YuArticle{
		Code:       code,
		UserId:     param.UserId,
		Wrapper:    param.Wrapper,
		Title:      param.Title,
		Summary:    param.Summary,
		IsTop:      param.IsTop,
		CategoryId: param.CategoryId,
		Kind:       param.Kind,
		SortNum:    param.SortNum,
		CmntStatus: param.CmntStatus,
		IsDelete:   param.IsDelete,
	}
	whr.Content = yutil.Base64EncodeStrToStr(param.Content)
	id := s.dao.Add(whr, param.TagIds)
	if id == 0 {
		return errcode.InsertError
	}
	go s.GetArticleById(id)
	return errcode.Success
}
