package controllers

import (
	"github.com/astaxie/beego"
	"test_url/models"
)

const SESSION_USER_KEY = "SESSION_USER_KEY"

type BaseConroller struct {
	beego.Controller
	User    models.User
	IsLogin bool
}

func (ctx *BaseConroller) Prepare() {
	ctx.Data["Path"] = ctx.Ctx.Request.RequestURI
	//类型判断
	u, ok := ctx.GetSession(SESSION_USER_KEY).(models.User)
	ctx.IsLogin = false
	//如果判断是这个类型，
	if ok {
		ctx.User = u
		ctx.IsLogin = true
		ctx.Data["User"] = ctx.User
	}
}

func (ctx *BaseConroller) Abort500(err error) {
	ctx.Data["error"] = err
	ctx.Abort("500")
}
