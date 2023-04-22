package routers

import (
	"FengfengStudy/global"
	"FengfengStudy/routers/class"
	"FengfengStudy/routers/classuser"
	"FengfengStudy/routers/course"
	"FengfengStudy/routers/file"
	"FengfengStudy/routers/role"
	"FengfengStudy/routers/user"
	"FengfengStudy/routers/worklist"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	gin.SetMode(global.Config.System.Env)
	router := gin.Default()
	// 中间件 not use
	// router.Use(middleware.Cors())
	router.MaxMultipartMemory = global.Config.FileSize.Size << 20

	// 路由分组
	routerGroup := router.Group("/api")
	class.InitClassRouter(routerGroup)
	user.InitUserRouter(routerGroup)
	classuser.InitClassUserRouter(routerGroup)
	course.InitCourseRouter(routerGroup)
	role.InitRoleRouter(routerGroup)
	file.InitFileRouter(routerGroup)
	worklist.InitWorklistRouter(routerGroup)

	// 系统配置api
	// InitSettingsRouter(router)
	// InitUserLogin(router)
	// InitUserRegister(router)

	// InitClassSearch(router)
	// InitClassDelete(router)

	// InitFileUpload(router)
	// InitFileDownload(router)
	// InitRoleSearch(router)
	// InitWorklistRouter(router)
	//
	// InitClassUserCreate(router)
	// InitClassUserDelete(router)
	// InitClassUserModify(router)
	// InitClassUserSearch(router)
	//
	// InitCourseCreate(router)
	// InitCourseDelete(router)

	return router
}
