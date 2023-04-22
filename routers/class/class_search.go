package class

import (
	"FengfengStudy/api"
	"github.com/gin-gonic/gin"
)

func InitClassSearch(classRouter *gin.RouterGroup) {
	classSearchAPi := api.ApiGroupApp.ClassApi
	classRouter.POST("classsearch", classSearchAPi.SearchClassView)
}
