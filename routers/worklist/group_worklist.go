package worklist

import "github.com/gin-gonic/gin"

func InitWorklistRouter(userGroup *gin.RouterGroup) {
	InitWorklistSearch(userGroup)
}
