package models

import (
	"errors"

	"github.com/astaxie/beego/logs"
)

func GetAllUsers() ([]SysOa, error) {
	//var users []SysOa
	users := make([]SysOa, 0)
	err := Engine.Find(&users)
	if err != nil {
		logs.Error(err)
		return nil, errors.New("操作失败")
	}

	return users, nil
}

func GetOneUser(id int) (*SysOa, error) {
	v := &SysOa{}
	has, err := Engine.Id(id).Get(v)
	if err != nil {
		logs.Error(err)
		return nil, err
	}
	if !has {
		return nil, nil
	}
	return v, nil
}
