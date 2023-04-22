package user

import (
	"FengfengStudy/api"
	"github.com/gin-gonic/gin"
)

func InitUserRegister(userGroup *gin.RouterGroup) {
	userRegisterApi := api.ApiGroupApp.UserApi
	userGroup.POST("/userregister", userRegisterApi.UserRegisterView)
}
