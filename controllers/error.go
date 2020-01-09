package controllers

import (
	"test_url/syserror"
	"github.com/astaxie/beego/logs"
)

type ErrorController struct {
	BaseConroller
}

// sjax {code,msg,reason_error  }
func (c *ErrorController) Error404() {

	c.TplName = "error/404.html"
	//如果是ajax请求，则返回一个json数据
	if c.IsAjax() {
		c.jsonerror(syserror.Error404{})
	} else {
		c.Data["content"] = "非法访问"
	}
}

func (c *ErrorController) Error500() {
	c.TplName = "error/500.html"
	err, ok := c.Data["error"].(error)
	// 如果不知道什么错误，
	if !ok {
		err = syserror.New("未知错误", nil)
	}
	// 因为可能err没有实现我们的系统错误，所以在这里判断一下Error类型
	serr, ok := err.(syserror.Error)
	// 如果不是系统Error类型,就生成一个系统的Error
	if !ok {
		serr = syserror.New(serr.Error(), nil)
	}
	// 输出错误原因
	if serr.ReasonError() != nil {
		logs.Info(serr.ReasonError().Error())
	}

	if c.IsAjax() {
		c.jsonerror(serr)
	} else {
		c.Data["constant"] = serr.Error()
	}

}

//一个封装的ajax请求的错误实现
func (c ErrorController) jsonerror(serr syserror.Error) {
	c.Ctx.Output.Status = 200 //返回的请求头
	c.Data["json"] = map[string]interface{}{
		"code":         serr.Code(),        //错误码
		"msg":          serr.Error(),       //错误信息
		"reason_error": serr.ReasonError(), //错误类型
	}
	c.ServeJSON()
}

//
//
//func (c *ErrorController) ErrorDb() {
//	c.Data["content"] = "database is now down"
//	c.TplName = "dberror.tpl"
//}
