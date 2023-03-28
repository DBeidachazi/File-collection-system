package routers

import (
	"FengfengStudy/api"
	"github.com/gin-gonic/gin"
)

func InitClassSearchUser(c *gin.Engine) {
	classSearchUserApi := api.ApiGroupApp.ClassUserApi
	c.POST("/api/classsearchuser", classSearchUserApi.SearchCLassUserView)
}
