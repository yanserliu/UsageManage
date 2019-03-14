package controllers

/*
用户接入权限说明
一级业务的管理员和非管理员可以接入对应一级业务和所有下属的二级业务
二级业务的管理员和非管理员可以接入对应二级业务
*/

import (
	//"encoding/json"
	"usage-api/models"
	//"github.com/astaxie/beego/logs"
)

// 通过IP/SN 和用户，查询匹配规则
type CheckController struct {
	BaseController
}

// URLMapping ...
func (u *CheckController) URLMapping() {
	u.Mapping("Get", u.Get)

}

// Get ...
// @Title Get
// @Description 校验用户是否有访问Server或者VMServer权限
// @Param	key		query	string	true	"参数key，例如：key=ip或key=sn"
// @Param	value   query	string	true	"参数value,例如:value=192.168.150.29或value=21345"
// @Param	user	query	string	true	"参数用户 例如:user=yanser.liu"
// @Success 200 校验通过
// @Failure 403 禁止接入
// @router / [get]
func (u *CheckController) Get() {
	user_name := u.GetString("user")
	key := u.GetString("key")
	value := u.GetString("value")
	is := models.IsSuperUser(user_name)
	if is {
		u.Ctx.Output.SetStatus(200)
		u.ServeJSON()
		return
	}
	t := models.AccessCheck(key, value, user_name)
	if !t {
		u.Ctx.Output.SetStatus(403)
		u.ServeJSON()
		return

	}
	u.Ctx.Output.SetStatus(200)

	u.ServeJSON()

}
