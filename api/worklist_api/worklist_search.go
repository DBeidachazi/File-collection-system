package worklist_api

import (
	"FengfengStudy/global"
	"FengfengStudy/models/orm/dal"
	"FengfengStudy/models/res"
	"github.com/gin-gonic/gin"
	"time"
)

type WorklistRequest struct {
	StuId int32 `json:"stu_id" binding:"required"`
}

func (WorklistApi) WorklistSearchView(c *gin.Context) {
	var req WorklistRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		global.Log.Warnln("绑定参数失败")
		global.Log.Warnln(err)
		res.FailWithMessage("绑定参数失败", c)
		return
	}
	_, err = dal.Worklist.Where(dal.Worklist.StuID.Eq(req.StuId)).First()
	if err != nil {
		global.Log.Warnln("没有worklist: ", req.StuId)
		res.FailWithMessage("没有发布任务", c)
		return
	}
	findAllWorklist, err := dal.Worklist.Where(dal.Worklist.StuID.Eq(req.StuId)).Find()
	if err != nil {
		global.Log.Warnln("查询工作列表失败: ", req.StuId)
		global.Log.Warnln(err)
		res.FailWithMessage("没有发布任务", c)
		return
	}

	var data []returnWorklist
	for _, v := range findAllWorklist {
		var returnWorklistOne returnWorklist
		returnWorklistOne.WorkID = v.WorkID
		returnWorklistOne.CourseID = v.CourseID
		returnWorklistOne.Status = v.Status
		returnWorklistOne.StuID = v.StuID
		returnWorklistOne.Deadline = v.Deadline
		returnWorklistOne.CourseName = v.CourseName

		classId, err := dal.Course.Where(dal.Course.CourseID.Eq(v.CourseID)).First()
		if err != nil {
			global.Log.Warnln("查询任务失败: ", v.CourseID)
			global.Log.Errorln(err)
			res.FailWithMessage("查询任务失败", c)
			return
		}
		returnWorklistOne.ClassId = classId.ClassID

		data = append(data, returnWorklistOne)
	}

	global.Log.Infoln("查询工作列表成功: ", req.StuId)
	res.Ok(data, "查询工作列表成功", c)
}

type returnWorklist struct {
	WorkID     int32     `gorm:"column:work_id;primaryKey;autoIncrement:true" json:"work_id"` // 工作流水号
	CourseID   int32     `gorm:"column:course_id;not null" json:"course_id"`                  // 任务id
	CourseName string    `gorm:"column:course_name;not null" json:"course_name"`              // 任务名称
	Status     int32     `gorm:"column:status;not null" json:"status"`                        // 提交人数
	StuID      int32     `gorm:"column:stu_id;not null" json:"stu_id"`
	Deadline   time.Time `gorm:"column:deadline;not null" json:"deadline"`
	ClassId    int32     `gorm:"column:class_id;not null" json:"class_id"`
}
