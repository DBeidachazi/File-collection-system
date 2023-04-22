package class

import (
	"FengfengStudy/api"
	"github.com/gin-gonic/gin"
)

func InitClassDelete(classRouter *gin.RouterGroup) {
	classDeleteApi := api.ApiGroupApp.ClassApi
	classRouter.POST("classdelete", classDeleteApi.DeleteClassView)
}
