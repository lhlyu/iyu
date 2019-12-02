package service

import (
	"github.com/lhlyu/iyu/cache"
	"github.com/lhlyu/iyu/common"
	"github.com/lhlyu/iyu/controller/vo"
	"github.com/lhlyu/iyu/errcode"
	"github.com/lhlyu/iyu/repository"
	"github.com/lhlyu/iyu/repository/po"
	"github.com/lhlyu/iyu/service/bo"
	"github.com/lhlyu/iyu/util"
	"sync"
)

type articleService struct {
}

func NewArticleService() *articleService {
	return &articleService{}
}

// query articles
func (*articleService) QueryArticles(param *vo.ArticleParam) *errcode.ErrCode {
	dao := repository.NewDao()
	total, err := dao.GetArticleCount(param)
	if err != nil {
		return errcode.QueryError
	}
	if total == 0 {
		return errcode.EmptyData
	}
	page := common.NewPage(param.PageNum, param.PageSize)
	page.SetTotal(total)
	ids, err := dao.QueryArticles(param, page)
	if err != nil {
		return errcode.QueryError
	}
	ch := cache.NewCache()
	articles := ch.GetArticles(ids...)
	if len(articles) > 0 {
		return errcode.Success.WithData(articles)
	}
	return errcode.EmptyData
}

// get article by id
func (*articleService) GetById(id int, reload bool) *errcode.ErrCode {
	// read from cache
	var (
		article    *po.YuArticle
		statArr    [5]int
		tags       []*bo.Tag
		nail       *bo.Nail
		category   *bo.Category
		articleErr error
	)

	ch := cache.NewCache()
	if !reload {
		articles := ch.GetArticles(id)
		if len(articles) > 0 {
			return errcode.Success.WithData(articles[0])
		}
	}
	dao := repository.NewDao()
	wg := sync.WaitGroup{}
	wg.Add(3)
	go func() {
		if article, articleErr = dao.GetArticle(id); articleErr == nil {
			che := cache.NewCache()
			categorys := che.GetCategory(article.CategoryId)
			if len(categorys) > 0 {
				category = categorys[0]
			}
			nails := che.GetNail(article.NailId)
			if len(nails) > 0 {
				nail = nails[0]
			}
		}
		wg.Done()
	}()
	go func() {
		if articleTags, err := dao.GetArticleTags(id); err == nil {
			var arr []int
			for _, v := range articleTags {
				arr = append(arr, v.TagId)
			}
			tags = cache.NewCache().GetTag(arr...)
		}
		wg.Done()
	}()
	go func() {
		if stats, err := dao.GetArticleStat(id); err == nil {
			for _, v := range stats {
				switch v.Action {
				case bo.FIRE:
					statArr[bo.FIRE] = v.Number
				case bo.CMNT:
					statArr[bo.CMNT] = v.Number
				case bo.LIKE:
					statArr[bo.LIKE] = v.Number
				case bo.UNLIKE:
					statArr[bo.UNLIKE] = v.Number
				}
			}
		}
		wg.Done()
	}()
	wg.Wait()
	if articleErr != nil {
		return errcode.QueryError
	}
	if article == nil {
		return errcode.NoExsistData
	}
	updatedAt := int(article.UpdatedAt.Unix())
	if updatedAt < 0 {
		updatedAt = int(article.CreatedAt.Unix())
	}
	articleData := &bo.ArticleData{
		ID:        article.Id,
		Kind:      article.Kind,
		CreatedAt: int(article.CreatedAt.Unix()),
		UpdatedAt: updatedAt,
		Title:     article.Title,
		Content:   util.Base64DecodeString(article.Content),
		Wraper:    article.Wraper,
		Tags:      tags,
		Nail:      nail,
		Category:  category,
		Fire:      statArr[bo.FIRE],
		CmntNum:   statArr[bo.CMNT],
		Like:      statArr[bo.LIKE],
		UnLike:    statArr[bo.UNLIKE],
	}
	ch.LoadArticle(articleData)
	return errcode.Success.WithData(articleData)
}

// insert one article
func (s *articleService) Insert(param *vo.ArticleVo) *errcode.ErrCode {
	article := &po.YuArticle{
		UserId:     param.UserId,
		Wraper:     param.Wraper,
		Title:      param.Title,
		Content:    util.Base64EncodeObj(param.Content),
		IsTop:      param.IsTop,
		IsDelete:   param.IsDelete,
		CategoryId: param.CategoryId,
		NailId:     param.NailId,
		Kind:       param.Kind,
	}
	if err := repository.NewDao().InsertArticle(article, param.TagArr); err != nil {
		return errcode.InsertError
	}
	// load map
	go s.GetById(article.Id, true)
	return errcode.Success
}

// update article
func (s *articleService) Update(param *vo.ArticleVo) *errcode.ErrCode {
	dao := repository.NewDao()
	article, err := dao.GetArticle(param.Id)
	if err != nil {
		return errcode.NoExsistData
	}
	param.Content = util.Base64EncodeObj(param.Content)
	util.CompareIntSet(&article.UserId, &param.UserId)
	util.CompareIntSet(&article.IsDelete, &param.IsDelete)
	util.CompareIntSet(&article.CategoryId, &param.CategoryId)
	util.CompareIntSet(&article.NailId, &param.NailId)
	util.CompareIntSet(&article.IsTop, &param.IsTop)
	util.CompareIntSet(&article.Kind, &param.Kind)
	util.CompareStrSet(&article.Title, &param.Title)
	util.CompareStrSet(&article.Content, &param.Content)
	util.CompareStrSet(&article.Wraper, &param.Wraper)

	NeedUpdateTag := false
	if len(param.TagArr) > 0 {
		if articleTags, err := dao.GetArticleTags(param.Id); err == nil {
			if len(articleTags) != len(param.TagArr) {
				NeedUpdateTag = true
			}
			if !NeedUpdateTag {
				for _, v := range articleTags {
					for _, w := range param.TagArr {
						if v.TagId != w {
							NeedUpdateTag = true
							break
						}
					}
				}
			}
		}

	}
	if NeedUpdateTag {
		err = dao.UpdateArticle(article, param.TagArr)
	} else {
		err = dao.UpdateArticle(article, nil)
	}
	if err != nil {
		return errcode.UpdateError
	}
	go s.GetById(param.Id, true)
	return errcode.Success
}

// batch delete article
func (s *articleService) Delete(param *vo.ArticleDeleteParam) *errcode.ErrCode {
	dao := repository.NewDao()
	if err := dao.DeleteArticle(param.Real, param.Ids...); err != nil {
		return errcode.DeleteError
	}
	if param.Real {
		cache.NewCache().DelArticles(param.Ids...)
	} else {
		go s.LoadArticles(param.Ids)
	}
	return errcode.Success
}

// load articles
func (s *articleService) LoadArticles(ids []int) {
	dao := repository.NewDao()
	ids, err := dao.QueryAllArticle(ids...)
	if err != nil {
		common.Ylog.Debug("load articles is err :", err)
		return
	}
	for _, v := range ids {
		s.GetById(v, true)
	}
}
