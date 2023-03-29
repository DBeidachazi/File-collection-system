package searchRole_api

import (
	"FengfengStudy/global"
	"FengfengStudy/models/orm/dal"
	"FengfengStudy/models/res"
	"github.com/gin-gonic/gin"
	"time"
)

type SearchRoleRequest struct {
	StuId int32 `json:"stu_id" binding:"required"`
}

func (SearchRoleApi) SearchRoleView(c *gin.Context) {
	var req SearchRoleRequest
	err := c.ShouldBindJSON(&req)
	global.Log.Infoln(req.StuId)
	if err != nil {
		global.Log.Warnln("绑定参数失败")
		global.Log.Warnln(err)
		res.FailWithMessage("绑定参数失败", c)
		return
	}

	findAllRole, err := dal.Role.Where(dal.Role.StuID.Eq(req.StuId)).Find()
	if err != nil {
		global.Log.Warnln("查询任务失败: ", req.StuId)
		global.Log.Errorln(err)
		res.FailWithMessage("查询任务失败", c)
		return
	}

	for k := 0; k < len(findAllRole); { // for range 循环在每次循环时都会创建一个新的迭代变量，这个变量会被重新赋值，而不是更新已有变量的值。 在此处，我们需要更新已有变量的值，因此使用传统的 for 循环。
		v := findAllRole[k]
		if v.Status == 3 {
			// a = append(a[:i], a[i+1:]...) // 删除中间1个元素
			findAllRole = append(findAllRole[:k], findAllRole[k+1:]...)
			continue
		}
		k++
	}

	// todo 优化

	var data []returnRole
	if len(findAllRole) != 0 {
		for _, v := range findAllRole {
			var returnRoleOne returnRole
			returnRoleOne.StuID = v.StuID
			returnRoleOne.RoleID = v.RoleID
			returnRoleOne.RoleName = v.RoleName
			returnRoleOne.ClassID = v.ClassID
			returnRoleOne.Status = v.Status
			returnRoleOne.Deadline = v.Deadline
			returnRoleOne.CourseID = v.CourseID

			findFileName, err := dal.Course.Where(dal.Course.CourseID.Eq(v.CourseID)).First()
			if err != nil {
				return
			}
			returnRoleOne.FileType = findFileName.FileType

			data = append(data, returnRoleOne)
		}
	}

	if len(findAllRole) == 0 {
		global.Log.Infoln("没有任务: ", req.StuId)
		res.Fail(data, "没有任务", c)
	} else {
		global.Log.Infoln("查询成功: ", req.StuId)
		res.Ok(data, "查询成功", c)
	}
	// 不包含自己发布的任务
}

type returnRole struct {
	StuID    int32     `gorm:"column:stu_id;not null" json:"stu_id"`                        // 用户id
	RoleID   int32     `gorm:"column:role_id;primaryKey;autoIncrement:true" json:"role_id"` // 任务id
	RoleName string    `gorm:"column:role_name;not null" json:"role_name"`                  // 任务名(course_name)
	ClassID  int32     `gorm:"column:class_id;not null" json:"class_id"`                    // class_id
	Status   int32     `gorm:"column:status;not null" json:"status"`                        // 1:已提交 2:未提交 3:发布者(不显示)
	Deadline time.Time `gorm:"column:deadline;not null" json:"deadline"`                    // 截止时间
	CourseID int32     `gorm:"column:course_id;not null" json:"course_id"`
	FileType string    `gorm:"column:file_type;not null" json:"file_type"` // 文件类型
}
