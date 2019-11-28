package service

import (
	"github.com/lhlyu/iyu/cache"
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

func (*articleService) GetArticles(param *bo.ArticleParam) *errcode.ErrCode {

	return nil
}

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
	//if !reload{
	//    articles := ch.GetArticles(id)
	//    if len(articles) > 0{
	//        return errcode.Success.WithData(articles[0])
	//    }
	//}
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
	articleData := &bo.ArticleData{
		ID:        article.Id,
		Kind:      article.Kind,
		CreatedAt: int(article.CreatedAt.Unix()),
		UpdatedAt: int(article.UpdatedAt.Unix()),
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

func (*articleService) Insert(param *vo.ArticleVo) *errcode.ErrCode {
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
	return errcode.Success
}

func (*articleService) Update(param *vo.ArticleVo) *errcode.ErrCode {
	//dao := repository.NewDao()
	//article,err := dao.GetArticle(param.Id)
	//if err != nil{
	//    return errcode.NoExsistData
	//}

	return nil
}
