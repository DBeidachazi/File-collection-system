package course

import (
	"FengfengStudy/api"
	"github.com/gin-gonic/gin"
)

func InitCourseDelete(courseGroup *gin.RouterGroup) {
	courseApi := api.ApiGroupApp.CourseApi
	courseGroup.POST("/coursedelete", courseApi.CourseDeleteView)
}
