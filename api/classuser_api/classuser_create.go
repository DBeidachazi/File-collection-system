package classuser_api

import (
	"FengfengStudy/models/orm/dal"
	"FengfengStudy/models/orm/model"
	"FengfengStudy/models/res"
	"github.com/gin-gonic/gin"
)

type CreateClassUserRequest struct {
	classId int32 `json:"classId" bind:"required"`
	stuId   int32 `json:"userId" bind:"required"`
	//permission int32 `json:"permission" bind:"required"`
}

func (ClassUserApi) CreateClassUserView(c *gin.Context) {
	var req CreateClassUserRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		return
	}
	_, err = dal.Class.Where(dal.Class.ClassID.Eq(req.classId)).First()
	if err != nil {
		return
	}
	stu, err := dal.User.Where(dal.User.StuID.Eq(req.stuId)).First()
	if err != nil {
		return
	}

	// 检查用户是否已经加入班级
	_, err = dal.ClassUser.Where(dal.ClassUser.ClassID.Eq(req.classId), dal.ClassUser.StuID.Eq(req.stuId)).First()
	if err != nil {
		// 用户未加入班级
		UserJoinClass(c, req.classId, req.stuId, 0, stu.Username) // 默认权限为0
		res.OkWithMessage("用户加入班级成功", c)
		return
	}
	// 用户已加入班级
	res.FailWithMessage("用户已加入班级", c)
	return
}

// 提取函数 for class_create and classuser_create
func UserJoinClass(c *gin.Context, classId int32, stuId int32, permission int32, name string) {
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
