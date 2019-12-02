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
	switch method {
	case "GET":
		if err := ctx.ReadQuery(v); err != nil {
			return errcode.IllegalParam
		}
	case "POST", "PUT", "DELETE":
		contentType := ctx.GetHeader("Content-Type")
		switch contentType {
		case "application/json":
			if err := ctx.ReadJSON(v); err != nil {
				return errcode.IllegalParam
			}
		case "application/x-www-form-urlencoded":
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

func (controller) getToken(user *common.XUser) string {
	m := make(map[string]interface{})
	m[common.X_ID] = user.Id
	m[common.X_ROLE] = user.Role
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

func (controller) GetUser(ctx iris.Context) *common.XUser {
	return ctx.Values().Get(common.X_USER).(*common.XUser)
}

type Controller struct {
	userController
	articleController
	tagController
	nailController
	categoryController
	quantaController
}
