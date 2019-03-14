package helper

import (
	"os"

	"encoding/json"
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

//日志变量
var Logger = logs.NewLogger()

//日志初始化

func InitLogs() {
	if _, err := os.Stat("logs"); err != nil {
		os.Mkdir("logs", os.ModePerm)
	}
	config := make(map[string]interface{})
	//mode := beego.BConfig.RunMode

	maxLines, err := beego.AppConfig.Int64("logs::max_lines")
	if err != nil {
		maxLines = 20000
	}
	maxSize, err := beego.AppConfig.Int64("logs::max_size")
	if err != nil {
		maxSize = 102400
	}
	maxDays, err := beego.AppConfig.Int64("logs::max_days")
	if err != nil {
		maxDays = 7
	}

	daily, err := beego.AppConfig.Bool("logs::daily")
	if err != nil {
		daily = true
	}
	level, error := beego.AppConfig.Int64(beego.BConfig.RunMode + "::logs_level")
	if error != nil {
		level = 4
		if Debug {
			level = 7
		}
	}

	config["filename"] = beego.AppConfig.String("logs::log_path")
	config["level"] = level
	config["daily"] = daily
	config["maxlines"] = maxLines
	config["maxsize"] = maxSize
	config["maxdays"] = maxDays
	// map 转 json
	configStr, err := json.Marshal(config)
	if err != nil {
		fmt.Println("initLogger failed, marshal err:", err)
		return
	}

	beego.SetLogger(logs.AdapterFile, string(configStr))
	beego.SetLogFuncCall(true)
	fmt.Println(string(configStr))
}

// func InitLogs() {
// 	//创建日志目录
// 	if _, err := os.Stat("logs"); err != nil {
// 		os.Mkdir("logs", os.ModePerm)
// 	}
// 	var level = 7
// 	if Debug {
// 		level = 4
// 	}
// 	maxLines := GetConfigInt64("logs", "max_lines")
// 	if maxLines <= 0 {
// 		maxLines = 10000
// 	}
// 	maxDays := GetConfigInt64("logs", "max_days")
// 	if maxDays <= 0 {
// 		maxDays = 7
// 	}
// 	//初始化日志各种配置
// 	LogsConf := fmt.Sprintf(`{"filename":"logs/dochub.log","level":%v,"maxlines":%v,"maxsize":120,"daily":true,"maxdays":%v}`, level, maxLines, maxDays)
// 	Logger.SetLogger(logs.AdapterFile, LogsConf)
// 	if Debug {
// 		Logger.SetLogger("console")
// 		beego.Info("日志配置信息：" + LogsConf)
// 	} else {
// 		//是否异步输出日志
// 		Logger.Async(1e3)
// 	}
// 	Logger.EnableFuncCallDepth(true) //是否显示文件和行号
// }
