package routers

import (
	"FengfengStudy/api"
	"github.com/gin-gonic/gin"
)

func InitClassUserDelete(router *gin.Engine) {
	classUserApi := api.ApiGroupApp.ClassUserApi
	router.POST("/api/classuserdelete", classUserApi.ClassUserDeleteView)
}
