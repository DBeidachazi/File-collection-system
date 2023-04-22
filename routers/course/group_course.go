package course

import "github.com/gin-gonic/gin"

func InitCourseRouter(courseGroup *gin.RouterGroup) {
	InitCourseCreate(courseGroup)
	InitCourseDelete(courseGroup)
}
