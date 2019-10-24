package repository

import (
	"fmt"
	"github.com/lhlyu/iyu/common"
	"github.com/lhlyu/iyu/module"
	po "github.com/lhlyu/iyu/repository/po"
	"testing"
)

func init() {
	module.Register(module.CfgModule, module.DbModule)
	module.Init()
}

// 测试 添加一个nail
func TestDao_AddNailOne(t *testing.T) {
	d := NewDao()
	nail := &po.YuNail{
		Color: "#0f0f0f",
		Name:  "置顶",
	}
	e := d.AddNailOne(nail)
	fmt.Println(e)
}

// 测试 更新一个nail
func TestDao_UpdateNailOne(t *testing.T) {
	d := NewDao()
	nail := &po.YuNail{
		Id:    1,
		Color: "#0f0f0f",
		Name:  "TOP10",
	}
	e := d.UpdateNailOne(nail)
	fmt.Println(e)
}

// 测试 查询所有nail
func TestDao_QueryNailRepository(t *testing.T) {
	d := NewDao()
	nails, e := d.QueryNail()
	if e != nil {
		fmt.Println(e)
		return
	}
	for _, v := range nails {
		fmt.Println(v)
	}
}

func TestRepository(t *testing.T) {
	if common.DB == nil {
		return
	}
	d := NewDao()
	//d.AddNailOne("置顶", "#0000ff")
	a := d.convertToInterface([]int{1, 2, 3, 4, 5})
	fmt.Println(a...)
}

func BenchmarkSprintf(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {

	}
}
