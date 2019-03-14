package controllers

/*
一二级业务管理权限说明
一级业务只有超级管理员有权增，删，改
二级业务由一级业务管理员和超级管理员进行增，删，改
*/
import (
	"encoding/json"
	"strings"

	"strconv"
	"usage-api/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

// 业务管理接口
type UsageNewController struct {
	BaseController
}
type DeleteErrorInfo struct {
	DeleteSuccess []string            `json:"删除成功"`
	DeleteFailed  []map[string]string `json:"删除失败"`
	NoAuthDelete  []string            `json:"无权限删除"`
}

// URLMapping ...
func (u *UsageNewController) URLMapping() {
	u.Mapping("Post", u.Post)
	u.Mapping("GetOne", u.GetOne)
	u.Mapping("GetAll", u.GetAll)
	u.Mapping("Put", u.Put)
	u.Mapping("Delete", u.Delete)
}

// @Post ...
// @Title Post One
// @Description 创建新的一二级业务
// @Param	body		body 	models.UsageNew	true		"一二级业务信息"
// @Param	cookie		header 	string			true		"sso token"
// @Success 201 {object} RespUsage
// @Failure 403 "插入失败"
// @router / [Post]
func (u *UsageNewController) Post() {
	var v models.UsageNew
	var name string
	var err error
	if beego.AppConfig.String("type") == "dev" {
		name = "it.song"
	} else {
		token := u.Ctx.GetCookie("INNER_AUTH_TOKEN")
		if token == "" {
			u.BaseController.ResponseBody(401, false, "请登录sso认证")
		}

		name, err = models.GetUserByCookie("INNER_AUTH_TOKEN=" + token)
		if err != nil {
			u.BaseController.ResponseBody(500, false, "sso校验出错，请稍后再试", err.Error())
		}
	}
	if err = json.Unmarshal(u.Ctx.Input.RequestBody, &v); err != nil {
		u.BaseController.ResponseBody(400, false, "请求数据格式有误", err.Error())
	}
	logs.Info(v)
	if (v.SecUsage == "" && v.HierarchyLevel != 1) || (v.SecUsage != "" && v.HierarchyLevel != 2) {
		u.BaseController.ResponseBody(500, false, "请求数据不合法")
	}
	if !models.OpUsageNewAuth(v, "Post", name) {
		u.BaseController.ResponseBody(403, false, "你无权限修改对应的人员")
	}

	v.LastUpdatePerson = name
	w, err := models.AddUsage(&v)
	if err != nil {
		u.BaseController.ResponseBody(500, false, "处理出错", err.Error())
	}
	u.BaseController.ResponseBody(201, true, "添加成功", w)

}

// @Put ...
// @Title Put
// @Description 修改一二级业务信息
// @Param	body		body 	models.UsageNew	true		"需要修改的信息"
// @Success 201 {object} RespUsage
// @Failure 403 "更新失败"
// @router / [Put]
func (u *UsageNewController) Put() {
	var v models.UsageNew
	var name string
	var err error
	if beego.AppConfig.String("type") == "dev" {
		name = "it.song"
	} else {
		token := u.Ctx.GetCookie("INNER_AUTH_TOKEN")
		if token == "" {
			u.BaseController.ResponseBody(401, false, "请登录sso认证")
		}

		name, err = models.GetUserByCookie("INNER_AUTH_TOKEN=" + token)
		if err != nil {
			u.BaseController.ResponseBody(500, false, "sso校验出错，请稍后再试", err.Error())
		}
	}
	if err = json.Unmarshal(u.Ctx.Input.RequestBody, &v); err != nil {

		u.BaseController.ResponseBody(400, false, "请求数据格式有误", err.Error())
	}

	if !models.OpUsageNewAuth(v, "Put", name) {
		u.BaseController.ResponseBody(403, false, "你无权限修改对应的人员")
	}

	v.LastUpdatePerson = name
	if err := models.UpdateUsage(&v); err != nil {
		u.BaseController.ResponseBody(500, false, "处理出错", err.Error())
	}

	u.BaseController.ResponseBody(201, true, "更新成功", v)

}

// @GetOne ...
// @Title Get One
// @Description 获取指定的一二级业务信息
// @Param	id		path 	string	true		"一二级业务ID，例如：1"
// @Success 200 {object} RespUsage
// @Failure 403 :id is empty
// @router /:id [get]
func (u *UsageNewController) GetOne() {
	idStr := u.Ctx.Input.Param(":id")
	//id, _ := strconv.Atoi(idStr)
	v, err := models.GetOneUsageById(idStr)
	if err != nil {
		u.BaseController.ResponseBody(500, false, "查询失败", err.Error())

	}
	u.BaseController.ResponseBody(200, true, "查询成功", v)

}

// @GetAll ...
// @Title Get All
// @Description 获取所有的一二级业务信息，或者通过一级业务名称获取二级业务信息
// @Param	usage		query	string	false	"一级业务名称，例如：usage=uhost"
// @Param	level		query	init	false	"获取一级业务列表或者二级业务列表，例如：level=1"
// @Success 200 {object} RespUsage
// @Failure 403 body is empty
// @router / [get]
func (u *UsageNewController) GetAll() {
	usage_name := u.GetString("usage")
	level, _ := u.GetInt("level")
	v, err := models.GetAllUsage(usage_name, level)
	if err != nil {
		u.BaseController.ResponseBody(500, false, "查询失败", err.Error())
	}
	u.BaseController.ResponseBody(200, true, "查询成功", v)
}

// @Delete
// @Title Delete
// @Description 删除指定的一二级业务信息
// @Param	id		path 	[]string	true		"业务ID，例如：1,2,3"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (u *UsageNewController) Delete() {
	var d DeleteErrorInfo
	var name string
	var err error

	ids := strings.Split(u.Ctx.Input.Param(":id"), ",")

	if beego.AppConfig.String("type") == "dev" {
		name = "it.song"
	} else {
		token := u.Ctx.GetCookie("INNER_AUTH_TOKEN")
		if token == "" {
			u.BaseController.ResponseBody(401, false, "请登录sso认证")
		}

		name, err = models.GetUserByCookie("INNER_AUTH_TOKEN=" + token)
		if err != nil {
			u.BaseController.ResponseBody(500, false, "sso校验出错，请稍后再试", err.Error())
		}
	}

	for i := 0; i < len(ids); i++ {
		// /logs.Info(ids[i])
		// if err = models.IsExitUsage(id); err != nil {
		// 	//u.BaseController.ResponseBody(200, true, "要删除的资源已经被删除或不存在!", err.Error())
		// }
		if !models.DeleteUsageNewAuth(ids[i], name) {

			d.NoAuthDelete = append(d.NoAuthDelete, ids[i])
			continue
			//u.BaseController.ResponseBody(403, false, "你无权限修改对应的人员")
		}

		id, _ := strconv.Atoi(ids[i])
		if err := models.DeleteUsage(id); err != nil {
			m := make(map[string]string)
			m[ids[i]] = err.Error()
			d.DeleteFailed = append(d.DeleteFailed, m)
			continue
		}
		d.DeleteSuccess = append(d.DeleteSuccess, ids[i])
	}
	if len(d.DeleteSuccess) != len(ids) {
		u.BaseController.ResponseBody(500, false, "删除失败", d)
	}
	u.BaseController.ResponseBody(200, true, "删除成功!")

}
