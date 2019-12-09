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
	"strconv"
	"strings"
	"sync"
)

type articleService struct {
	*Service
}

func NewArticleService(traceId string) *articleService {
	return &articleService{
		Service: &Service{traceId},
	}
}

// query articles
func (s *articleService) QueryArticlePage(param *vo.ArticleParam) *errcode.ErrCode {
	dao := repository.NewDao(s.TraceId)
	total, err := dao.GetArticleCount(param)
	if err != nil {
		return errcode.QueryError
	}
	if total == 0 {
		return errcode.EmptyData
	}
	page := common.NewPage(param.PageNum, param.PageSize)
	page.SetTotal(total)
	ids, err := dao.QueryArticlePage(param, page)
	if err != nil {
		return errcode.QueryError
	}
	result := s.Query(false, ids...)
	if result.IsSuccess() {
		return errcode.Success.WithPage(page, result.Data)
	}
	return result
}

func (s *articleService) Query(reload bool, id ...int) *errcode.ErrCode {
	s.Info(reload, id)
	c := cache.NewCache(s.TraceId)
	var values []*bo.Article
	if !reload {
		values = c.GetArticle(id...)
	}
	if len(values) > 0 {
		return errcode.Success.WithData(values)
	}
	var (
		articles    []*po.YuArticle
		categoryIds []int
		nailIds     []int
		tagMap      map[int][]*bo.Tag
		statMap     map[int][]int
		articleErr  error
	)
	wg := sync.WaitGroup{}
	wg.Add(3)
	go func() {
		dao := repository.NewDao(s.TraceId)
		if articles, articleErr = dao.QueryArticle(id...); articleErr == nil {
			for _, v := range articles {
				categoryIds = append(categoryIds, v.CategoryId)
				nailIds = append(nailIds, v.NailId)
			}
		}
		wg.Done()
	}()
	go func() {
		tagMap = s.getTagMap(id...)
		wg.Done()
	}()
	go func() {
		statMap = s.getStatMap(id...)
		wg.Done()
	}()
	wg.Wait()
	if articleErr != nil {
		return errcode.QueryError
	}
	if len(articles) == 0 {
		return errcode.EmptyData
	}
	categoryMap := s.getCategoryMap(categoryIds...)
	nailMap := s.getNailMap(nailIds...)

	for _, v := range articles {
		updatedAt := v.UpdatedAt.Unix()
		if updatedAt < 0 {
			updatedAt = v.CreatedAt.Unix()
		}
		value := &bo.Article{
			Id:        v.Id,
			Kind:      v.Kind,
			CreatedAt: int(v.CreatedAt.Unix()),
			UpdatedAt: int(updatedAt),
			Title:     v.Title,
			Content:   util.Base64DecodeString(v.Content),
			Wraper:    v.Wraper,
			IsOpen:    v.IsOpen,
		}
		if d, ok := tagMap[v.Id]; ok {
			value.Tags = d
		}
		if d, ok := categoryMap[v.Id]; ok {
			value.Category = d
		}
		if d, ok := nailMap[v.Id]; ok {
			value.Nail = d
		}
		if d, ok := statMap[v.Id]; ok {
			value.Fire = d[bo.FIRE]
			value.CmntNum = d[bo.CMNT]
			value.Like = d[bo.LIKE]
			value.UnLike = d[bo.UNLIKE]
		}
		values = append(values, value)
	}
	go c.SetArticle(values...)
	return errcode.Success.WithData(values)
}

// add update
func (s *articleService) Edit(param *vo.ArticleVo) *errcode.ErrCode {
	dao := repository.NewDao(s.TraceId)
	param.Content = util.Base64EncodeObj(param.Content)
	if param.Id == 0 {
		id, err := dao.InsertArticle(param)
		if err != nil {
			return errcode.InsertError
		}
		go s.Query(true, id)
		return errcode.Success
	}
	data, err := dao.GetArticleById(param.Id)
	if err != nil {
		return errcode.QueryError
	}
	if data == nil {
		return errcode.NoExsistData
	}
	util.CompareIntSet(&data.UserId, &param.UserId)
	util.CompareIntSet(&data.IsDelete, &param.IsDelete)
	util.CompareIntSet(&data.IsOpen, &param.IsOpen)
	util.CompareIntSet(&data.CategoryId, &param.CategoryId)
	util.CompareIntSet(&data.NailId, &param.NailId)
	util.CompareIntSet(&data.IsTop, &param.IsTop)
	util.CompareIntSet(&data.Kind, &param.Kind)
	util.CompareStrSet(&data.Title, &param.Title)
	util.CompareStrSet(&data.Content, &param.Content)
	util.CompareStrSet(&data.Wraper, &param.Wraper)

	NeedUpdateTag := false
	if len(param.TagArr) > 0 {
		if articleTagArr, err := dao.GetArticleTags(param.Id); err == nil && len(articleTagArr) > 0 {
			articleTags := strings.Split(articleTagArr[0].TagIds, ",")
			if len(articleTags) != len(param.TagArr) {
				NeedUpdateTag = true
			}
			if !NeedUpdateTag {
				for _, v := range articleTags {
					for _, w := range param.TagArr {
						if v != strconv.Itoa(w) {
							NeedUpdateTag = true
							break
						}
					}
				}
			}
		}
	}
	if NeedUpdateTag {
		err = dao.UpdateArticle(data, param.TagArr)
	} else {
		err = dao.UpdateArticle(data, nil)
	}
	if err != nil {
		return errcode.UpdateError
	}
	go s.Query(true, data.Id)
	return errcode.Success
}

func (s *articleService) getTagMap(id ...int) map[int][]*bo.Tag {
	s.Info(id)
	dao := repository.NewDao(s.TraceId)
	values, err := dao.GetArticleTags(id...)
	if err != nil || len(values) == 0 {
		return nil
	}
	idMap := make(map[int]bool)
	m := make(map[int][]*bo.Tag)
	for _, v := range values {
		if v.TagIds != "" {
			var tags []*bo.Tag
			ids := strings.Split(v.TagIds, ",")
			for _, w := range ids {
				if id, err := strconv.Atoi(w); err == nil {
					tag := &bo.Tag{
						Id: id,
					}
					tags = append(tags, tag)
					idMap[id] = true
				}
			}
			m[v.ArticleId] = tags
		}
	}
	var ids []int
	for k := range idMap {
		ids = append(ids, k)
	}
	result := NewTagService(s.TraceId).Query(false, ids...)
	if !result.IsSuccess() {
		return nil
	}
	tagMap := make(map[int]*bo.Tag)
	for _, v := range result.Data.([]*bo.Tag) {
		tagMap[v.Id] = v
	}
	for _, v := range m {
		for _, w := range v {
			if value, has := tagMap[w.Id]; has {
				w.IsDelete = value.IsDelete
				w.Name = value.Name
			}
		}
	}
	return m
}

func (s *articleService) getStatMap(id ...int) map[int][]int {
	dao := repository.NewDao(s.TraceId)
	values, err := dao.GetArticleStat(id...)
	if err != nil || len(values) == 0 {
		return nil
	}
	length := len(bo.Stat)
	m := make(map[int][]int)
	for _, v := range values {
		arr, ok := m[v.BusinessId]
		if !ok {
			arr = make([]int, length)
		}
		arr[v.Action] = v.Number
		m[v.BusinessId] = arr
	}
	return m
}

func (s *articleService) getCategoryMap(id ...int) map[int]*bo.Category {
	svc := NewCategoryService(s.TraceId)
	result := svc.Query(false, id...)
	if !result.IsSuccess() {
		return nil
	}
	m := make(map[int]*bo.Category)
	for _, v := range result.Data.([]*bo.Category) {
		m[v.Id] = v
	}
	return m
}

func (s *articleService) getNailMap(id ...int) map[int]*bo.Nail {
	svc := NewNailService(s.TraceId)
	result := svc.Query(false, id...)
	if !result.IsSuccess() {
		return nil
	}
	m := make(map[int]*bo.Nail)
	for _, v := range result.Data.([]*bo.Nail) {
		m[v.Id] = v
	}
	return m
}
