package controllers

import (
	"strconv"
	"usage-api/models"
	//"github.com/astaxie/beego"
)

// 查询OA人员信息接口
type UserController struct {
	BaseController
}

// URLMapping ...
func (u *UserController) URLMapping() {
	u.Mapping("GetOne", u.GetOne)
	u.Mapping("GetAll", u.GetAll)
}

// GetAll ...
// @Title GetAll
// @Description 获取所有OA人员信息
// @Success 200 {object} []models.SysOa
// @Failure 403 body is empty
// @router / [get]
func (u *UserController) GetAll() {
	v, err := models.GetAllUsers()
	if err != nil {
		u.BaseController.ResponseJson(false, "查询失败", err.Error())
		u.Ctx.Output.SetStatus(403)
	}
	u.BaseController.ResponseJson(true, "查询成功", v)
}

// GetOne ...
// @Title GetOne
// @Description 获取指定OA人员信息
// @Param	id		path 	string	true		"OA人员ID，例如：1"
// @Success 200 {object} models.User
// @Failure 403 :id is empty
// @router /:id [get]
func (u *UserController) GetOne() {
	idStr := u.Ctx.Input.Param(":id")

	if idStr != "" {
		id, _ := strconv.Atoi(idStr)
		v, err := models.GetOneUser(id)
		if err != nil {
			u.BaseController.ResponseJson(false, "查询失败", err.Error())
			u.Ctx.Output.SetStatus(403)

		}
		u.BaseController.ResponseJson(true, "查询成功", v)
	}
	//u.ServeJSON()
}
