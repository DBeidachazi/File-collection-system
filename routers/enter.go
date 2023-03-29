package routers

import (
	"FengfengStudy/global"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	gin.SetMode(global.Config.System.Env)
	router := gin.Default()
	// 中间件
	//router.Use(middleware.Cors())
	router.MaxMultipartMemory = global.Config.FileSize.Size << 20
	// 系统配置api
	InitSettingsRouter(router)
	InitUserLoginRouter(router)
	InitUserRegisterRouter(router)
	InitClassCreate(router)
	InitClassUserCreate(router)
	InitClassUserModify(router)
	InitCourseCreate(router)
	InitFileUpload(router)
	InitFileDownload(router)
	InitSearchRole(router)
	InitWorklistRouter(router)
	InitClassSearch(router)
	InitClassDelete(router)
	InitClassSearchUser(router)
	InitClassUserDelete(router)
	InitCourseDelete(router)
	return router
}
