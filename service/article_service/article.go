package article_service

import (
	"github.com/lhlyu/iyu/cache"
	"github.com/lhlyu/iyu/common"
	"github.com/lhlyu/iyu/controller/dto"
	"github.com/lhlyu/iyu/errcode"
	"github.com/lhlyu/iyu/repository/article_repository"
	"github.com/lhlyu/iyu/service/category_service"
	"github.com/lhlyu/iyu/service/tag_service"
	"github.com/lhlyu/iyu/service/user_service"
	"github.com/lhlyu/iyu/service/vo"
	"github.com/lhlyu/yutil"
	"sync"
	"time"
)

type Service struct {
	common.BaseService
	tagSvc      *tag_service.Service
	categorySvc *category_service.Service
	userSvc     *user_service.Service
	dao         *article_repository.Dao
	che         *cache.Cache
}

func NewService(traceId string) *Service {
	svc := &Service{}
	svc.tagSvc = tag_service.NewService(traceId)
	svc.categorySvc = category_service.NewService(traceId)
	svc.userSvc = user_service.NewService(traceId)
	svc.che = cache.NewCache(traceId)
	svc.dao = article_repository.NewDao(traceId)
	svc.SetTraceId(traceId)
	return svc
}

func (s *Service) QueryArticlePage(param *dto.ArticleDto) (*common.Page, []*vo.ArticleVo) {
	page := param.Page
	total := s.dao.Count(param)
	page.SetTotal(total)
	if total == 0 {
		return page, nil
	}
	datas := s.dao.QueryPage(param, page)
	var codes []string
	var ids []int
	categorySet := yutil.NewSet()
	userSet := yutil.NewSet()
	for _, v := range datas {
		codes = append(codes, v.Code)
		ids = append(ids, v.Id)
		categorySet.Add(v.CategoryId)
		userSet.Add(v.UserId)
	}
	items := s.che.GetArticle(codes...)
	if len(items) > 0 {
		return page, items
	}
	categoryIds := categorySet.ToIntArray()
	userIds := userSet.ToIntArray()
	var (
		categoryMap   map[int]*vo.CategoryVo
		userMap       map[int]*vo.UserVo
		articleTagMap map[int][]*vo.TagVo
	)
	var wg sync.WaitGroup
	wg.Add(3)
	go func() {
		categoryMap = s.getCategoryMap(categoryIds)
		wg.Done()
	}()
	go func() {
		userMap = s.getUserMap(userIds)
		wg.Done()
	}()
	go func() {
		articleTagMap = s.getArticleTagMap(ids)
		wg.Done()
	}()
	wg.Wait()
	for _, v := range datas {
		item := &vo.ArticleVo{
			Id:       v.Id,
			Kind:     v.Kind,
			IsDelete: v.IsDelete,
			SortNum:  v.SortNum,
		}
		item.Code = v.Code
		item.Wrapper = v.Wrapper
		item.Title = v.Title
		item.Summmary = v.Summary
		item.Content = yutil.Base64DecodeStrToStr(v.Content)
		item.IsTop = v.IsTop
		item.CmntStatus = v.CmntStatus
		item.CreatedAt = v.CreatedAt.Unix()
		item.UpdateAt = v.UpdatedAt.Unix()
		if item.UpdateAt <= 0 {
			item.UpdateAt = item.CreatedAt
		}
		stat := &vo.ArticleStat{
			Fire:    v.Fire,
			Like:    v.Like,
			CmntNum: v.CmntNum,
		}
		item.Stat = stat
		if w, ok := categoryMap[v.CategoryId]; ok {
			data := &vo.ArticleCategory{
				ID:    w.Id,
				Name:  w.Name,
				Color: w.Color,
			}
			item.Category = data
		}
		if w, ok := userMap[v.UserId]; ok {
			data := &vo.ArticleAuthor{
				ID:     w.Id,
				Name:   w.UserName,
				Avatar: w.AvatarUrl,
			}
			item.Author = data
		}
		if w, ok := articleTagMap[v.Id]; ok {
			var data []*vo.ArticleTags
			for _, v := range w {
				data = append(data, &vo.ArticleTags{
					ID:   v.Id,
					Name: v.Name,
				})
			}
			item.Tags = data
		}
		items = append(items, item)
	}
	go s.che.SetArticle(items)
	return page, items
}

func (s *Service) GetArticleByCode(code string) *errcode.ErrCode {
	page := common.NewPageOne()
	param := &dto.ArticleDto{
		Page:     page,
		Code:     code,
		IsDelete: 1,
	}
	_, data := s.QueryArticlePage(param)
	if len(data) == 0 {
		return errcode.NoExsistData
	}
	return errcode.Success.WithData(data[0].ArticleData)
}

func (s *Service) GetArticleById(id int) *errcode.ErrCode {
	page := common.NewPageOne()
	param := &dto.ArticleDto{
		Page: page,
		Id:   id,
	}
	_, data := s.QueryArticlePage(param)
	if len(data) == 0 {
		return errcode.NoExsistData
	}
	return errcode.Success.WithData(data[0])
}

/*** private ***/

func (s *Service) getCategoryMap(categoryIds []int) map[int]*vo.CategoryVo {
	if len(categoryIds) == 0 {
		return nil
	}
	items := s.categorySvc.QueryCategoryByKeys(categoryIds...)
	m := make(map[int]*vo.CategoryVo)
	for _, v := range items {
		m[v.Id] = v
	}
	return m
}

func (s *Service) getUserMap(userIds []int) map[int]*vo.UserVo {
	if len(userIds) == 0 {
		return nil
	}
	items := s.userSvc.QueryUserById(userIds...)
	m := make(map[int]*vo.UserVo)
	for _, v := range items {
		m[v.Id] = v
	}
	return m
}

func (s *Service) getArticleTagMap(articles []int) map[int][]*vo.TagVo {
	if len(articles) == 0 {
		return nil
	}
	datas := s.dao.QueryTagsByArticleIds(articles...)
	if len(datas) == 0 {
		return nil
	}
	var tagIds []int
	for _, v := range datas {
		tagIds = append(tagIds, v.TagId)
	}
	tags := s.tagSvc.QueryTagByIds(tagIds...)
	tagMap := make(map[int]string)
	for _, v := range tags {
		tagMap[v.Id] = v.Name
	}
	m := make(map[int][]*vo.TagVo)
	for _, v := range datas {
		w, ok := m[v.ArticleId]
		if !ok {
			w = make([]*vo.TagVo, 0)
		}
		tag := &vo.TagVo{
			Id: v.TagId,
		}
		if u, ok := tagMap[v.TagId]; ok {
			tag.Name = u
		}
		w = append(w, tag)
		m[v.ArticleId] = w
	}
	return m
}

func (s *Service) getUniqueCode() string {
	for i := 0; i < 20; i++ {
		code := yutil.RandString(6)
		count := s.dao.GetCodeCount(code)
		if count == 0 {
			return code
		}
	}
	return time.Now().Format("02150405")
}
