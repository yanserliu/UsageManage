package controllers

import (
	//	"strings"

	//"usage-api/models"

	// "time"

	// "fmt"

	// "usage-api/helper"

	"github.com/astaxie/beego"
)

type BaseController struct {
	beego.Controller
	//TplTheme  string //模板主题
	//TplStatic string //模板静态文件
	//AdminId   int    //管理员是否已登录，如果已登录，则管理员ID大于0
	//	Sys       models.Sys
}

type RespUsage struct {
	Status  bool                `json:"status"`
	Message string              `json:"message"`
	Data    []map[string]string `json:"data"`
}

//404
func (this *BaseController) Error404() {
	this.Layout = ""
	this.Data["content"] = "Page Not Foud"
	this.Data["code"] = "404"
	this.Data["content_zh"] = "页面被外星人带走了"
	this.TplName = "error.html"
}

//501
func (this *BaseController) Error501() {
	this.Layout = ""
	this.Data["code"] = "501"
	this.Data["content"] = "Server Error"
	this.Data["content_zh"] = "服务器被外星人戳炸了"
	this.TplName = "error.html"
}

//数据库错误
func (this *BaseController) ErrorDb() {
	this.Layout = ""
	this.Data["content"] = "Database is now down"
	this.Data["content_zh"] = "数据库被外星人抢走了"
	this.TplName = "error.html"
}

//更新内容
// func (this *BaseController) Update() {
// 	id := strings.Split(this.GetString("id"), ",")
// 	i, err := models.UpdateByIds(this.GetString("table"), this.GetString("field"), this.GetString("value"), id)
// 	ret := map[string]interface{}{"status": 0, "msg": "更新失败，可能您未对内容作更改"}
// 	if i > 0 && err == nil {
// 		ret["status"] = 1
// 		ret["msg"] = "更新成功"
// 	}
// 	if err != nil {
// 		ret["msg"] = err.Error()
// 	}
// 	this.Data["json"] = ret
// 	this.ServeJSON()
// }

//删除内容
// func (this *BaseController) Del() {
// 	id := strings.Split(this.GetString("id"), ",")
// 	i, err := models.DelByIds(this.GetString("table"), id)
// 	ret := map[string]interface{}{"status": 0, "msg": "删除失败，可能您要删除的内容已经不存在"}
// 	if i > 0 && err == nil {
// 		ret["status"] = 1
// 		ret["msg"] = "删除成功"
// 	}
// 	if err != nil {
// 		ret["msg"] = err.Error()
// 	}
// 	this.Data["json"] = ret
// 	this.ServeJSON()
// }

//响应json
func (this *BaseController) ResponseJson(isSuccess bool, msg string, data ...interface{}) {
	// status := 0
	// if isSuccess {
	// 	status = 1
	// }
	ret := map[string]interface{}{"status": isSuccess, "msg": msg}
	if len(data) > 0 {
		ret["data"] = data[0]
	}
	this.Data["json"] = ret
	this.ServeJSON()
	this.StopRun()
}

func (this *BaseController) ResponseBody(code int, isSuccess bool, msg string, data ...interface{}) {
	this.Ctx.Output.SetStatus(code)
	ret := map[string]interface{}{"status": isSuccess, "msg": msg}
	if len(data) > 0 {
		ret["data"] = data[0]
	}
	this.Data["json"] = ret
	this.ServeJSON()
	this.StopRun()
}

//响应json
func (this *BaseController) Response(data map[string]interface{}) {
	this.Data["json"] = data
	this.ServeJSON()
	this.StopRun()
}
