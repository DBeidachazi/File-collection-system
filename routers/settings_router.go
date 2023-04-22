package routers

import (
	"FengfengStudy/api"
	"github.com/gin-gonic/gin"
)

// @Description: 测试
func InitSettingsRouter(router *gin.Engine) {
	settingsApi := api.ApiGroupApp.SettingsApi
	router.GET("/api/settings", settingsApi.SettingsInfoView)
}
