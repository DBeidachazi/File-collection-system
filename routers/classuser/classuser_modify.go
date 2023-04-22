package classuser

import (
	"FengfengStudy/api"
	"github.com/gin-gonic/gin"
)

func InitClassUserModify(classUserRouterGroup *gin.RouterGroup) {
	classUserApi := api.ApiGroupApp.ClassUserApi
	classUserRouterGroup.POST("/classusermodify", classUserApi.ModifyClassUserView)
}
