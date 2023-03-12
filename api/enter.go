package api

import (
	"FengfengStudy/api/class_api"
	"FengfengStudy/api/classuser_api"
	"FengfengStudy/api/course_api"
	"FengfengStudy/api/setting_api"
	"FengfengStudy/api/user_api"
)

type ApiGroup struct {
	SettingsApi  setting_api.SettingsApi
	UserApi      user_api.UserApi
	ClassApi     class_api.ClassApi
	ClassUserApi classuser_api.ClassUserApi
	CourseApi    course_api.CourseApi
}

var ApiGroupApp = new(ApiGroup)
