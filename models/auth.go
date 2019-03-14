/*
增删改一二级业务权限规则：
一级业务：只有超级用户有权限进行增加，删除，修改
二级业务：超级用户和一级业务的管理员有权限增加删除和修改
*/

/*
增删改一二级业务对应管理员规规则：
一级业务管理员的规则：只有超级用户有权限进行增加，删除，修改
二级业务：超级用户和一级业务的管理员有权限增加删除和修改
*/
package models

import (
	"errors"
	// "database/sql"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

type Info struct {
	key   string
	value string
}

type Resource struct {
	Data Data `json:"data"`
}

type Data struct {
	Category string `json:"category"`
	Info     Body   `json:"info"`
}

type Body struct {
	IP       string      `json:"IP"`
	Usage    string      `json:"Usage"`
	SecUsage interface{} `json:"SecUsage"`
}

type Userinfo struct {
	Account    string `json:"account"`
	FirstLogin string `json:"first_login"`
	Name       string `json:"name"`
	LastLogin  string `json:"last_login"`
}

type LoginInfo struct {
	Action    string   `json:"action"`
	RetCode   int      `json:"ret_code"`
	State     int      `json:"state"`
	LoginInfo Userinfo `json:"login_info"`
}

func GetUserByCookie(cookie string) (string, error) {
	logs.Info(cookie)
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	if cookie == beego.AppConfig.String("tectToken") {
		logs.Info("test dev mode")
		return "yanser.liu", nil
	}
	client := &http.Client{Transport: tr}
	//生成要访问的url
	url := "https://ussov2.ucloudadmin.com/api/sso/check"
	//提交请求
	reqest, err := http.NewRequest("GET", url, nil)

	//增加header选项
	reqest.Header.Add("Cookie", cookie)
	if err != nil {
		panic(err)

		return "", err
	}
	//处理返回结果
	response, err := client.Do(reqest)
	if err != nil {
		return "", err
	}
	var user Userinfo
	defer response.Body.Close()

	logs.Info(response.Status)

	decoder := json.NewDecoder(response.Body)
	err = decoder.Decode(&user)
	logs.Info("请求用户是:", user.Account)
	if user.Account == "" {
		return "", errors.New("sso认证错误")
	}
	return user.Account, nil

}

func AccessCheck(key, value, user string) bool {
	resource, err := GetServerInfo(key, value)
	if err != nil {
		logs.Error(err)
		return false
	}

	if resource.Data.Info.SecUsage == nil {
		//若果资源只隶属于一级业务，则只有一级业务的管理员和用户可以访问，一级业务下的所有二级业务的管理员和用户无法访问
		sql := fmt.Sprintf("select * from (select Id as UId from sys_usage_new where `Usage`='%s' and HierarchyLevel=1 and DeletedAt is NULL) a  JOIN (select UsageId,OAId from sys_owner_usage  where  DeletedAt is NULL) b on a.UId = b.UsageId join (select Id as OId from sys_oa where UserName='%s' and DeletedAt is NULL) c on b.OAId=c.OId", resource.Data.Info.Usage, user)
		logs.Info("查询一级业务", sql)
		c, err := Engine.Sql(sql).Query().Count()
		if err != nil || c == 0 {
			return false
		}
		return true
	}
	//若果资源只隶属于二级业务，则只有该二级业务的管理员和用户，和一级业务的管理员和用户访问。和该二级业务平行的其他二级业务管理员及用户无法访问
	sql := fmt.Sprintf("select * from (select Id as UId from sys_usage_new where `Usage`='%s' and `SecUsage`='%s' and DeletedAt is NULL UNION ALL select Id as UId from sys_usage_new where `Usage`='%s' and HierarchyLevel=1 and DeletedAt is NULL  ) a  JOIN (select UsageId,OAId from sys_owner_usage  where  DeletedAt is NULL) b on a.UId = b.UsageId join (select Id as OId from sys_oa where UserName='%s' and DeletedAt is NULL) c on b.OAId=c.OId", resource.Data.Info.Usage, resource.Data.Info.SecUsage, resource.Data.Info.Usage, user)
	logs.Info("查询二级业务", sql)
	c, err := Engine.Sql(sql).Query().Count()
	if err != nil || c == 0 {
		return false
	}
	return true
}

func IsSuperUser(user string) bool {
	has, _ := Engine.Exist(&SysSuperUser{
		Name: user,
	})
	return has
}

func AddCheck(user, usage, secUsage string) bool {
	if secUsage == "" {
		return true
	}
	return true
}

func DeleteCheck() bool {
	return true
}

func UpdateCheck() bool {
	return true
}

func GetServerInfo(key, value string) (*Resource, error) {

	client := &http.Client{}
	//生成要访问的url
	cmdb_url := fmt.Sprintf("%s::cmdb-api", beego.BConfig.RunMode)

	url := fmt.Sprintf(beego.AppConfig.String(cmdb_url), "?%s=%s", key, value)
	//提交请求
	reqest, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	response, err := client.Do(reqest)
	if err != nil {
		return nil, err
	}
	var resource *Resource
	defer response.Body.Close()

	decoder := json.NewDecoder(response.Body)
	err = decoder.Decode(&resource)
	if err != nil {
		return nil, err
	}
	return resource, nil
}

func GetOAIdByName(name string) (int, error) {
	var oa SysOa
	has, err := Engine.Where("UserName=?", name).Get(&oa)
	if err != nil || !has {
		return 0, errors.New("没有查到该用户")
	}

	return oa.Id, nil
}

func IsAdmin(v SysUsageNew, name string) bool {
	if v.HierarchyLevel == 1 {
		has, _ := Engine.Exist(&SysSuperUser{
			Name: name,
		})
		return has
	}

	return true
}

func OpUsageNewAuth(v UsageNew, method, name string) bool {
	isOk := IsSuperUser(name)
	if isOk {
		return true
	}
	if v.HierarchyLevel == 2 && method != "Get" {
		return CheckOpSecUsageAuth(v, name)
	}

	return true
}

func DeleteUsageNewAuth(id string, name string) bool {
	//超级用户可以删除一级业务和二级业务
	isOk := IsSuperUser(name)
	if isOk {
		return true
	}
	return CheckDeleteSecUsageAuth(id, name)

}

func CheckOpSecUsageAuth(v UsageNew, name string) bool {
	//根只有一级业务的管理员才能删除二级业务
	sql := "select * from (select Id as UId from sys_usage_new where `Usage`='" + v.Usage + "' and HierarchyLevel=1 and DeletedAt is NULL)a join (select UsageId,OAId,IsManager from sys_owner_usage where DeletedAt is NULL and IsManager=1)b on a.UId = b.UsageId join (select Id as OId from sys_oa where UserName='" + name + "' and DeletedAt is NULL) c on b.OAId = c.OId"
	logs.Info(sql)
	has, err := Engine.SQL(sql).Exist()
	if err != nil {
		return false
	}
	logs.Info(has)
	return has
	//return IsUsageAdmin(v, name)

}

func CheckDeleteSecUsageAuth(id string, name string) bool {
	//根据业务名称和用户名，查询用户是否是一级业务的管理员
	//sql := "select * from (select Id as UId,`usage`,SecUsage from sys_usage_new where id=43 and HierarchyLevel=2 and DeletedAt is NULL) a left join (select OAId,UsageId,IsManager from sys_owner_usage where DeletedAt is NULL and IsManager=1) b on a.UId = b.UsageId join (select Id as NameId,UserName from sys_oa where UserName='yanser.liu' and DeletedAt is NULL) c on b.OAId = c.NameId "
	a := "(select `usage`,SecUsage from sys_usage_new where id=" + id + " and HierarchyLevel=2 and DeletedAt is NULL) a"
	b := "(select Id as UsageId,`usage` as UsageName from  sys_usage_new where DeletedAt is NULL and HierarchyLevel=1) b"
	c := "(select OAId,UsageId,IsManager from sys_owner_usage where DeletedAt is NULL and IsManager=1) c"
	d := "(select Id as NameId,UserName from sys_oa where UserName='" + name + "' and DeletedAt is NULL) d"
	sql := "select * from " + a + " join " + b + " on a.`usage` = b.UsageName join " + c + " on b.UsageId = c.UsageId join " + d + " on c.OAId = d.NameId "
	logs.Info(sql)
	has, err := Engine.SQL(sql).Exist()
	if err != nil {
		return false
	}
	logs.Info(has)
	return has
	//return IsUsageAdmin(v, name)

}

func CheckOwnerUsageAuth(name string, usageId int) bool {
	if IsSuperUser(name) {
		return true
	}

	oaId, err := GetOAIdByName(name)
	if err != nil {
		return false
	}
	return checkOwnerAuth(oaId, usageId)
}

func checkOwnerAuth(oaid int, usageId int) bool {
	var usage SysUsageNew
	if has, err := Engine.Where("Id=?", usageId).Get(&usage); err != nil || !has {
		logs.Info(err)
		return false
	}
	//如果是一级业务，则只能由一级业务管理员操作
	if usage.HierarchyLevel == 1 {
		has, err := Engine.Exist(&SysOwnerUsage{
			OAId:      oaid,
			UsageId:   usageId,
			IsManager: 1,
		})
		logs.Info("是不是一级业务管理员：", has)
		if err != nil {
			return false
		}

		return has
	}
	//如果不是一级业务，则只能由二级业务管理员 或者对应的一级管理员
	has, _ := Engine.Exist(&SysOwnerUsage{ //是不是二级管理员
		OAId:      oaid,
		UsageId:   usageId,
		IsManager: 1,
	})
	logs.Info("二级管理员：", has)
	if has {
		return true
	}
	//是不是对应的一级管理员
	sql := "select * from (select Id as UsId from sys_usage_new where `Usage`='" + usage.Usage + "' and HierarchyLevel=1 and DeletedAt is NULL)a join (select * from sys_owner_usage where DeletedAt is NULL  and IsManager=1 and OAId='" + strconv.Itoa(oaid) + "' )b on a.UsId=b.UsageId "
	logs.Info(sql)
	has, _ = Engine.Sql(sql).Exist()
	logs.Info("一级管理员：", has)
	if !has {
		return false
	}

	return true
}
