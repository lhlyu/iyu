package module

import (
	"github.com/lhlyu/iyu/service"
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

	// clear all cache
	//che := cache.NewCache()
	//keys := che.JoinSep(common.Cfg.GetString("redis_key.iyu"), "*")
	//che.ClearCache(keys)
	// init data

	go loadCache()
}

func loadCache() {
	time.AfterFunc(time.Second*5, func() {
		log.Println("load nail datas ...")
		service.NewNailService().GetAll(true)
		log.Println("load category datas ...")
		service.NewCategoryService().GetAll(true)
		log.Println("load tag datas ...")
		service.NewTagService().GetAll(true)
		log.Println("load quanta datas ...")
		service.NewQuantaService().GetAll(nil, true)
		log.Println("load article datas ...")
		service.NewArticleService().LoadArticles(nil)
	})
}

var InitiateModule = initiate{}
