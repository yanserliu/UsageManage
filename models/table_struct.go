package models

//"time"

//SameMapper 支持结构体名称和对应的表名称相同
//以及结构体field名称与对应的表字段名称相同的命名；

type SysUsageNew struct {
	Id               int    `json:"Id" xorm:"not null pk autoincr INT(11)"`
	Usage            string `json:"Usage" xorm:"not null VARCHAR(255)"`
	SecUsage         string `json:"SecUsage" xorm:"VARCHAR(255)"`
	HierarchyLevel   int    `json:"HierarchyLevel" xorm:"INT(11)"`
	ResourceNum      int    `json:"ResourceNum" xorm:"default 0 INT(11)"`
	CreatedAt        string `json:"CreatedAt" xorm:"created DATETIME"`
	DeletedAt        string `json:"DeletedAt" xorm:"deleted DATETIME"`
	LastUpdatePerson string `json:"LastUpdatePerson" xorm:"VARCHAR(255)"`
	LastUpdateTime   string `json:"LastUpdateTime" xorm:"updated DATETIME"`
	Remark           string `json:"Remark" xorm:"VARCHAR(255)"`
}

type SysOwnerUsage struct {
	Id               int    `json:"Id" xorm:"not null pk autoincr unique INT(11)"`
	OAId             int    `json:"OAId" xorm:"index INT(11)"`
	UsageId          int    `json:"UsageId" xorm:"index INT(11)"`
	IsManager        int    `json:"IsManager" xorm:"INT(11) default 0"`
	CreatedAt        string `json:"CreatedAt" xorm:"created DATETIME"`
	LastUpdateTime   string `json:"LastUpdateTime" xorm:"updated updated DATETIME"`
	DeletedAt        string `json:"DeletedAt" xorm:"deleted DATETIME"`
	LastUpdatePerson string `json:"LastUpdatePerson" xorm:"VARCHAR(255)"`
}

type SysOa struct {
	Id               int    `json:"Id" xorm:"not null pk unique INT(11)"`
	UserName         string `json:"UserName" xorm:"index VARCHAR(255)"`
	JobNumber        int64  `json:"JobNumber" xorm:"BIGINT(11)"`
	BossId           int    `json:"BossId" xorm:"INT(11)"`
	Email            string `json:"Email" xorm:"VARCHAR(255)"`
	Namecn           string `json:"NameCn" xorm:"index VARCHAR(255)"`
	MemberState      int    `json:"MemberState" xorm:"INT(11)"`
	Job              string `json:"Job" xorm:"VARCHAR(255)"`
	Position         int    `json:"Position" xorm:"INT(11)"`
	Leader           int    `json:"Leader" xorm:"INT(11)"`
	PhoneNumber      string `json:"PhoneNumber" xorm:"VARCHAR(255)"`
	ExtraEmail       string `json:"ExtraEmail" xorm:"VARCHAR(255)"`
	CXO              int    `json:"CXO" xorm:"INT(11)"`
	HRBP             int    `json:"HRBP" xorm:"INT(11)"`
	DepartmentCn     string `json:"DepartmentCN" xorm:"VARCHAR(255)"`
	CreatedAt        string `json:"CreatedAt" xorm:"created DATETIME"`
	LastUpdateTime   string `json:"LastUpdateTime" xorm:"updated DATETIME"`
	LastUpdatePerson string `json:"LastUpdatePerson" xorm:"DATETIME"`
	DeletedAt        string `json:"DeletedAt" xorm:"deleted DATETIME"`
}

type SysSuperUser struct {
	Id   int    `json:"id" xorm:" not null pk autoincr INT(11)"`
	Name string `json:"name" xorm:"VARCHAR(255)"`
	OAId int    `json:"OAId" xorm:"INT(11)"`
}

// type VMServerOwner struct {
// 	Id               int       `json:"Id" xorm:"'id' not null pk autoincr INT(11)"`
// 	UUID             string    `json:"UUID" xorm:"not null VARCHAR(255)"`
// 	Usage            string    `json:"Usage" xorm:"VARCHAR(255)"`
// 	SecUsage         string    `json:"SecUsage" xorm:"VARCHAR(255)"`
// 	Level            int       `json:"Level" xorm:"'level' default 1 INT(11)"`
// 	Owner            string    `json:"Owner" xorm:"VARCHAR(255)"`
// 	ResourceNum      int       `json:"ResourceNum" xorm:"default 0 INT(11)"`
// 	AssetFlag        int       `json:"AssetFlag" xorm:"TINYINT(1)"`
// 	Remark           string    `json:"Remark" xorm:"VARCHAR(255)"`
// 	LastUpdatePerson string    `json:"LastUpdatePerson" xorm:"VARCHAR(255)"`
// 	Tags             string    `json:"tags" xorm:"VARCHAR(255)"`
// 	CreatedAt        time.Time `json:"createdAt" xorm:"'createdAt' created not null DATETIME"`
// 	LastUpdateTime   time.Time `json:"LastUpdateTime" xorm:"updated DATETIME"`
// 	DeletedAt        time.Time `json:"deletedAt" xorm:"'deletedAt' deleted DATETIME"`
// 	CategoryId       int       `json:"categoryId" xorm:"'categoryId' index INT(11)"`
// }

// type ServerOwner struct {
// 	Id               int       `json:"id" xorm:"'id' not null pk autoincr INT(11)"`
// 	UUID             string    `json:"UUID" xorm:"not null VARCHAR(255)"`
// 	Usage            string    `json:"Usage" xorm:"VARCHAR(255)"`
// 	SecUsage         string    `json:"SecUsage" xorm:"VARCHAR(255)"`
// 	Level            int       `json:"level" xorm:"'level' default 1 INT(11)"`
// 	Owner            string    `json:"Owner" xorm:"VARCHAR(255)"`
// 	AssetFlag        int       `json:"AssetFlag" xorm:"TINYINT(1)"`
// 	Remark           string    `json:"Remark" xorm:"VARCHAR(255)"`
// 	Lastupdateperson string    `json:"LastUpdatePerson" xorm:"VARCHAR(255)"`
// 	Tags             string    `json:"tags" xorm:"'tags' VARCHAR(255)"`
// 	CreatedAt        time.Time `json:"createdAt" xorm:"'createdAt' created not null DATETIME"`
// 	LastUpdateTime   time.Time `json:"LastUpdateTime" xorm:"updated not null DATETIME"`
// 	DeletedAt        time.Time `json:"deletedAt" xorm:"'deletedAt' deleted DATETIME"`
// 	CategoryId       int       `json:"categoryId" xorm:"'categoryId' index INT(11)"`
// 	ResourceNum      int       `json:"ResourceNum" xorm:"default 0 INT(11)"`
// }
