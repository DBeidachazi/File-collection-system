package routers

import (
	"FengfengStudy/api"
	"github.com/gin-gonic/gin"
)

func InitFileDownload(router *gin.Engine) {
	fileDownloadAPi := api.ApiGroupApp.FileDownloadApi
	router.POST("/api/filedownload", fileDownloadAPi.FileDownloadApi)
}
