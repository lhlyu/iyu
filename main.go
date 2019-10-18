package main

import (
    "github.com/kataras/golog"
    "github.com/kataras/iris"
    "github.com/kataras/iris/middleware/logger"
    "github.com/kataras/iris/middleware/recover"
    "github.com/lhlyu/iyu/common"
    "github.com/lhlyu/iyu/module"
)

func init(){
    module.Register(module.CfgModule,module.DbModule)
    module.Init()
}

func main() {
    app := iris.New()
    app.Use(recover.New())
    app.Use(logger.New())
    golog.SetLevel("debug")
    app.Run(iris.Addr(":" + common.Cfg.GetString("server.port")))
}
