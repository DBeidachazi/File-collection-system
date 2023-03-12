package class_api

import (
	"FengfengStudy/api/classuser_api"
	"FengfengStudy/global"
	"FengfengStudy/models/orm/dal"
	"FengfengStudy/models/orm/model"
	"FengfengStudy/models/res"
	"github.com/gin-gonic/gin"
)

type ClassCreateRequest struct {
	ClassId   int32  `json:"classId" bind:"required"`
	StuId     int32  `json:"stuId" bind:"required"`
	ClassName string `json:"className" bind:"required"`
}

func (ClassApi) CreateClassView(c *gin.Context) {
	var req ClassCreateRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		return
	}

	// 检查用户是否有权限创建班级
	_, err = dal.Class.Where(dal.Class.ClassID.Eq(req.ClassId)).First()
	if err == nil {
		// 班级已存在
		global.Log.Warnln("班级已存在", req.ClassId)
		res.FailWithMessage("班级已存在", c)
		return
	}
	//  查询用户信息
	name, err := dal.User.Where(dal.User.StuID.Eq(req.StuId)).First()
	if err != nil {
		global.Log.Warnln("创建班级的用户不存在", req.StuId)
		res.FailWithMessage("用户不存在", c)
		return
	}
	// 入库
	class := model.Class{
		ClassID:   req.ClassId,
		ClassName: req.ClassName, // 班级名称
	}
	err = dal.Class.Create(&class)
	if err != nil {
		global.Log.Warnln("入库失败", req.ClassId, req.ClassName)
		res.FailWithMessage("创建班级失败", c)
		return
	} else {
		// 传入班级id，用户id，权限，用户姓名
		classuser_api.UserJoinClass(req.ClassId, req.StuId, 2, name.Username)
		global.Log.Infoln(req.ClassId, "创建班级成功")
		res.OkWithMessage("创建班级成功", c)
	}

}
