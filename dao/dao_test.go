package dao

import (
	"github.com/lhlyu/iyu/common"
	"github.com/lhlyu/iyu/dao/po"
	"github.com/lhlyu/iyu/module"
	"github.com/lhlyu/iyu/trace"
	"testing"
)

func init() {
	module.Register(
		module.CfgModule, // 配置模块 <必须>
		module.DbModule,  // DB模块
	)
	module.Init()
}

func TestRecordDao_Add(t *testing.T) {
	v := &po.Record{
		UserId:   123,
		TargetId: 1,
		Kind:     3,
		Content:  "测试",
	}
	d := NewRecordDao(trace.NewTracker())
	d.Add(v)
}

func TestRecordDao_QueryVisit(t *testing.T) {
	d := NewRecordDao(trace.NewTracker())
	m := d.QueryVisit(0)
	for k, v := range m {
		t.Log(k, v)
	}
}

func TestRecordDao_Query(t *testing.T) {
	v := &po.Record{
		Kind: 3,
	}
	d := NewRecordDao(trace.NewTracker())
	items := d.Query(v, common.NewPage(1, 10))
	for _, item := range items {
		t.Logf("%+v\n", item)
	}
}
