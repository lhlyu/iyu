package cache

import (
	"github.com/lhlyu/iyu/common"
	"log"
	"strconv"
	"strings"
	"time"
)

/**
LHLYU-BLOG:AUTHOR  -  存放作者信息【string】
LHLYU-BLOG:CATALOG   -   分类【Hash】
LHLYU-BLOG:TAGS    -   标签【Hash】
LHLYU-BLOG:NAIL    -   钉子【Hash】


LHLYU-BLOG:ARTICLE:LIST - 文章列表【list】
-- LHLYU-BLOG:ARTICLE:LIST:KEY  【string】 =》 LHLYU-BLOG:ARTICLE:LIST

LHLYU-BLOG:ARTICLE:MAP - 文章MAP【hash】
- field: 文章ID

LHLYU-BLOG:ARTICLE:IVEAW:id  - 文章浏览量【string】

LHLYU-BLOG:IVEAW      - 全站浏览量【string】
-- LHLYU-BLOG:IVEAW:KEY  【string】 =》 LHLYU-BLOG:IVEAW
*/

const (
	_MAP  = ":MAP"
	_LIST = ":LIST"
)

const (
	_ONE_HOUR  = time.Hour
	_ONE_DAY   = _ONE_HOUR * 24
	_ONE_WEEK  = _ONE_DAY * 7
	_ONE_MONTH = _ONE_DAY * 30
)

type cache struct {
}

func NewCache() *cache {
	return &cache{}
}

func (c *cache) hasRedis() bool {
	if common.Redis == nil || c == nil {
		log.Println("redis is not initialized")
		return false
	}
	return true
}

func (c *cache) JoinSep(key ...string) string {
	return strings.Join(key, ":")
}

func (c *cache) getTimestamp() string {
	return ":" + strconv.FormatInt(time.Now().Unix(), 10)
}

// regex clear cache
func (c *cache) ClearCache(key string) {
	if c.hasRedis() {
		keys := common.Redis.Keys(key).Val()
		for _, k := range keys {
			common.Redis.Del(k)
		}
	}
}

func (c *cache) AddCatalogList(key string, v []interface{}) {
	if c.hasRedis() {
		common.Redis.RPush(key, v...)
	}
}

func (c *cache) mutexHandler(key string, f func()) {
	keyMutex := key + ":MUTEX"
	if rs, _ := common.Redis.SetNX(keyMutex, 1, time.Second*5).Result(); !rs {
		return
	}
	f()
	common.Redis.Del(keyMutex)
}
