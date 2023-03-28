package routers

import (
	"FengfengStudy/api"
	"github.com/gin-gonic/gin"
)

func InitClassDelete(c *gin.Engine) {
	classDeleteApi := api.ApiGroupApp.ClassApi
	c.POST("/api/classdelete", classDeleteApi.DeleteClassView)
}
