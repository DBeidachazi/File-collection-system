package routers

import (
	"FengfengStudy/api"
	"github.com/gin-gonic/gin"
)

func InitUserRegisterRouter(router *gin.Engine) {
	userRegisterApi := api.ApiGroupApp.UserApi
	router.POST("/api/userregister", userRegisterApi.UserRegisterView)
}
