package service

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
	"github.com/lhlyu/iyu/module"
	"github.com/lhlyu/iyu/trace"
	"github.com/lhlyu/yutil/v2"
	"testing"
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

func TestArticleService_GetArticleByCode(t *testing.T) {
	svc := NewArticleService(ctx)
	t.Log(svc.GetArticleByCode("abcde"))
}

func TestArticleService_QueryArticles(t *testing.T) {
	svc := NewArticleService(ctx)
	t.Log(yutil.Json.Marshal(svc.QueryArticles()))
}
