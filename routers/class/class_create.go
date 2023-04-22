package class

import (
	"FengfengStudy/api"
	"github.com/gin-gonic/gin"
)

func InitClassCreate(classRouter *gin.RouterGroup) {
	classCreateApi := api.ApiGroupApp.ClassApi
	classRouter.POST("/classcreate", classCreateApi.CreateClassView)
}
