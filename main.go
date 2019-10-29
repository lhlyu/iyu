package main

import (
	"github.com/iris-contrib/middleware/cors"
	"github.com/kataras/golog"
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/recover"
	"github.com/lhlyu/iyu/common"
	"github.com/lhlyu/iyu/module"
	"github.com/lhlyu/iyu/router"
	"time"
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

	crs := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedHeaders: []string{"Content-Type"},
		AllowedMethods: []string{"GET", "POST", "PUT", "HEAD"},
		ExposedHeaders: []string{"X-Header"},
		MaxAge:         int((24 * time.Hour).Seconds()),
		// Debug:          true,
	})
	app.Use(crs)
	app.AllowMethods(iris.MethodOptions)
	// todo
	//var  j = jwt.New(jwt.Config{
	//        ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
	//            return []byte("My Secret"), nil
	//        },
	//        SigningMethod: jwt.SigningMethodHS256,
	//})

	app.OnErrorCode(iris.StatusNotFound, func(ctx iris.Context) {
		ctx.HTML("<b>Resource Not found</b>")
	})
	router.SetRouter(app)

	app.Run(iris.Addr("localhost:" + common.Cfg.GetString("server.port")))
}
