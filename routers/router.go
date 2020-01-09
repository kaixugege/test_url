package routers

import (
	"github.com/astaxie/beego"
	"test_url/controllers"
)

func init() {
	//beego.Router("/", &controllers.MainController{})
	beego.Include(&controllers.IndexController{})
	beego.ErrorController(&controllers.ErrorController{})
}
