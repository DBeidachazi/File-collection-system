package routers

import (
	"FengfengStudy/api"
	"github.com/gin-gonic/gin"
)

func InitWorklistRouter(router *gin.Engine) {
	worklistApi := api.ApiGroupApp.WorklistApi
	router.POST("/api/searchworklist", worklistApi.WorklistSearchView)
}
