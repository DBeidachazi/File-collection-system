package routers

import (
	"FengfengStudy/api"
	"github.com/gin-gonic/gin"
)

func InitClassCreate(router *gin.Engine) {
	classCreateApi := api.ApiGroupApp.ClassApi
	router.POST("/api/classcreate", classCreateApi.CreateClassView)
}
