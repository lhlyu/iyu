package module

import (
	"github.com/lhlyu/iyu/cache"
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
	che.ClearCache("*")
	// init data

}

var InitiateModule = initiate{}
