package routers

import (
	"FengfengStudy/api"
	"github.com/gin-gonic/gin"
)

func InitUserLoginRouter(router *gin.Engine) {
	userLoginApi := api.ApiGroupApp.UserApi
	router.POST("/api/userlogin", userLoginApi.UserLoginView)
}
