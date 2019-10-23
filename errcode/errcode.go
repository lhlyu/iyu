package errcode

import (
	"fmt"
)

type ErrCode struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func (e *ErrCode) IsSuccess() bool {
	if e.Code == 0 {
		return true
	}
	return false
}

func (e *ErrCode) String() string {
	return fmt.Sprintf("code=%d,msg=%s,data=%v", e.Code, e.Msg, e.Data)
}

func (e *ErrCode) GetErrCode() *ErrCode {
	return e
}

func NewErrcode(code int, data interface{}) *ErrCode {
	return &ErrCode{
		Code: code,
		Data: data,
		Msg:  errCodeMap[code],
	}
}

const (
	ERROR   = iota - 1 // -1
	SUCCESS            // 0
	FAILURE
)

const (
	EMPTY_DATA = iota + 1000
	EXISTS_DATA
	NO_EXISTS_DATA
)

// 常用
var (
	Error        = NewErrcode(ERROR, nil)
	Success      = NewErrcode(SUCCESS, nil)
	Failure      = NewErrcode(FAILURE, nil)
	EmptyData    = NewErrcode(EMPTY_DATA, nil)
	ExsistData   = NewErrcode(EXISTS_DATA, nil)
	NoExsistData = NewErrcode(NO_EXISTS_DATA, nil)
)

var errCodeMap = map[int]string{
	ERROR:          "异常",
	SUCCESS:        "成功",
	FAILURE:        "失败",
	EMPTY_DATA:     "数据为空",
	EXISTS_DATA:    "数据已存在",
	NO_EXISTS_DATA: "数据不存在",
}
