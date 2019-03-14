package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["usage-api/controllers:CheckController"] = append(beego.GlobalControllerRouter["usage-api/controllers:CheckController"],
        beego.ControllerComments{
            Method: "Get",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["usage-api/controllers:OwnerUsageController"] = append(beego.GlobalControllerRouter["usage-api/controllers:OwnerUsageController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"Post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["usage-api/controllers:OwnerUsageController"] = append(beego.GlobalControllerRouter["usage-api/controllers:OwnerUsageController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/`,
            AllowHTTPMethods: []string{"Put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["usage-api/controllers:OwnerUsageController"] = append(beego.GlobalControllerRouter["usage-api/controllers:OwnerUsageController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["usage-api/controllers:OwnerUsageController"] = append(beego.GlobalControllerRouter["usage-api/controllers:OwnerUsageController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["usage-api/controllers:OwnerUsageController"] = append(beego.GlobalControllerRouter["usage-api/controllers:OwnerUsageController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["usage-api/controllers:QueryUsageController"] = append(beego.GlobalControllerRouter["usage-api/controllers:QueryUsageController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["usage-api/controllers:QueryUsageController"] = append(beego.GlobalControllerRouter["usage-api/controllers:QueryUsageController"],
        beego.ControllerComments{
            Method: "Get",
            Router: `/:level`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["usage-api/controllers:QueryUserController"] = append(beego.GlobalControllerRouter["usage-api/controllers:QueryUserController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["usage-api/controllers:UsageNewController"] = append(beego.GlobalControllerRouter["usage-api/controllers:UsageNewController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"Post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["usage-api/controllers:UsageNewController"] = append(beego.GlobalControllerRouter["usage-api/controllers:UsageNewController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/`,
            AllowHTTPMethods: []string{"Put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["usage-api/controllers:UsageNewController"] = append(beego.GlobalControllerRouter["usage-api/controllers:UsageNewController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["usage-api/controllers:UsageNewController"] = append(beego.GlobalControllerRouter["usage-api/controllers:UsageNewController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["usage-api/controllers:UsageNewController"] = append(beego.GlobalControllerRouter["usage-api/controllers:UsageNewController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["usage-api/controllers:UserController"] = append(beego.GlobalControllerRouter["usage-api/controllers:UserController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["usage-api/controllers:UserController"] = append(beego.GlobalControllerRouter["usage-api/controllers:UserController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
