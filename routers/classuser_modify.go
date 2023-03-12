package routers

import (
	"FengfengStudy/api"
	"github.com/gin-gonic/gin"
)

func InitClassUserModify(router *gin.Engine) {
	classUserApi := api.ApiGroupApp.ClassUserApi
	router.POST("/api/classusermodify", classUserApi.ModifyClassUserView)
}
