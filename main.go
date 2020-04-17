package main

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/recover"
	"github.com/lhlyu/iyu/common"
	"github.com/lhlyu/iyu/middleware"
	"github.com/lhlyu/iyu/module"
	"github.com/lhlyu/iyu/router"
)

func init() {
	// 加载所需的模块
	module.Register(
		module.CfgModule, // 配置模块 <必须>
		module.LgModule,  // 日志模块
		//module.DbModule,       // DB模块
		module.InitiateModule, // 初始化模块
		//module.TimerModule,    // 定时任务模块
	)
	module.Init()
}

func main() {

	app := iris.New()

	// 前置中间件
	app.Use(middleware.Before())
	app.Use(recover.New())
	app.Use(middleware.Limiter()) // 限制每秒访问数量
	app.Use(middleware.Jwt())
	app.Use(middleware.Cors())
	app.Use(middleware.Log())

	// 后置 Post-Middleware
	app.Done(middleware.After())
	app.SetExecutionRules(iris.ExecutionRules{
		Done: iris.ExecutionOptions{Force: true},
	})
	router.SetRouter(app)
	app.Run(iris.Addr(common.Cfg.GetString("server.host") + ":" + common.Cfg.GetString("server.port")))
}
