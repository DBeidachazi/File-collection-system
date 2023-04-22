package role

import (
	"FengfengStudy/api"
	"github.com/gin-gonic/gin"
)

func InitRoleSearch(roleGroup *gin.RouterGroup) {
	SearchRoleAPi := api.ApiGroupApp.SearchRoleApi
	roleGroup.POST("/searchrole", SearchRoleAPi.SearchRoleView)
}
