package worklist

import (
	"FengfengStudy/api"
	"github.com/gin-gonic/gin"
)

func InitWorklistSearch(userGroup *gin.RouterGroup) {
	worklistApi := api.ApiGroupApp.WorklistApi
	userGroup.POST("/searchworklist", worklistApi.WorklistSearchView)
}
