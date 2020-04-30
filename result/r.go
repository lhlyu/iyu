package result

import (
	"fmt"
	"github.com/lhlyu/iyu/common"
)

// 统一响应处理
type R struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func New(code int, msg string) *R {
	return &R{
		Code: code,
		Msg:  msg,
	}
}

func (r *R) IsSuccess() bool {
	if r.Code == 0 {
		return true
	}
	return false
}

func (r *R) String() string {
	return fmt.Sprintf("code=%d,msg=%s,data=%v", r.Code, r.Msg, r.Data)
}

func (r *R) WithData(data interface{}) *R {
	nr := New(r.Code, r.Msg)
	nr.Data = data
	return nr
}

func (r *R) WithPage(data interface{}, page *common.Page) *R {
	nr := New(r.Code, r.Msg)
	nr.Data = map[string]interface{}{
		"list": data,
		"page": page,
	}
	return nr
}

func (r *R) WithMsg(msg ...interface{}) *R {
	nr := New(r.Code, r.Msg)
	if nr.Msg == "" {
		nr.Msg = fmt.Sprint(msg...)
	} else {
		nr.Msg += ":" + fmt.Sprint(msg...)
	}
	return nr
}

var (
	Error   = New(-1, "系统异常")
	Success = New(0, "成功")
	Failure = New(1, "失败")

	EmptyData     = New(1000, "数据为空")
	ExistsData    = New(1001, "数据已存在")
	NotExistsData = New(1002, "数据不存在")
	IllegalParam  = New(1003, "参数不合法")
	CacheErr      = New(1004, "缓存读取异常")
	DbErr         = New(1005, "数据读取异常")

	QueryError  = New(10001, "查询失败")
	InsertError = New(10002, "插入失败")
	UpdateError = New(10003, "更新失败")
	DeleteError = New(10004, "删除失败")

	ProhibitDeleteError = New(100001, "分类下包含文章，禁止删除！")
)
