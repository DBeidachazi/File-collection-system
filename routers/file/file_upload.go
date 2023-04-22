package file

import (
	"FengfengStudy/api"
	"github.com/gin-gonic/gin"
)

func InitFileUpload(fileGroup *gin.RouterGroup) {
	fileUploadApi := api.ApiGroupApp.FileUploadApi
	fileGroup.POST("/upload/:classid/:stuid/:courseid", fileUploadApi.FileUploadView)
}
