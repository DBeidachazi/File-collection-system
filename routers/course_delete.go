package routers

import (
	"FengfengStudy/api"
	"github.com/gin-gonic/gin"
)

func InitCourseDelete(router *gin.Engine) {
	courseApi := api.ApiGroupApp.CourseApi
	router.POST("/api/coursedelete", courseApi.CourseDeleteView)
}
