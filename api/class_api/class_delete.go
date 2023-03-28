package class_api

import (
	"FengfengStudy/global"
	"FengfengStudy/models/orm/dal"
	"FengfengStudy/models/res"
	"github.com/gin-gonic/gin"
)

type ClassDeleteRequest struct {
	ClassId int32 `json:"class_id" bind:"required"`
}

func (ClassApi) DeleteClassView(c *gin.Context) {
	var req ClassDeleteRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		global.Log.Warnln("绑定参数失败", err)
		return
	}
	// search class
	_, err = dal.Class.Where(dal.Class.ClassID.Eq(req.ClassId)).First()
	if err != nil {
		global.Log.Warnln("班级不存在", req.ClassId)
		res.FailWithMessage("班级不存在", c)
		return
	}
	// delete class
	_, err = dal.Class.Where(dal.Class.ClassID.Eq(req.ClassId)).Delete()
	if err != nil {
		global.Log.Warnln("删除班级失败", req.ClassId)
		return
	}
	// delete class_user
	_, err = dal.ClassUser.Where(dal.ClassUser.ClassID.Eq(req.ClassId)).Delete()
	if err != nil {
		global.Log.Warnln("删除班级成员失败", req.ClassId)
		return
	}
	// delete role
	_, err = dal.Role.Where(dal.Role.ClassID.Eq(req.ClassId)).Delete()
	if err != nil {
		global.Log.Warnln("删除任务失败", req.ClassId)
		return
	}
	// search worklist_course_id
	worklistCourseId, err := dal.Course.Where(dal.Course.ClassID.Eq(req.ClassId)).Find()
	if err != nil {
		global.Log.Warnln("查询课程失败", req.ClassId)
		return
	}
	// delete course
	_, err = dal.Course.Where(dal.Course.ClassID.Eq(req.ClassId)).Delete()
	if err != nil {
		global.Log.Warnln("删除课程失败", req.ClassId)
		return
	}
	for _, v := range worklistCourseId {
		// delete worklist
		_, err = dal.Worklist.Where(dal.Worklist.CourseID.Eq(v.CourseID)).Delete()
		if err != nil {
			global.Log.Warnln("删除作业失败", v.CourseID)
			return
		}
	}
	res.OkWithMessage("删除班级成功", c)

}
