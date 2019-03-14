package controllers

import (
	// "encoding/json"
	// "strconv"
	"usage-api/models"
	//"github.com/astaxie/beego/logs"
)

// Query Usage By UserName Controller
type QueryUsageController struct {
	BaseController
}

// URLMapping ...
func (q *QueryUsageController) URLMapping() {
	q.Mapping("GetOne", q.Get)
	q.Mapping("GetAll", q.GetAll)
}

// GetAll ...
// @Title GetAll
// @Description get All Usage By UserName
// @Param	name		path 	string	true		"用户名, 例如:yanser.liu"
// @Success 200 {object} RespUsage
// @Failure 403 body is empty
// @router / [get]
func (q *QueryUsageController) GetAll() {
	name := q.Ctx.Input.Param(":name")
	v, err := models.GetAllUsageByUserName(name)
	if err != nil {
		q.BaseController.ResponseBody(500, false, "查询失败", err)
	}
	q.BaseController.ResponseBody(200, true, "查询成功", v)
}

// Get ...
// @Title Get
// @Description get All 1 Level or 2 Level Usage By UserName
// @Param	name		path 	string	true		"用户名, 例如:yanser.liu"
// @Param	level		path 	string	true		"一二级业务等级, 例如：1或2"
// @Success 200 {object} RespUsage
// @Failure 403 body is empty
// @router /:level [get]
func (q *QueryUsageController) Get() {
	name := q.Ctx.Input.Param(":name")
	level := q.Ctx.Input.Param(":level")
	v, err := models.GetLevelUsageByUserName(name, level)
	if err != nil {
		q.BaseController.ResponseBody(500, false, "查询失败", err)
	}
	q.BaseController.ResponseBody(200, true, "查询成功", v)
}
