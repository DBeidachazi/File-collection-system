package routers

import (
	"FengfengStudy/api"
	"github.com/gin-gonic/gin"
)

func InitCourseCreate(router *gin.Engine) {
	courseApi := api.ApiGroupApp.CourseApi
	router.POST("/api/coursecreate", courseApi.CreateCourseView)
}
