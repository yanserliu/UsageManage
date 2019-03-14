// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"usage-api/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/usage",
			beego.NSInclude(
				&controllers.UsageNewController{},
			),
		),
		beego.NSNamespace("/user",
			beego.NSInclude(
				&controllers.UserController{},
			),
		),
		beego.NSNamespace("/owner",
			beego.NSInclude(
				&controllers.OwnerUsageController{},
			),
		),

		beego.NSNamespace("/check",
			beego.NSInclude(
				&controllers.CheckController{},
			),
		),
		beego.NSNamespace("/query/usage/user/:name",
			beego.NSInclude(
				&controllers.QueryUsageController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
