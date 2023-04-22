package classuser

import "github.com/gin-gonic/gin"

func InitClassUserRouter(classUserRouterGroup *gin.RouterGroup) {
	InitClassUserCreate(classUserRouterGroup)
	InitClassUserSearch(classUserRouterGroup)
	InitClassUserDelete(classUserRouterGroup)
	InitClassUserModify(classUserRouterGroup)
}
