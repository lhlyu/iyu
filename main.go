package main

import (
	"github.com/kataras/golog"
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/pprof"
	"github.com/kataras/iris/middleware/recover"
	"github.com/lhlyu/iyu/common"
	"github.com/lhlyu/iyu/middleware"
	"github.com/lhlyu/iyu/module"
	"github.com/lhlyu/iyu/router"
)

func init() {
	module.Register(module.CfgModule, // 读取配置 <必须>
		module.DbModule,    // 连接数据库
		module.TimerModule) // 启用定时任务
	module.Init()
}

func main() {

	app := iris.New()

	golog.SetLevel("debug")

	app.Use(recover.New())
	app.Use(middleware.Log())
	app.Use(middleware.Cors())

	p := pprof.New()
	app.Any("/debug/pprof", p)
	app.Any("/debug/pprof/{action:path}", p)

	router.SetRouter(app)

	app.Run(iris.Addr(common.Cfg.GetString("server.host") + ":" + common.Cfg.GetString("server.port")))
}
