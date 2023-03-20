package routers

import (
	"FengfengStudy/api"
	"github.com/gin-gonic/gin"
)

func InitSearchRole(router *gin.Engine) {
	SearchRoleAPi := api.ApiGroupApp.SearchRoleApi
	router.POST("/api/searchrole", SearchRoleAPi.SearchRoleView)
}
