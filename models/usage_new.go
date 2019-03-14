package models

import (
	"errors"
	//"strings"
	"strconv"
	//"time"

	"github.com/astaxie/beego/logs"
)

type UsageNewInfo struct {
	Id               int    `json:"Id" xorm:"not null pk autoincr INT(11)"`
	Usage            string `json:"Usage" xorm:"not null VARCHAR(255)"`
	SecUsage         string `json:"SecUsage" xorm:"VARCHAR(255)"`
	HierarchyLevel   int    `json:"HierarchyLevel" xorm:"INT(11)"`
	ResourceNum      int    `json:"ResourceNum" xorm:"default 0 INT(11)"`
	Admin            string `json:"Admin" xorm:"VARCHAR(255)"`
	CreatedAt        string `json:"CreatedAt" xorm:"created DATETIME"`
	DeletedAt        string `json:"DeletedAt" xorm:"deleted DATETIME"`
	LastUpdatePerson string `json:"LastUpdatePerson" xorm:"VARCHAR(255)"`
	LastUpdateTime   string `json:"LastUpdateTime" xorm:"updated DATETIME"`
	Remark           string `json:"Remark" xorm:"VARCHAR(255)"`
}

type UsageNew struct {
	Id               int    `json:"Id"`
	Usage            string `json:"Usage"`
	SecUsage         string `json:"SecUsage"`
	HierarchyLevel   int    `json:"HierarchyLevel"`
	ResourceNum      int    `json:"ResourceNum"`
	Admin            []int  `json:"Admin"`
	LastUpdatePerson string `json:"LastUpdatePerson"`
	Remark           string `json:"Remark"`
}

func IsExitUsage(id int) error {

	has, err := Engine.Exist(&SysUsageNew{
		Id: id,
	})
	if err != nil {
		logs.Error(err)
		return err
	}
	if !has {
		return errors.New("对应的资源不存在")
	}
	return nil
}

func DeleteUsage(id int) error {
	v := &SysUsageNew{}

	has, err := Engine.ID(id).Get(v)
	if err != nil {
		logs.Error(err)
		//return errors.New("需要的删除记录不存在")
		return err
	}
	if !has {
		return nil
	}
	if v.HierarchyLevel == 1 {
		has, err := Engine.Exist(&SysUsageNew{
			Usage:          v.Usage,
			HierarchyLevel: 2,
		})
		if err != nil {
			logs.Error(err)
			return err
		}
		if has {
			return errors.New("该一级业务下存在二级业务，禁止删除")
		}
	}
	//若是二级业务，判定是否绑定有资源ps:20190314不判断是否绑定资源
	// has, err = Engine.Where("Id=? and ResourceNum > ?", id, 0).Exist(&SysUsageNew{})
	// if err != nil {
	// 	return err
	// }
	// if has {
	// 	return errors.New("该业务下绑定有资源，禁止删除")
	// }
	//事物处理，删除sys_owner_usage和sys_usage_new
	session := Engine.NewSession()
	defer session.Close()
	err = session.Begin()
	if err != nil {
		return err
	}

	_, err = session.Where("UsageId = ?", id).Delete(&SysOwnerUsage{})
	if err != nil {
		session.Rollback()
		return err
	}

	_, err = session.ID(id).Delete(&SysUsageNew{})
	if err != nil {
		session.Rollback()
		return err
	}

	err = session.Commit()
	if err != nil {
		return err
	}
	return nil
}

func GetAllUsage(usage string, level int) ([]UsageNewInfo, error) {
	//var usageInfo []UsageNewInfo
	usageInfo := make([]UsageNewInfo, 0)
	var sql string
	if usage != "" {
		sql = "select a.*,group_concat(c.UserName) as Admin from (SELECT * from sys_usage_new where DeletedAt is NULL and `Usage`='" + usage + "' and HierarchyLevel=2)a left join  (select UsageId,OAId from sys_owner_usage  where DeletedAt is NULL and IsManager=1)b on a.Id =b.UsageId left join  (select Id,UserName,NameCn from sys_oa  where deletedAt is NULL) c on b.OAId=c.Id group by Id"
	} else if level != 0 {
		sql = "select a.*,group_concat(c.UserName) as Admin from (SELECT * from sys_usage_new where DeletedAt is NULL and HierarchyLevel='" + strconv.Itoa(level) + "')a left join  (select UsageId,OAId from sys_owner_usage  where DeletedAt is NULL and IsManager=1)b on a.Id =b.UsageId left join  (select Id,UserName,NameCn from sys_oa  where deletedAt is NULL) c on b.OAId=c.Id group by Id"
	} else {
		sql = "select a.*,group_concat(c.UserName) as Admin from (SELECT * from sys_usage_new where DeletedAt is NULL)a left join  (select UsageId,OAId from sys_owner_usage  where DeletedAt is NULL and IsManager=1)b on a.Id =b.UsageId left join  (select Id,UserName,NameCn from sys_oa  where deletedAt is NULL) c on b.OAId=c.Id group by Id"
	}
	// err := Engine
	// .Find(&usageInfo)
	err := Engine.Sql(sql).Find(&usageInfo)
	logs.Info(sql)
	if err != nil {
		logs.Error(err)
		return nil, err
	}
	return usageInfo, nil

}

func GetOneUsageById(id string) ([]UsageNewInfo, error) {
	sql := "select a.*,group_concat(c.UserName) as Admin from (SELECT * from sys_usage_new where DeletedAt is NULL and Id = " + id + ")a left join  (select UsageId,OAId from sys_owner_usage  where DeletedAt is NULL and IsManager=1)b on a.Id =b.UsageId left join  (select Id,UserName,NameCn from sys_oa  where deletedAt is NULL) c on b.OAId=c.Id group by Id"
	//var usageInfo []UsageNewInfo
	usageInfo := make([]UsageNewInfo, 0)
	err := Engine.Sql(sql).Find(&usageInfo)
	logs.Info(sql)

	if err != nil {
		logs.Error(err)
		return nil, err
	}

	return usageInfo, nil
}
func AddUsage(usageNew *UsageNew) (*SysUsageNew, error) {
	var v SysUsageNew

	v.Usage = usageNew.Usage
	v.SecUsage = usageNew.SecUsage
	v.HierarchyLevel = usageNew.HierarchyLevel
	v.ResourceNum = usageNew.ResourceNum
	v.LastUpdatePerson = usageNew.LastUpdatePerson
	v.Remark = usageNew.Remark

	has, err := Engine.Exist(&SysUsageNew{
		Usage:          usageNew.Usage,
		SecUsage:       usageNew.SecUsage,
		HierarchyLevel: usageNew.HierarchyLevel,
	})

	if err != nil {
		logs.Error(err)
		return nil, err
	}
	if has {
		return nil, errors.New("重复插入，记录已存在")
	}
	row, err := Engine.Insert(&v)
	if err != nil {
		logs.Error(err)
		return nil, err
	}
	if row == 0 {
		return nil, errors.New("插入失败")
	}
	//user := strings.Split(usageNew.Admin, ",")
	if len(usageNew.Admin) == 0 {
		return nil, errors.New("请添加管理员")
	}
	for i := 0; i < len(usageNew.Admin); i++ {
		//var id int
		var owner SysOwnerUsage
		// has, err := Engine.Table("sys_oa").Where("UserName=?", usageNew.Admin[i]).Cols("Id").Get(&id)
		// if err != nil {
		// 	continue
		// 	logs.Info(err)
		// }
		// if !has {
		// 	logs.Error("没有查到对应管理员人员", usageNew.Admin[i])
		// 	continue
		// }
		owner.OAId = usageNew.Admin[i]
		owner.UsageId = v.Id
		owner.IsManager = 1
		row, err := Engine.Insert(&owner)
		if err != nil {
			continue
			logs.Error(err)

		}
		if row == 0 {
			continue
			logs.Error("插入失败")
		}
	}
	//}
	return &v, nil
}

func UpdateUsage(usageNew *UsageNew) error {
	row, err := Engine.Where("DeletedAt is NULL").ID(usageNew.Id).AllCols().Update(usageNew)
	if err != nil {
		logs.Error(err)
		return err
	}
	if row == 0 {
		return errors.New("更新失败")
	}
	return nil
}

type Usage struct {
	Id               int      `json:"Id"`
	Usage            string   `json:"Usage"`
	SecUsage         string   `json:"SecUsage"`
	HierarchyLevel   int      `json:"HierarchyLevel"`
	ResourceNum      int      `json:"ResourceNum"`
	Admin            []string `json:"Admin"`
	LastUpdatePerson string   `json:"LastUpdatePerson"`
	Remark           string   `json:"Remark"`
}

func GetAllUsageTest(usageName string, level int) ([]Usage, error) {
	usage := make([]SysUsageNew, 0)
	err := Engine.Find(&usage)
	if err != nil {
		logs.Error(err)
	}
	usageNew := make([]Usage, 0)
	for i := 0; i < len(usage); i++ {
		var temp Usage
		temp.HierarchyLevel = usage[i].HierarchyLevel
		temp.Id = usage[i].Id
		temp.Usage = usage[i].Usage
		temp.SecUsage = usage[i].SecUsage

		err := Engine.Table("sys_owner_usage").Where("UsageId=? and IsManager = 1", usage[i].Id).Join("left", "sys_oa", "sys_owner_usage.OAId=sys_oa.Id").Cols("sys_oa.UserName").Find(&temp.Admin)
		if err != nil {
			logs.Error(err)
		}
		usageNew = append(usageNew, temp)

	}
	return usageNew, nil

}
