package cache

import (
	"github.com/lhlyu/iyu/common"
	"time"
)

type Cache struct {
	common.BaseCache
}

func NewCache(traceId string) *Cache {
	che := &Cache{}
	che.SetTraceId(traceId)
	return che
}

func (*Cache) getKey(name string) string {
	return common.Cfg.GetString(name)
}

func (c *Cache) mutexHandler(key string, f func()) {
	if common.Redis != nil {
		key += ":mutex"
		boolCmd := common.Redis.SetNX(key, "1", time.Second*10)
		if c.Error(boolCmd.Err(), key) {
			return
		}
		if !boolCmd.Val() {
			return
		}
		f()
		initCmd := common.Redis.Del(key)
		c.Error(initCmd.Err(), key)
		return
	}
}

func (c *Cache) ClearKeys(keys ...string) {
	if len(keys) == 0 {
		key := c.getKey(iyuKey) + "*"
		keys = common.Redis.Keys(key).Val()
	}
	if len(keys) > 0 {
		initCmd := common.Redis.Del(keys...)
		c.Error(initCmd.Err(), keys)
	}
}

const (
	tagKey          = "redis_key.tag"
	quantaKey       = "redis_key.quanta"
	categoryKey     = "redis_key.category"
	userKey         = "redis_key.user"
	globalVisitKey  = "redis_key.global_visit"
	articleVisitKey = "redis_key.article_visit"
	articleLikeKey  = "redis_key.article_like"
	cmntLikeKey     = "redis_key.cmnt_like"
	replyLikeKey    = "redis_key.reply_like"
	articleKey      = "redis_key.article"
	timelineKey     = "redis_key.timeline"
	tokenKey        = "redis_key.token"
	iyuKey          = "redis_key.iyu"
)

const (
	day   = time.Hour * 24
	week  = day * 7
	month = day * 30
)

/**
常用元素: 网站配置 分类 标签
type: hash
survival: week
named:
  iyu:quanta   key   value
  iyu:category id    value
  iyu:tag      id    value

用户:
type: hash
survival: day
named:
    iyu:user  userId   value

type: set
survival: day

named:
    iyu:global:visit  userId
    iyu:article:visit:{id}  userId
    iyu:article:like:{id}   userId
    iyu:cmnt:like:{id}      userId
    iyu:reply:like:{id}     userId

article

type: hash
survival: day

named:
    iyu:article  id  value


type : string  nx
token:
    iyu:token:{userId}  value

future:
网站访问量 文章访问量
*/
