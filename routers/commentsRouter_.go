package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["test_url/controllers:IndexController"] = append(beego.GlobalControllerRouter["test_url/controllers:IndexController"],
        beego.ControllerComments{
            Method: "Index",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["test_url/controllers:IndexController"] = append(beego.GlobalControllerRouter["test_url/controllers:IndexController"],
        beego.ControllerComments{
            Method: "IndexAbout",
            Router: `/about`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["test_url/controllers:IndexController"] = append(beego.GlobalControllerRouter["test_url/controllers:IndexController"],
        beego.ControllerComments{
            Method: "IndexMessage",
            Router: `/message`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["test_url/controllers:IndexController"] = append(beego.GlobalControllerRouter["test_url/controllers:IndexController"],
        beego.ControllerComments{
            Method: "IndexUser",
            Router: `/user`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
