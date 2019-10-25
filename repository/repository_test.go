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

// test add one user
func TestDao_AddUserOne(t *testing.T) {
	d := NewDao()
	e := d.AddUserOne(&po.YuUser{
		ThirdId:   123,
		IsAdmin:   1,
		From:      2,
		Status:    1,
		AvatarUrl: "http://xx.png",
		UserUrl:   "http://xx.com",
		UserName:  "lhlyu",
		Bio:       "AASADADASDASDADA ASD A",
		Ip:        "0.0.0.0",
	})
	if e != nil {
		fmt.Println(e)
	}
}

// test query users
func TestDao_QueryUser(t *testing.T) {
	d := NewDao()
	page := &common.Page{
		PageNum:  1,
		PageSize: 10,
	}
	users, e := d.QueryUser(nil, page)
	if e != nil {
		fmt.Println(e)
		return
	}
	fmt.Printf("%+v\n", page)
	for _, u := range users {
		fmt.Println(u)
	}

}

// test get a user
func TestDao_GetUser(t *testing.T) {
	d := NewDao()
	user, e := d.GetUser(&po.YuUser{
		Id: 1,
	})
	if e != nil {
		fmt.Println(e)
		return
	}
	fmt.Printf("%+v\n", user)
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
