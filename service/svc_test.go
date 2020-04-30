package service

import (
	"github.com/lhlyu/iyu/module"
)

func init() {
	// 加载所需的模块
	module.Register(
		module.CfgModule, // 配置模块 <必须>
		module.LgModule,  // 日志模块
		module.DbModule,  // DB模块
		module.RedisModule,
	)
	module.Init()
}
