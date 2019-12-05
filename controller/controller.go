package controller

import (
	"github.com/iris-contrib/middleware/jwt"
	"github.com/kataras/iris"
	"github.com/lhlyu/iyu/cache"
	"github.com/lhlyu/iyu/common"
	"github.com/lhlyu/iyu/controller/vo"
	"github.com/lhlyu/iyu/errcode"
	"github.com/lhlyu/iyu/service"
	"github.com/lhlyu/iyu/util"
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

/**
jwt 通用
iss: 签发者
sub: 面向的用户
aud: 接收方
exp: 过期时间
nbf: 生效时间
iat: 签发时间
jti: 唯一身份标识
*/

func (controller) getToken(user *common.XUser) string {
	itv := common.Cfg.GetInt("jwt.itv")
	if itv == 0 {
		itv = common.ITV
	}
	m := make(map[string]interface{})
	m[common.X_ID] = user.Id
	m[common.X_ROLE] = user.Role
	now := time.Now()
	m["iat"] = now.Unix()
	m["nbf"] = now.Unix()
	m["exp"] = now.Add(time.Second * time.Duration(itv)).Unix()
	m["iss"] = common.Cfg.GetString("author")
	token := jwt.NewTokenWithClaims(jwt.SigningMethodHS256, jwt.MapClaims(m))
	tokenString, _ := token.SignedString([]byte(common.Cfg.GetString("jwt.secret")))
	cache.NewCache().SetJwt(tokenString)
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

// 记录
func (c controller) Record(ctx iris.Context, BusinessId, BusinessKind, Action int) {
	user := c.GetUser(ctx)
	svc := service.NewRecordService()
	param := &vo.RecordParam{
		UserId:       user.Id,
		Ip:           user.Ip,
		BusinessId:   BusinessId,
		BusinessKind: BusinessKind,
		Action:       Action,
	}
	if Action == common.ACTION_CMNT {
		svc.Insert(param)
		return
	}
	key := util.Base64EncodeObj(param)
	// 限制
	cache.NewCache().Record(key, func() {
		svc.Insert(param)
	})
}

// 是否是管理员
func (c controller) IsAdmin(ctx iris.Context) bool {
	user := c.GetUser(ctx)
	return user.Role >= common.PERMISSION
}

func (c controller) IsAdminRouter(ctx iris.Context) bool {
	admin, ok := ctx.Values().Get(common.ADMIN).(bool)
	if !ok {
		return false
	}
	return admin
}

type Controller struct {
	userController
	articleController
	tagController
	nailController
	categoryController
	quantaController
}
