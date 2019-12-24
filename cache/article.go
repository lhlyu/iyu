package cache

import (
	"github.com/lhlyu/iyu/common"
	"github.com/lhlyu/iyu/service/vo"
	"github.com/lhlyu/yutil"
)

func (c *Cache) GetArticle(fields ...string) []*vo.ArticleVo {
	key := c.getKey(articleKey)
	sliceCmd := common.Redis.HMGet(key, fields...)
	if c.Error(sliceCmd.Err(), key, fields) {
		return nil
	}
	vals := sliceCmd.Val()
	var items []*vo.ArticleVo
	for _, v := range vals {
		if v == nil {
			continue
		}
		item := &vo.ArticleVo{}
		yutil.JsonStrToObj(v.(string), item)
		items = append(items, item)
	}
	return items
}

func (c *Cache) SetArticle(items []*vo.ArticleVo) {
	key := c.getKey(articleKey)
	c.mutexHandler(key, func() {
		m := make(common.MSF)
		for _, v := range items {
			m[v.Code] = yutil.JsonObjToStr(v)
		}
		statusCmd := common.Redis.HMSet(key, m)
		if c.Error(statusCmd.Err(), key) {
			return
		}
		common.Redis.Expire(key, week)
	})
}

func (c *Cache) SetTimeline(items []*vo.ArticleTimeline) {
	key := c.getKey(timelineKey)
	c.mutexHandler(key, func() {
		value := yutil.JsonObjToStr(items)
		common.Redis.Set(key, value, month)
	})
}

func (c *Cache) GetTimeline() []*vo.ArticleTimeline {
	key := c.getKey(timelineKey)
	val := common.Redis.Get(key).Val()
	if val == "" {
		return nil
	}
	var items []*vo.ArticleTimeline
	yutil.JsonStrToObj(val, &items)
	return items
}
