package course

import (
	"FengfengStudy/api"
	"github.com/gin-gonic/gin"
)

func InitCourseCreate(courseGroup *gin.RouterGroup) {
	courseApi := api.ApiGroupApp.CourseApi
	courseGroup.POST("/coursecreate", courseApi.CreateCourseView)
}
