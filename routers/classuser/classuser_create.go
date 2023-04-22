package classuser

import (
	"FengfengStudy/api"
	"github.com/gin-gonic/gin"
)

func InitClassUserCreate(classUserRouterGroup *gin.RouterGroup) {
	classUserApi := api.ApiGroupApp.ClassUserApi
	classUserRouterGroup.POST("/classusercreate", classUserApi.CreateClassUserView)
}
