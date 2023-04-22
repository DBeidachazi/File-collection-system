package file

import "github.com/gin-gonic/gin"

func InitFileRouter(fileGroup *gin.RouterGroup) {
	InitFileDownload(fileGroup)
	InitFileUpload(fileGroup)
}
