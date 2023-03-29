package course_api

import (
	"FengfengStudy/global"
	"FengfengStudy/models/orm/dal"
	"FengfengStudy/models/res"
	"github.com/gin-gonic/gin"
)

type CourseDeleteRequest struct {
	CourseId int32 `json:"course_id" bind:"required"`
}

func (CourseApi) CourseDeleteView(c *gin.Context) {
	var req CourseDeleteRequest
	err := c.ShouldBindJSON(&req)
	global.Log.Warnln(err)
	global.Log.Warnln(req)
	if returnErr(err, c)() {
		return
	}

	_, err = dal.Course.Where(dal.Course.CourseID.Eq(req.CourseId)).First()
	if returnErr(err, c)() {
		return
	}

	_, err = dal.Course.Where(dal.Course.CourseID.Eq(req.CourseId)).Delete()
	if returnErr(err, c)() {
		return
	}

	_, err = dal.Role.Where(dal.Role.CourseID.Eq(req.CourseId)).Delete()
	if returnErr(err, c)() {
		return
	}

	_, err = dal.Worklist.Where(dal.Worklist.CourseID.Eq(req.CourseId)).Delete()
	if returnErr(err, c)() {
		return
	}

	global.Log.Infoln("删除课程成功: ", req.CourseId)
	res.OkWithMessage("删除课程成功", c)
}

func returnErr(err error, c *gin.Context) func() bool {
	return func() bool {
		if err != nil {
			global.Log.Warnln("删除课程失败", err)
			res.FailWithMessage("删除课程失败", c)
			return true
		}
		return false
	}
}
