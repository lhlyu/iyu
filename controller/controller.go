package controller

import (
	"github.com/iris-contrib/middleware/jwt"
	"github.com/kataras/iris"
	"github.com/lhlyu/iyu/common"
	"github.com/lhlyu/iyu/errcode"
	"gopkg.in/go-playground/validator.v9"
	"time"
)

var validate = validator.New()

type controller struct {
}

func (controller) getParams(ctx iris.Context, v interface{}, check bool) *errcode.ErrCode {
	// 根据方法获取参数
	// GET  -   query params
	// POST/PUT/DELETE  - body param
	method := ctx.Method()
	if method == "GET" {
		if err := ctx.ReadQuery(v); err != nil {
			return errcode.IllegalParam
		}
	} else if method == "POST" || method == "PUT" || method == "DELETE" {
		// application/json
		contentType := ctx.GetHeader("Content-Type")
		if contentType == "application/json" {
			if err := ctx.ReadJSON(v); err != nil {
				return errcode.IllegalParam
			}
		} else if contentType == "application/x-www-form-urlencoded" {
			if err := ctx.ReadForm(v); err != nil {
				return errcode.IllegalParam
			}
		}
	}
	if !check {
		return nil
	}
	if err := validate.Struct(v); err != nil {
		return errcode.IllegalParam
	}
	return nil
}

func (controller) getToken(m map[string]interface{}) string {
	if m == nil {
		m = make(map[string]interface{})
	}
	m["t"] = time.Now().Unix()
	token := jwt.NewTokenWithClaims(jwt.SigningMethodHS256, jwt.MapClaims(m))
	tokenString, _ := token.SignedString([]byte(common.Cfg.GetString("jwt.secret")))
	return tokenString
}

func (controller) checkUInt(v int) *errcode.ErrCode {
	if v <= 0 {
		return errcode.IllegalParam
	}
	return nil
}

func (controller) checkEmpty(v string) *errcode.ErrCode {
	if v == "" {
		return errcode.IllegalParam
	}
	return nil
}

type Controller struct {
	userController
	articleController
	tagController
	nailController
	categoryController
	quantaController
}
