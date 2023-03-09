package main

import (
	"FengfengStudy/core"
	"FengfengStudy/global"
)

func main() {
	// 读取配置文件
	core.InitConf()
	// 初始化日志
	global.Log = core.InitLogger()

	// 生成数据库代码 + 连接数据库
	global.DB = core.BuildGen()

}
