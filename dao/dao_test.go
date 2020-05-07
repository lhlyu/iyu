package dao

import (
	"github.com/lhlyu/iyu/common"
	"github.com/lhlyu/iyu/dao/po"
	"github.com/lhlyu/iyu/module"
	"github.com/lhlyu/iyu/trace"
	"github.com/lhlyu/yutil/v2"
	"testing"
)

func init() {
	// 加载所需的模块
	module.Register(
		module.CfgModule, // 配置模块 <必须>
		module.LgModule,  // 日志模块
		module.DbModule,  // DB模块
	)
	module.Init()
}

func TestBaseDao_QueryPage(t *testing.T) {
	d := NewBaseDao(trace.NewTracker())
	var datas []*po.Category
	page := common.NewPage(1, 10)
	err := d.QueryPage(&datas, page, "name desc", "name like ?", "%j%")
	if err != nil {
		t.Error(err)
	}
	t.Log(yutil.Json.Marshal(datas))
	t.Log(yutil.Json.Marshal(page))
}
