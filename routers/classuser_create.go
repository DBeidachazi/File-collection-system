package routers

import (
	"FengfengStudy/api"
	"github.com/gin-gonic/gin"
)

func InitClassUserCreate(router *gin.Engine) {
	classUserApi := api.ApiGroupApp.ClassUserApi
	router.POST("/api/classusercreate", classUserApi.CreateClassUserView)
}
