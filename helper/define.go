//定义一些常量和变量
package helper

import (
	"sync"

	"github.com/astaxie/beego"
	"github.com/huichen/sego"
)

const (
	//DocHub Version
	VERSION = "v2.1"
	//Cache Config
	CACHE_CONF = `{"CachePath":"./cache/runtime","FileSuffix":".cache","DirectoryLevel":2,"EmbedExpiry":120}`

	DEFAULT_STATIC_EXT    = ".txt,.html,.ico,.jpeg,.png,.gif,.xml"
	DEFAULT_COOKIE_SECRET = "dochub"

	//	扩展名
	EXT_CATE_WORD       = "word"
	EXT_NUM_WORD        = 1
	EXT_CATE_PPT        = "ppt"
	EXT_NUM_PPT         = 2
	EXT_CATE_EXCEL      = "excel"
	EXT_NUM_EXCEL       = 3
	EXT_CATE_PDF        = "pdf"
	EXT_NUM_PDF         = 4
	EXT_CATE_TEXT       = "text"
	EXT_NUM_TEXT        = 5
	EXT_CATE_OTHER      = "other"
	EXT_NUM_OTHER       = 6
	EXT_CATE_OTHER_MOBI = "mobi"
	EXT_CATE_OTHER_EPUB = "epub"
	EXT_CATE_OTHER_CHM  = "chm"
	EXT_CATE_OTHER_UMD  = "umd"
)

var (
	//develop mode
	Debug = beego.AppConfig.String("runmode") == "dev"

	//允许直接访问的文件扩展名
	StaticExt = make(map[string]bool)

	//分词器
	Segmenter sego.Segmenter

	//配置文件的全局map
	ConfigMap sync.Map

	//程序是否已经安装
	IsInstalled = false

	//允许上传的文档扩展名
	AllowedUploadExt = ",doc,docx,rtf,wps,odt,ppt,pptx,pps,ppsx,dps,odp,pot,xls,xlsx,et,ods,txt,pdf,chm,epub,umd,mobi,"
)
