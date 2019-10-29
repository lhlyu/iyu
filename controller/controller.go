package controller

import (
	"github.com/kataras/iris"
	"github.com/lhlyu/iyu/errcode"
	validator "gopkg.in/go-playground/validator.v9"
)

var validate = validator.New()

type controller struct {
}

func (*controller) GetParams(ctx iris.Context, v interface{}, check bool) *errcode.ErrCode {
	// 根据方法获取参数
	// GET  -   query params
	// POST/PUT/DELETE  - body param
	method := ctx.Method()
	if method == "GET" {
		ctx.ReadQuery(v)
	} else if method == "POST" || method == "PUT" || method == "DELETE" {
		// application/json
		ctx.ReadJSON(v)
	}
	if !check {
		return nil
	}
	if err := validate.Struct(v); err != nil {
		return errcode.IllegalParam
	}
	return nil
}
