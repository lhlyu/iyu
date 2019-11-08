package module

import (
	"github.com/lhlyu/iyu/cache"
	"github.com/lhlyu/iyu/common"
	"log"
)

// 启动时执行
type initiate struct {
}

func (initiate) seq() int {
	return 1 << 4
}

func (initiate) SetUp() {
	log.Println("init initiate module ->")
	// 初始化数据

	// clear all cache
	che := cache.NewCache()
	keys := che.JoinSep(common.Cfg.GetString("redis_key.iyu"), "*")
	che.ClearCache(keys)
	// init data

}

var InitiateModule = initiate{}
