package routers

import (
	"FengfengStudy/global"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	gin.SetMode(global.Config.System.Env)
	router := gin.Default()
	// 系统配置api
	InitSettingsRouter(router)
	InitUserLoginRouter(router)
	InitUserRegisterRouter(router)
	return router
}
