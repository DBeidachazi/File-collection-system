package course_api

import (
	"FengfengStudy/global"
	"FengfengStudy/models/orm/dal"
	"FengfengStudy/models/orm/model"
	"FengfengStudy/models/res"
	"github.com/gin-gonic/gin"
	"time"
)

type CourseCreateRequest struct {
	StuId      int32  `json:"stuId" bind:"required"`
	CourseName string `json:"courseName" bind:"required"`
	ClassId    int32  `json:"classId" bind:"required"`
	Username   string `json:"username" bind:"required"`
	Deadline   string `json:"deadline" bind:"required"` // 截止时间 "2020-12-12 12:12:12"
}

func (CourseApi) CreateCourseView(c *gin.Context) {
	var req CourseCreateRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		global.Log.Warnln("参数错误", err)
		res.FailWithMessage("参数错误", c)
		return
	}
	// 查看用户是否在班级中
	userIsInClass, err := dal.ClassUser.Where(dal.ClassUser.StuID.Eq(req.StuId), dal.ClassUser.ClassID.Eq(req.ClassId)).First()
	if err != nil {
		global.Log.Warnln("用户不在班级中", err)
		res.FailWithMessage("用户不在班级中", c)
		return
	}
	// 查看用户是否有权限创建课程
	if userIsInClass.PermissionType != 1 && userIsInClass.PermissionType != 2 {
		global.Log.Warnln("用户没有权限创建课程", req.StuId, req.ClassId)
		res.FailWithMessage("用户没有权限创建课程", c)
		return
	}
	// 入库
	t, err := time.Parse("2006-01-02 15:04:05", req.Deadline)
	if err != nil {
		global.Log.Warnln("时间格式有误", t, "+", err, "+", req.Deadline)
		res.FailWithMessage("时间格式有误", c)
		return
	}
	course := model.Course{
		CourseName: req.CourseName,
		ClassID:    req.ClassId,
		Username:   req.Username,
		Deadline:   t,
		StuID:      req.StuId,
	}
	err = dal.Course.Create(&course)
	if err != nil {
		global.Log.Warnln("入库失败", err)
		res.FailWithMessage("创建课程失败", c)
		return
	}
	global.Log.Infoln(req.StuId, req.CourseName, "创建课程成功")
	// 向role中插入 该班级的所有学生的这门课程数据
	InsertIntoRole(&course, c)
	res.OkWithMessage("创建课程成功", c)
	return
}
func InsertIntoRole(course *model.Course, c *gin.Context) {
	findAllUser, err := dal.ClassUser.Where(dal.ClassUser.ClassID.Eq(course.ClassID)).Find()
	if err != nil {
		return
	}
	for _, v := range findAllUser {
		var status int32
		if v.PermissionType == 2 {
			status = 3 // 发布者
		} else {
			status = 2 // 未提交
		}
		// 构建role
		role := model.Role{
			StuID:    v.StuID,
			RoleName: course.CourseName,
			ClassID:  v.ClassID,
			Status:   status,
		}
		err = dal.Role.Create(&role)
		if err != nil {
			global.Log.Warnln("role入库失败", err)
			res.FailWithMessage("role入库失败", c)
			return
		}
	}
	global.Log.Infoln("role入库成功")
}
