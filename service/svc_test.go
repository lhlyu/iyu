package service

import (
	"github.com/lhlyu/iyu/controller/dto"
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
		module.RedisModule,
	)
	module.Init()
}

func TestCategoryService_AddOne(t *testing.T) {
	svc := NewCategoryService(trace.NewTracker())
	result := svc.AddOne(&dto.Category{
		Name: "vue",
	})
	t.Log(result.String())
}

func TestCategoryService_DelById(t *testing.T) {
	svc := NewCategoryService(trace.NewTracker())
	result := svc.DelById(&dto.Category{
		Id: uint(11),
	})
	t.Log(result.String())
}

func TestCategoryService_GetAll(t *testing.T) {
	svc := NewCategoryService(trace.NewTracker())
	datas, r := svc.GetAll()
	t.Log(yutil.Json.Marshal(datas))
	t.Log(r.String())
}

func TestCategoryService_QueryByName(t *testing.T) {
	svc := NewCategoryService(trace.NewTracker())
	datas, r := svc.QueryByName("J")
	t.Log(yutil.Json.Marshal(datas))
	t.Log(r.String())
}

func TestCategoryService_UpdateOne(t *testing.T) {
	svc := NewCategoryService(trace.NewTracker())
	r := svc.UpdateOne(&dto.Category{
		Id:   9,
		Name: "c++",
	})
	t.Log(r.String())
}
