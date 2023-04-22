package classuser

import (
	"FengfengStudy/api"
	"github.com/gin-gonic/gin"
)

func InitClassUserSearch(classUserRouterGroup *gin.RouterGroup) {
	classSearchUserApi := api.ApiGroupApp.ClassUserApi
	classUserRouterGroup.POST("/classsearchuser", classSearchUserApi.SearchCLassUserView)
}
