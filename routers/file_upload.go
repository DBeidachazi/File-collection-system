package routers

import (
	"FengfengStudy/api"
	"github.com/gin-gonic/gin"
)

func InitFileUpload(router *gin.Engine) {
	fileUploadApi := api.ApiGroupApp.FileUploadApi
	router.POST("/upload/:classid/:stuid/:courseid", fileUploadApi.FileUploadView)
}
