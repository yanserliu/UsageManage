package controllers

/*
业务管理员权限说明
超级管理员可以管理多有（多有一级业务）
一级业务的管理员可以管理对应的一级业务及其下属的所有人员与业务的关系（对应一级管理员，一级非管理员，二级管理员，二级非管理员）
二级业务的管理员只能管理该二级业务的管理员和非管理员
一二级业务的非管理员均没有业务和人员的管理权限。
*/

import (
	"encoding/json"
	"strconv"
	"strings"
	"usage-api/models"

	"github.com/astaxie/beego"
)

// 管理“人员-业务”之间的关系
type OwnerUsageController struct {
	BaseController
}

// URLMapping ...
func (o *OwnerUsageController) URLMapping() {
	o.Mapping("Post", o.Post)
	o.Mapping("GetOne", o.GetOne)
	o.Mapping("GetAll", o.GetAll)
	o.Mapping("Put", o.Put)
	o.Mapping("Delete", o.Delete)
}

// Post ...
// @Title Post One
// @Description 创建“人员-业务”关联关系
// @Param	body		body 	models.SysOwnerUsage	true		"业务ID，人员ID，是否是管理员为必填字段，其他字段均不需要填写"
// @Success 201 {object} RespUsage
// @Failure 403 "插入失败"
// @router / [Post]
func (o *OwnerUsageController) Post() {
	var v models.SysOwnerUsage
	var name string
	var err error
	if beego.AppConfig.String("type") == "dev" {
		name = "it.song"
	} else {
		token := o.Ctx.GetCookie("INNER_AUTH_TOKEN")

		if token == "" {
			o.BaseController.ResponseBody(401, false, "请登录sso认证")
		}
		name, err = models.GetUserByCookie("INNER_AUTH_TOKEN=" + token)
		if err != nil {
			o.BaseController.ResponseBody(500, false, "sso校验出错，请稍后再试", err.Error())
		}
	}
	if err := json.Unmarshal(o.Ctx.Input.RequestBody, &v); err != nil {
		o.BaseController.ResponseBody(400, false, "请求数据格式有误", err.Error())
	}

	if !models.CheckOwnerUsageAuth(name, v.UsageId) {
		o.BaseController.ResponseBody(403, false, "你无权限修改对应的人员")

	}
	if err := models.AddOwnerUsage(&v); err != nil {
		o.BaseController.ResponseBody(500, false, "处理出错", err.Error())
	}

	o.BaseController.ResponseBody(201, true, "新增成功", v)
}

// Put ...
// @Title Put
// @Description 修改业务-人员关系
// @Param	body		body 	models.SysOwnerUsage	true		"Id，IsManager 是必填字段，其他字段均不需要填写"
// @Success 201 {object} RespUsage
// @Failure 403 "更新失败"
// @router / [Put]
func (o *OwnerUsageController) Put() {
	var v models.SysOwnerUsage
	var name string
	var err error
	if beego.AppConfig.String("type") == "dev" {
		name = "it.song"
	} else {
		token := o.Ctx.GetCookie("INNER_AUTH_TOKEN")
		if token == "" {
			o.BaseController.ResponseBody(401, false, "请登录sso认证")
		}
		name, err = models.GetUserByCookie("INNER_AUTH_TOKEN=" + token)
		if err != nil {
			o.BaseController.ResponseBody(500, false, "sso校验失败", err.Error())
		}
	}
	if err := json.Unmarshal(o.Ctx.Input.RequestBody, &v); err != nil {
		o.BaseController.ResponseBody(400, false, "请求数据格式有误", err.Error())
	}
	//查找要更新的记录原信息
	w, err := models.GetOwnerUsageById(strconv.Itoa(v.Id))
	if err != nil {
		o.BaseController.ResponseBody(400, false, "所要更新的记录不存在", err.Error())
	}

	if !models.CheckOwnerUsageAuth(name, w.UsageId) {
		o.BaseController.ResponseBody(403, false, "该条记录你无权限修改")
	}
	v.LastUpdatePerson = name
	if err := models.UpdateOwnerUsage(&v); err != nil {
		o.BaseController.ResponseBody(500, false, "内部处理错误", err.Error())

	}

	o.BaseController.ResponseBody(201, true, "更新成功", v)
}

// GetOne ...
// @Title Get One
// @Description 通过人员业务关联ID查询该条数据下的信息
// @Param	id		path 	string	true		"人员业务关联ID，例如：1"
// @Success 200 {object} RespUsage
// @Failure 403 :id is empty
// @router /:id [get]
func (o *OwnerUsageController) GetOne() {
	idStr := o.Ctx.Input.Param(":id")
	v, err := models.GetOwnerUsageById(idStr)
	if err != nil {
		o.BaseController.ResponseBody(500, false, "查询失败", err.Error())

	}
	o.BaseController.ResponseBody(200, true, "查询成功", v)

}

// GetAll ...
// @Title Get All
// @Description 获取所有人员业务关系信息，或者某人名下的业务信息，某业务下的所属人信息
// @Param	user		query  	string	false	"人员名称，例如：yanser.liu"
// @Param	usage_id	query  	int	    false	"业务ID,例如：1"
// @Success 200 {object} RespUsage
// @Failure 403 body is empty
// @router / [get]
func (o *OwnerUsageController) GetAll() {
	user := o.GetString("user")
	usageId, err := o.GetInt("usage_id")
	if err != nil {
		usageId = 0
	}
	v, err := models.GetAllOwnerUsage(user, usageId)
	if err != nil {
		o.BaseController.ResponseBody(500, false, "查询失败", err.Error())
	}
	o.BaseController.ResponseBody(200, true, "查询成功", v)
}

// @Title Delete
// @Description 删除人员-业务关联关系
// @Param	id		path 	[]string	true		"OwnerUsageId,例如：1,2,3,4"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (o *OwnerUsageController) Delete() {
	var d DeleteErrorInfo
	var name string
	var err error
	ids := strings.Split(o.Ctx.Input.Param(":id"), ",")

	if beego.AppConfig.String("type") == "dev" {
		name = "it.song"
	} else {
		token := o.Ctx.GetCookie("INNER_AUTH_TOKEN")
		if token == "" {
			o.BaseController.ResponseBody(401, false, "请登录sso认证")
		}

		name, err = models.GetUserByCookie("INNER_AUTH_TOKEN=" + token)
		if err != nil {
			o.BaseController.ResponseBody(500, false, "sso校验出错，请稍后再试", err.Error())
		}
	}
	for i := 0; i < len(ids); i++ {
		//查找要删除的记录原信息
		w, err := models.GetOwnerUsageById(ids[i])
		if err != nil {
			d.DeleteSuccess = append(d.DeleteSuccess, ids[i])
			continue
			//o.BaseController.ResponseBody(200, true, "要删除的资源已经被删除或不存在!", err.Error())
		}

		if !models.CheckOwnerUsageAuth(name, w.UsageId) {
			d.NoAuthDelete = append(d.NoAuthDelete, ids[i])
			continue
			//o.BaseController.ResponseBody(403, false, "你无权限删除对应的人员")
		}
		id, _ := strconv.Atoi(ids[i])
		if err := models.DeleteOwnerUsage(id); err != nil {
			m := make(map[string]string)
			m[ids[i]] = err.Error()
			d.DeleteFailed = append(d.DeleteFailed, m)
			continue
		}
		d.DeleteSuccess = append(d.DeleteSuccess, ids[i])
	}

	if len(d.DeleteSuccess) != len(ids) {
		o.BaseController.ResponseBody(500, false, "删除失败", d)
	}
	o.BaseController.ResponseBody(200, true, "删除成功!")
	//o.BaseController.ResponseBody(200, true, "删除成功!")

}
