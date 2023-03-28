package classuser_api

import (
	"FengfengStudy/global"
	"FengfengStudy/models/orm/dal"
	"FengfengStudy/models/res"
	"github.com/gin-gonic/gin"
)

type SearchClassUserRequest struct {
	ClassId int32 `json:"class_id" bind:"required"`
}

func (ClassUserApi) SearchCLassUserView(c *gin.Context) {
	var req SearchClassUserRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		global.Log.Warnln("绑定参数失败", err)
		res.FailWithMessage("绑定参数失败", c)
		return
	}
	findAllStudents, err := dal.ClassUser.Where(dal.ClassUser.ClassID.Eq(req.ClassId)).Find()
	if err != nil {
		global.Log.Warnln("查询班级学生失败", req.ClassId)
		res.FailWithMessage("查询班级学生失败", c)
		return
	}
	// 删除班级创建者
	for i, v := range findAllStudents {
		if v.PermissionType == 2 {
			findAllStudents = append(findAllStudents[:i], findAllStudents[i+1:]...)
			break
		}
	}
	res.Ok(findAllStudents, "查询班级学生成功", c)
}
