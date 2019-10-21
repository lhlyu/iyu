package main

import (
	"github.com/kataras/golog"
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/recover"
	"github.com/lhlyu/iyu/common"
	"github.com/lhlyu/iyu/module"
	"github.com/lhlyu/iyu/router"
)

func init() {
	module.Register(module.CfgModule, // 读取配置
		module.DbModule,    // 连接数据库
		module.TimerModule) // 启用定时任务
	module.Init()
}

func main() {
	app := iris.New()
	app.Use(recover.New())
	golog.SetLevel("debug")

	router.SetRouter(app)

	app.Run(iris.Addr(":" + common.Cfg.GetString("server.port")))
}
