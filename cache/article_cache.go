package cache

import (
	"github.com/lhlyu/iyu/trace"
	"time"
)

type ArticleCache struct {
	BaseCache
}

func NewArticleCache(tracker trace.ITracker) ArticleCache {
	return ArticleCache{
		BaseCache: NewBaseCache(tracker),
	}
}

func (c ArticleCache) SetOne(code string, v interface{}) {
	key := c.Key("redis_key.article_prefix") + code
	c.BaseCache.Set(key, v, time.Second*600)
}

func (c ArticleCache) GetOne(code string, dist interface{}) (bool, error) {
	key := c.Key("redis_key.article_prefix") + code
	return c.BaseCache.Get(key, dist)
}

func (c ArticleCache) ClearOne(code string) {
	key := c.Key("redis_key.article_prefix") + code
	c.BaseCache.Clear(key)
}
