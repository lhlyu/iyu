package cache

import (
	"github.com/lhlyu/iyu/common"
	"log"
	"strconv"
	"strings"
	"time"
)

const (
	_MAP   = ":map"
	_LIST  = ":list"
	_MUTEX = ":mutex"
)

const (
	_ONE_HOUR  = time.Hour
	_ONE_DAY   = _ONE_HOUR * 24
	_ONE_WEEK  = _ONE_DAY * 7
	_ONE_MONTH = _ONE_DAY * 30
)

type cache struct {
	TraceId string
}

func NewCache(traceId string) *cache {
	return &cache{traceId}
}

func (s *cache) Error(param ...interface{}) {
	common.Ylog.Log(3, "error", s.TraceId, "cache", param...)
}

func (s *cache) Info(param ...interface{}) {
	common.Ylog.Log(3, "info", s.TraceId, "cache", param...)
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
		common.Redis.Del(keys...)
	}
}

func (c *cache) mutexHandler(key string, f func()) {
	keyMutex := key + _MUTEX
	if rs, _ := common.Redis.SetNX(keyMutex, 1, time.Second*5).Result(); !rs {
		return
	}
	f()
	common.Redis.Del(keyMutex)
}

func (c cache) setMap(keyName string, vm map[string]interface{}, duration time.Duration) {
	key := common.Cfg.GetString(keyName)
	if key == "" {
		return
	}
	mapKey := key + _MAP
	c.mutexHandler(mapKey, func() {
		common.Redis.HMSet(mapKey, vm)
		common.Redis.Expire(mapKey, duration)
	})
}

func (c cache) getMap(keyName string, field ...string) []string {
	key := common.Cfg.GetString(keyName)
	if key == "" {
		return nil
	}
	var arr []string
	mapKey := key + _MAP
	c.mutexHandler(mapKey, func() {
		if len(field) == 0 {
			m := common.Redis.HGetAll(mapKey).Val()
			for _, v := range m {
				arr = append(arr, v)
			}
		} else {
			m := common.Redis.HMGet(mapKey, field...).Val()
			for _, v := range m {
				s, ok := v.(string)
				if ok {
					arr = append(arr, s)
				}
			}
		}
	})
	return arr
}
