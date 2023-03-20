package routers

import (
	"FengfengStudy/api"
	"github.com/gin-gonic/gin"
)

func InitUserLoginRouter(router *gin.Engine) {
	userLoginApi := api.ApiGroupApp.UserApi
	router.POST("/api/userlogin", userLoginApi.UserLoginView)
	//router.POST("/api/userlogin", func(c *gin.Context) {
	//	c.JSON(200, gin.H{"message": "Hello World"})
	//})
}
