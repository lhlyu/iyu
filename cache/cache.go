package cache

import (
	"github.com/lhlyu/iyu/common"
	"github.com/lhlyu/iyu/util"
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

func (c *cache) mutexHandler(key string, f func()) {
	keyMutex := key + ":MUTEX"
	if rs, _ := common.Redis.SetNX(keyMutex, 1, time.Second*5).Result(); !rs {
		return
	}
	f()
	common.Redis.Del(keyMutex)
}

// load data to map and list
func (c *cache) LoadMapAndList(keyName string, vm map[string]interface{}) {
	if c.hasRedis() {
		key := common.Cfg.GetString(keyName)
		if key == "" {
			return
		}
		mapKey := key + _MAP
		targetMapKey := mapKey + c.getTimestamp()
		listKey := key + _LIST
		targetListKey := listKey + c.getTimestamp()
		c.mutexHandler(mapKey, func() {
			var arr []interface{}
			for k, v := range vm {
				value := util.ObjToJsonStr(v)
				common.Redis.HSet(targetMapKey, k, value)
				arr = append(arr, value)
			}
			common.Redis.Expire(targetMapKey, _ONE_WEEK)
			oldTargetMapKey := common.Redis.Get(mapKey).Val()
			common.Redis.Set(mapKey, targetMapKey, _ONE_WEEK)
			if oldTargetMapKey != "" {
				common.Redis.Del(oldTargetMapKey)
			}
			if len(arr) > 0 {
				common.Redis.RPush(targetListKey, arr...)
			}
			oldTargetListKey := common.Redis.Get(listKey).Val()
			common.Redis.Set(listKey, targetListKey, _ONE_WEEK)
			common.Redis.Expire(targetListKey, _ONE_WEEK)
			if oldTargetListKey != "" {
				common.Redis.Del(oldTargetListKey)
			}
		})
	}
}

// get data from list
func (c *cache) GetListData(keyName string) []string {
	if c.hasRedis() {
		key := common.Cfg.GetString(keyName)
		if key == "" {
			return nil
		}
		listKey := key + _LIST
		targetListKey := common.Redis.Get(listKey).Val()
		if targetListKey == "" {
			return nil
		}
		return common.Redis.LRange(targetListKey, 0, -1).Val()
	}
	return nil
}

// get data from list page
func (c *cache) GetListDataPage(keyName string, page *common.Page) []string {
	if c.hasRedis() {
		key := common.Cfg.GetString(keyName)
		if key == "" {
			return nil
		}
		listKey := key + _LIST
		targetListKey := common.Redis.Get(listKey).Val()
		if targetListKey == "" {
			return nil
		}
		total := common.Redis.LLen(targetListKey).Val()
		if total == 0 {
			return nil
		}
		page.SetTotal(int(total))
		return common.Redis.LRange(targetListKey, int64(page.StartRow), int64(page.StopRow)).Val()
	}
	return nil
}
