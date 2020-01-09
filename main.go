package main

import (
	_ "test_url/routers"
	_ "test_url/models" //这是只引用models 包的 init 函数
	"github.com/astaxie/beego"
	"strings"
)

func main() {
	initSession()
	initTemplate()
	beego.Run()
}
func initSession() {
	beego.BConfig.WebConfig.Session.SessionOn = true
	beego.BConfig.WebConfig.Session.SessionName = "liteblog-key"
	beego.BConfig.WebConfig.Session.SessionProvider = "file"
	beego.BConfig.WebConfig.Session.SessionProviderConfig = "data/session"
}

func initTemplate() {
	beego.AddFuncMap("equr1", func(x, y string) bool {
		//当前两个路径进行对比
		x1 := strings.Trim(x, "/")
		y1 := strings.Trim(y, "/")
		return strings.Compare(x1, y1) == 0
	})
}
