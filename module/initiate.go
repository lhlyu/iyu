package module

import (
	"github.com/lhlyu/iyu/cache"
	"github.com/lhlyu/iyu/util"
	"github.com/lhlyu/yutil"
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
	// 工具包不忽略错误
	yutil.NotIgnore()
	// 初始化数据
	traceId := util.GetGID()
	// 清除所有缓存
	cache.NewCache(traceId).ClearKeys()

	go loadCache(traceId)
}

func loadCache(traceId string) {

}

var InitiateModule = initiate{}
