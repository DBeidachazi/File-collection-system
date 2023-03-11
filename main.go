package main

import (
	"FengfengStudy/core"
	"FengfengStudy/global"
	"FengfengStudy/routers"
)

func main() {
	// 读取配置文件
	core.InitConf()
	// 初始化日志
	global.Log = core.InitLogger()
	// 生成数据库代码 + 连接数据库
	global.DB = core.BuildGen()

	router := routers.InitRouter()

	addr := global.Config.System.GetHost()
	global.Log.Infof("server run at %s", addr)
	err := router.Run(addr)
	if err != nil {
		global.Log.Errorln(err)
		return
	}

}
