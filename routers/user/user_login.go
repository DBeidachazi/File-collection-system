package user

import (
	"FengfengStudy/api"
	"github.com/gin-gonic/gin"
)

func InitUserLogin(userGroup *gin.RouterGroup) {
	userLoginApi := api.ApiGroupApp.UserApi
	userGroup.POST("/userlogin", userLoginApi.UserLoginView)
}
