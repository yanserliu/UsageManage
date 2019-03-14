package models

import (
	"fmt"
	"os"
	"path"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	_ "github.com/go-sql-driver/mysql"

	"github.com/xormplus/core"
	"github.com/xormplus/xorm"
)

var Engine *xorm.Engine

func init() {
	var err error
	dbconn := fmt.Sprintf("%s::DBConnect", beego.BConfig.RunMode)
	Engine, err = xorm.NewEngine("mysql", beego.AppConfig.String(dbconn))
	if err != nil {
		logs.Error(err.Error())
		os.Exit(-1)
	}

	if err := Engine.Ping(); err != nil {
		logs.Error(err.Error())
		os.Exit(-1)
	}
	if err := SetEngine(); err != nil {
		logs.Error(err.Error())
		os.Exit(-1)
	}

	//SameMapper 支持结构体名称和对应的表名称以及结构体field名称与对应的表字段名称相同的命名； *
	Engine.SetColumnMapper(core.SameMapper{})
	Engine.SetTableMapper(core.GonicMapper{})
	//Engine.Table("ServerOwner").CreateTable(&ServerOwner{})                            //重命名表名称
	//Engine.Table("VMServerOwner").CreateTable(&VMServerOwner{})                        //重命名表名称
	err = Engine.CreateTables(&SysOwnerUsage{}, &SysOa{}, &SysUsageNew{}, &SysSuperUser{}) //SameMapper规则创建表
	if err != nil {
		logs.Error("manage.infrastructure.db.initdb.init() CreateTables errors:", err)
	}

}

func SetEngine() error {
	logDir := "logs"
	if _, e := os.Open(logDir); e != nil {
		os.MkdirAll(logDir, os.ModePerm)
	}

	logPath := path.Join(logDir, "sql.log")
	f, err := os.Create(logPath)
	if err != nil {
		return err
	}
	//记录sql执行命令
	Engine.SetLogger(xorm.NewSimpleLogger(f))
	Engine.ShowSQL(true)

	if location, err := time.LoadLocation("Asia/Shanghai"); err == nil {
		Engine.TZLocation = location
	}

	return err
}
