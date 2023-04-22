package user

import "github.com/gin-gonic/gin"

func InitUserRouter(userGroup *gin.RouterGroup) {
	InitUserLogin(userGroup)
	InitUserRegister(userGroup)
}
