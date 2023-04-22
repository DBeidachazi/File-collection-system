package classuser

import (
	"FengfengStudy/api"
	"github.com/gin-gonic/gin"
)

func InitClassUserDelete(classUserRouterGroup *gin.RouterGroup) {
	classUserApi := api.ApiGroupApp.ClassUserApi
	classUserRouterGroup.POST("/classuserdelete", classUserApi.ClassUserDeleteView)
}
