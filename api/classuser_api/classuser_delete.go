package classuser_api

import (
	"FengfengStudy/global"
	"FengfengStudy/models/orm/dal"
	"FengfengStudy/models/res"
	"github.com/gin-gonic/gin"
)

type ClassUserDeleteRequest struct {
	ClassId int32 `json:"class_id" bind:"required"`
	StuId   int32 `json:"stu_id" bind:"required"`
}

func (ClassUserApi) ClassUserDeleteView(c *gin.Context) {
	var req ClassUserDeleteRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		global.Log.Warnln("绑定参数失败", err)
		res.FailWithMessage("绑定参数失败", c)
		return
	}

	_, err = dal.ClassUser.Where(dal.ClassUser.ClassID.Eq(req.ClassId), dal.ClassUser.StuID.Eq(req.StuId)).Delete()
	if err != nil {
		global.Log.Warnln("删除班级学生失败", req.ClassId, req.StuId)
		res.FailWithMessage("删除班级学生失败", c)
	}

	// 删除role
	_, err = dal.Role.Where(dal.Role.ClassID.Eq(req.ClassId), dal.Role.StuID.Eq(req.StuId)).Delete()
	if err != nil {
		global.Log.Warnln("删除role失败", req.ClassId, req.StuId)
		res.FailWithMessage("删除role失败", c)
	}

	res.OkWithMessage("删除班级学生成功", c)

}
