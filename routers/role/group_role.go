package role

import "github.com/gin-gonic/gin"

func InitRoleRouter(roleGroup *gin.RouterGroup) {
	InitRoleSearch(roleGroup)
}
