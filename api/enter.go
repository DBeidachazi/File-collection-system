package api

import (
	"FengfengStudy/api/setting_api"
	"FengfengStudy/api/user_api"
)

type ApiGroup struct {
	SettingsApi setting_api.SettingsApi
	UserApi     user_api.UserApi
}

var ApiGroupApp = new(ApiGroup)
