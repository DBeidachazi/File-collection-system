package class

import (
	"github.com/gin-gonic/gin"
)

func InitClassRouter(classRouterGroup *gin.RouterGroup) {
	InitClassCreate(classRouterGroup)
	InitClassSearch(classRouterGroup)
	InitClassDelete(classRouterGroup)
}
