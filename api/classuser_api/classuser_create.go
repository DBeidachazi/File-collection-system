package classuser_api

import (
	"FengfengStudy/global"
	"FengfengStudy/models/orm/dal"
	"FengfengStudy/models/orm/model"
	"FengfengStudy/models/res"
	"github.com/gin-gonic/gin"
)

type CreateClassUserRequest struct {
	ClassId int32 `json:"classId" bind:"required"`
	StuId   int32 `json:"stuId" bind:"required"`
}

func (ClassUserApi) CreateClassUserView(c *gin.Context) {
	var req CreateClassUserRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		return
	}
	// 检测班级和用户是否存在
	_, err = dal.Class.Where(dal.Class.ClassID.Eq(req.ClassId)).First()
	if err != nil {
		global.Log.Warnln("班级不存在", req.ClassId)
		res.FailWithMessage("班级不存在", c)
		return
	}
	stu, err := dal.User.Where(dal.User.StuID.Eq(req.StuId)).First()
	if err != nil {
		global.Log.Warnln("用户不存在", req.StuId)
		res.FailWithMessage("用户不存在", c)
		return
	}

	// 检查用户是否已经加入班级
	_, err = dal.ClassUser.Where(dal.ClassUser.ClassID.Eq(req.ClassId), dal.ClassUser.StuID.Eq(req.StuId)).First()
	if err != nil {
		// 用户未加入班级
		UserJoinClass(req.ClassId, req.StuId, 0, stu.Username) // 默认权限为0
		res.OkWithMessage("用户加入班级成功", c)
		return
	} else {
		// 用户已加入班级
		global.Log.Warnln("用户已加入班级", req.ClassId, req.StuId)
		res.FailWithMessage("用户已加入班级", c)
		return
	}
}

// UserJoinClass 提取函数 for class_create and classuser_create
func UserJoinClass(classId int32, stuId int32, permission int32, name string) {
	err := dal.ClassUser.Create(&model.ClassUser{
		ClassID:        classId,
		StuID:          stuId,
		PermissionType: permission,
		Username:       name,
	})
	if err != nil {
		return
	}
}
