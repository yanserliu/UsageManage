package models

import (
	"errors"
	//	"strconv"
	//"time"
	//"fmt"
	"strconv"

	"github.com/astaxie/beego/logs"
)

type OwnerUsageInfo struct {
	DepartmentCN   string
	HierarchyLevel int
	Id             int
	IsManager      int
	NameCn         string
	OAId           int
	Usage          string
	SecUsage       string
	UsageId        int
	UserName       string
}

type UsageInfo struct {

	//IsManager      int

	Id             int
	Usage          string
	SecUsage       string
	HierarchyLevel int
}

func IsExitOwnerUsage(id int) error {

	has, err := Engine.Exist(&SysOwnerUsage{
		Id: id,
	})
	if err != nil {
		logs.Error(err)
		return err
	}
	if !has {
		logs.Warn("查找的人员-业务关系不存在")
		return errors.New("对应的资源不存在")
	}
	return nil
}

func DeleteOwnerUsage(id int) error {
	v := &SysOwnerUsage{Id: id}
	_, err := Engine.Delete(v)
	if err != nil {
		logs.Error(err)
		return err
	}
	return nil
}

func GetOwnerUsageById(id string) (*OwnerUsageInfo, error) {
	// sql := "select Id,UserName,NameCn,OAId,DepartmentCN,UsageId,`Usage`,SecUsage,Hierarchylevel,IsManager from (select Id,OAId,IsManager,UsageId from sys_owner_usage where DeletedAt is NULL and Id =" + id +
	// 	") a left join (select UserName,NameCn,Id as Oid,DepartmentCN from sys_oa where DeletedAt is NULL) b  on a.OAId = b.Oid left join (select Id as uId, `Usage`,SecUsage,Hierarchylevel from sys_usage_new where Deletedat is NULL) c on a.UsageId =c.uId"
	// logs.Info(sql)

	//has, err := Engine.Sql(sql).Get(&ownerUsageInfo)

	var ownerUsageInfo OwnerUsageInfo
	has, err := Engine.Table("sys_owner_usage").
		Join("LEFT", "sys_oa", "sys_owner_usage.OAId=sys_oa.Id").
		Join("left", "sys_usage_new", "sys_owner_usage.UsageId =sys_usage_new.Id").
		Where("sys_owner_usage.Id=?", id).
		Get(&ownerUsageInfo)

	if err != nil {
		logs.Info(err)
		return nil, err
	}
	if !has {
		logs.Warn("查找的人员-业务关系不存在")
		return nil, errors.New("查找的资源不存在")
	}
	return &ownerUsageInfo, nil
}

func GetAllOwnerUsage(user string, usage_id int) ([]OwnerUsageInfo, error) {
	var err error
	//var ownerUsageInfo []OwnerUsageInfo
	ownerUsageInfo := make([]OwnerUsageInfo, 0)
	if user != "" {
		err = Engine.Table("sys_owner_usage").
			Join("LEFT", "sys_oa", "sys_owner_usage.OAId=sys_oa.Id").
			Join("left", "sys_usage_new", "sys_owner_usage.UsageId =sys_usage_new.Id").
			Where("sys_oa.UserName=?", user).
			Find(&ownerUsageInfo)
	} else if usage_id > 0 {
		err = Engine.Table("sys_owner_usage").
			Join("LEFT", "sys_oa", "sys_owner_usage.OAId=sys_oa.Id").
			Join("left", "sys_usage_new", "sys_owner_usage.UsageId =sys_usage_new.Id").
			Where("sys_owner_usage.UsageId=?", strconv.Itoa(usage_id)).
			Find(&ownerUsageInfo)
	} else {
		err = Engine.Table("sys_owner_usage").
			Join("LEFT", "sys_oa", "sys_owner_usage.OAId=sys_oa.Id").
			Join("left", "sys_usage_new", "sys_owner_usage.UsageId =sys_usage_new.Id").
			Find(&ownerUsageInfo)
	}

	if err != nil {
		logs.Error(err)
		return nil, err
	}
	return ownerUsageInfo, nil
}

func GetAllUsageByUserName(name string) ([]UsageInfo, error) {

	// sql := "select UsageId,`Usage`,SecUsage,HierarchyLevel,IsManager from (select Id as UserId,UserName,NameCn from sys_oa where UserName='" + name +
	// 	"' and DeletedAt is NULL) a left join (select OAId,UsageId,IsManager from sys_owner_usage where DeletedAt is NULL) b on a.UserId=b.OAId left join (select Id,`Usage`,SecUsage,HierarchyLevel from sys_usage_new where DeletedAt is NULL) c on b.UsageId=c.Id"

	sql := "select distinct Id,e.* from  (select d.* from (select * from sys_owner_usage where DeletedAt is NULL)a join (select Id as UId from sys_oa where UserName='" + name + "' and DeletedAt is NULL)b on a.OAId=b.UId" +
		" join (select Id as UsId,`Usage` as `Usage2`,SecUsage,HierarchyLevel from sys_usage_new where DeletedAt is NULL and HierarchyLevel=1)c on c.UsId=a.UsageId" +
		" left join (select Id,`Usage`,SecUsage,HierarchyLevel from sys_usage_new where DeletedAt is NULL)d on c.`Usage2`=d.`Usage` union all" +
		" select c.* from (select UsageId as UsId,OAId from sys_owner_usage where DeletedAt is NULL)a join (select Id as UId from sys_oa where UserName='yanser.liu' and DeletedAt is NULL)b on a.OAId=b.UId" +
		" join (select Id ,`Usage`,SecUsage,HierarchyLevel from sys_usage_new where DeletedAt is NULL and HierarchyLevel=2)c on c.Id=a.UsId) e"

	logs.Info(sql)
	//var usageInfo []UsageInfo
	usageInfo := make([]UsageInfo, 0)
	err := Engine.Sql(sql).Find(&usageInfo)
	if err != nil {
		logs.Error(err)
		return nil, err
	}

	return usageInfo, nil
}

func GetLevelUsageByUserName(name, level string) ([]UsageInfo, error) {
	var sql string
	if level == "1" {
		sql = "select UsageId,`Usage`,SecUsage,HierarchyLevel from (select Id as UserId,UserName,NameCn from sys_oa where UserName='" + name +
			"' and DeletedAt is NULL) a join (select OAId,UsageId from sys_owner_usage where DeletedAt is NULL) b on a.UserId=b.OAId join (select Id,`Usage`,SecUsage,HierarchyLevel from sys_usage_new " +
			"where DeletedAt is NULL and HierarchyLevel =1) c on b.UsageId=c.Id"
	}
	if level == "2" {
		sql = "select distinct Id,`Usage`,SecUsage,HierarchyLevel from  (select d.* from (select * from sys_owner_usage where DeletedAt is NULL)a join (select Id as UId from sys_oa where UserName='" + name +
			"' and DeletedAt is NULL)b on a.OAId=b.UId join (select * from sys_usage_new where DeletedAt is NULL and HierarchyLevel=1)c on c.Id=a.UsageId left join (select Id ,`Usage`,SecUsage,HierarchyLevel" +
			" from sys_usage_new where DeletedAt is NULL and HierarchyLevel=2) d on d.`Usage`=c.`Usage` union all select c.* from (select UsageId as UsId,OAId from sys_owner_usage where DeletedAt is NULL)" +
			"a join (select Id as UId from sys_oa where UserName='" + name + "' and DeletedAt is NULL)b on a.OAId=b.UId join (select Id ,`Usage`,SecUsage,HierarchyLevel from sys_usage_new where DeletedAt is NULL and HierarchyLevel=2)c on c.Id=a.UsId) e"
	}
	logs.Info(sql)
	//var usageInfo []UsageInfo
	usageInfo := make([]UsageInfo, 0)
	err := Engine.Sql(sql).Find(&usageInfo)
	if err != nil {
		logs.Error(err)
		return nil, err
	}

	return usageInfo, nil
}

func AddOwnerUsage(ownerUsage *SysOwnerUsage) error {
	if ownerUsage.IsManager > 1 {
		return errors.New("添加的角色不存在吗，只能添加1：管理员，0：访客")
	}

	has, err := Engine.Exist(&SysOwnerUsage{
		UsageId:   ownerUsage.UsageId,
		OAId:      ownerUsage.OAId,
		IsManager: ownerUsage.IsManager,
	})

	if err != nil {
		logs.Error(err)
		return err
	}
	if has {
		return errors.New("重复插入，记录已存在")
	}
	if CheckIsOneUsageAuth(ownerUsage) {
		return errors.New("该用户已经是二级业务的一级业务的管理员，具有该二级业务的最高权限，无需再添加到该二级业务管理员或访客角色")
	}
	row, err := Engine.Insert(ownerUsage)
	if err != nil {
		logs.Error(err)
		return err
	}
	if row == 0 {
		return errors.New("插入失败")
	}
	return nil
}

func CheckIsOneUsageAuth(ownerUsage *SysOwnerUsage) bool {
	var usage string
	var id int
	has, err := Engine.Table("sys_usage_new").Where("Id=? and HierarchyLevel=2", ownerUsage.UsageId).Cols("Usage").Get(&usage)
	if err != nil {
		logs.Error(err)
		return false
	}
	if !has {
		logs.Error("没有查到二级业务信息，该业务是一级业务")
		return false
	}
	logs.Error(usage)
	has, err = Engine.Table("sys_usage_new").Where("`Usage`=? and HierarchyLevel=1", usage).Cols("Id").Get(&id)
	if err != nil {
		logs.Error(err)
		return false
	}
	if !has {
		logs.Error("没有查到对应的一级业务信息")
		return false
	}
	logs.Error(id)
	has, err = Engine.Where("UsageId=? and IsManager=1", id).Exist(&SysOwnerUsage{})
	if err != nil {
		logs.Error(err)
		return false
	}
	if !has {
		logs.Error("不是一级业务管理员")
		return false
	}
	return true
}

func UpdateOwnerUsage(ownerUsage *SysOwnerUsage) error {
	if ownerUsage.IsManager > 1 {
		return errors.New("更新的角色不存在吗，只能更新成 1：管理员，0：访客")
	}

	//ownerUsage只更新"IsManager", "LastUpdatePerson"，LastUpdateTime 三个字段
	row, err := Engine.ID(ownerUsage.Id).Cols("IsManager", "LastUpdatePerson").Update(ownerUsage)
	if err != nil {
		logs.Error(err)
		return err
	}
	if row == 0 {
		return errors.New("要更新的记录不存在")
	}
	return nil
}

func CheckOwnerUsage(name, ip, sn string) ([]map[string]interface{}, error) {
	myArray := make([]map[string]interface{}, 0)
	sql := "select * from sys_owner_usage"
	logs.Info(sql)
	res, err := Engine.QueryInterface(sql)
	if err != nil {
		logs.Error(err)
		return nil, errors.New("操作失败")
	}
	if len(res) == 0 {
		return myArray, nil
	}
	logs.Info(res)
	return res, nil
}
