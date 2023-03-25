package routers

import (
	"FengfengStudy/api"
	"github.com/gin-gonic/gin"
)

func InitClassSearch(router *gin.Engine) {
	classSearchAPi := api.ApiGroupApp.ClassApi
	router.POST("/api/classsearch", classSearchAPi.SearchClassView)
}
