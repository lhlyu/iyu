package module

import (
	"github.com/lhlyu/iyu/cache"
	"github.com/lhlyu/iyu/common"
	"github.com/lhlyu/iyu/service"
	"github.com/lhlyu/iyu/util"
	"log"
	"time"
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
	traceId := util.GetGID()
	che := cache.NewCache(traceId)
	keys := che.JoinSep(common.Cfg.GetString("redis_key.iyu"), "*")
	che.ClearCache(keys)

	go loadCache(traceId)
}

func loadCache(traceId string) {
	time.AfterFunc(time.Second*5, func() {
		log.Println("load nail datas ...")
		service.NewNailService(traceId).Query(true)
		log.Println("load category datas ...")
		service.NewCategoryService(traceId).Query(true)
		log.Println("load tag datas ...")
		service.NewTagService(traceId).Query(true)
		log.Println("load quanta datas ...")
		service.NewQuantaService(traceId).Query(true)
		log.Println("load article datas ...")
		service.NewArticleService(traceId).Query(true)
	})
}

var InitiateModule = initiate{}
