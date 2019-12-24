package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/recover"
	"github.com/lhlyu/iyu/common"
	"github.com/lhlyu/iyu/middleware"
	"github.com/lhlyu/iyu/module"
	"github.com/lhlyu/iyu/router"
)

func init() {
	module.Register(module.CfgModule, // 读取配置 <必须>
		module.LgModule,       // 日志
		module.DbModule,       // 连接数据库
		module.RedisModule,    // redis
		module.InitiateModule, // 初始化
		module.TimerModule)    // 启用定时任务
	module.Init()
}

func main() {

	app := iris.New()

	// 前置
	app.Use(middleware.Before())
	app.Use(recover.New())
	app.Use(middleware.Limiter()) // 限制每秒访问数量
	app.Use(middleware.Log())
	app.Use(middleware.Cors())
	app.Use(middleware.Jwt())

	// 后置 Post-Middleware
	app.Done(middleware.After())
	app.SetExecutionRules(iris.ExecutionRules{
		Done: iris.ExecutionOptions{Force: true},
	})

	router.SetRouter(app)

	app.Run(iris.Addr(common.Cfg.GetString("server.host") + ":" + common.Cfg.GetString("server.port")))
}
