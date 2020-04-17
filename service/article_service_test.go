package service

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
	"github.com/lhlyu/iyu/module"
	"github.com/lhlyu/iyu/trace"
)

var ctx iris.Context

func init() {
	// 加载所需的模块
	module.Register(
		module.CfgModule, // 配置模块 <必须>
		module.LgModule,  // 日志模块
		module.DbModule,  // DB模块
		//module.TimerModule,    // 定时任务模块
	)
	module.Init()
	ctx = context.NewContext(iris.Default())
	ctx.Values().Set(trace.TRACKER, trace.NewTracker())
}
