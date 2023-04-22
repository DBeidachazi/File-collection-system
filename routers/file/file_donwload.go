package file

import (
	"FengfengStudy/api"
	"github.com/gin-gonic/gin"
)

func InitFileDownload(fileGroup *gin.RouterGroup) {
	fileDownloadAPi := api.ApiGroupApp.FileDownloadApi
	fileGroup.POST("/filedownload", fileDownloadAPi.FileDownloadApi)
}
