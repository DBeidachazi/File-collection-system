package classuser_api

import (
	"FengfengStudy/global"
	"FengfengStudy/models/orm/dal"
	"FengfengStudy/models/res"
	"github.com/gin-gonic/gin"
)

type ModifyClassUser struct {
	ClassId     int32 `json:"classId" bind:"required"`
	StuId       int32 `json:"stuId" bind:"required"`       // 提出修改的用户
	ModifyStuId int32 `json:"modifyStuId" bind:"required"` // 修改的用户
	Permission  int32 `json:"permission" bind:"required"`  // 修改的权限
}

func (ClassUserApi) ModifyClassUserView(c *gin.Context) {
	var req ModifyClassUser
	err := c.ShouldBindJSON(&req)
	if err != nil {
		return
	}
	if req.Permission < 0 || req.Permission >= 2 {
		global.Log.Warnln("权限不合法", req.Permission)
		res.FailWithMessage("权限不合法", c)
		return
	}
	// todo 重复
	classStu, err := dal.ClassUser.Where(dal.ClassUser.ClassID.Eq(req.ClassId), dal.ClassUser.StuID.Eq(req.StuId)).First()
	if err != nil {
		global.Log.Warnln("管理员权限不足", req.ClassId, req.StuId)
		res.FailWithMessage("管理员权限不足", c)
		return
	}

	if classStu.PermissionType != 2 {
		global.Log.Warnln("管理员权限不足", req.ClassId, req.StuId)
		res.FailWithMessage("管理员权限不足", c)
		return
	}
	if req.ModifyStuId == req.StuId {
		global.Log.Warnln("不能修改自己的权限", req.ModifyStuId)
		res.FailWithMessage("不能修改自己的权限", c)
		return
	}
	// 检测班级和用户是否存在
	_, err = dal.Class.Where(dal.Class.ClassID.Eq(req.ClassId)).First()
	if err != nil {
		global.Log.Warnln("班级不存在", req.ClassId)
		res.FailWithMessage("班级不存在", c)
		return
	}
	_, err = dal.User.Where(dal.User.StuID.Eq(req.StuId)).First()
	if err != nil {
		global.Log.Warnln("管理员用户不存在", req.StuId)
		res.FailWithMessage("管理员用户不存在", c)
		return
	}
	_, err = dal.User.Where(dal.User.StuID.Eq(req.ModifyStuId)).First()
	if err != nil {
		global.Log.Warnln("被修改的用户不存在", req.ModifyStuId)
		res.FailWithMessage("被修改的用户不存在", c)
		return
	}

	// 检查被修改的用户是否在传入的班级中
	_, err = dal.ClassUser.Where(dal.ClassUser.StuID.Eq(req.ModifyStuId), dal.ClassUser.ClassID.Eq(req.ClassId)).First()
	if err != nil {
		global.Log.Warnln("被修改的用户不在传入的班级中", req.ClassId, req.ModifyStuId)
		res.FailWithMessage("被修改的用户不在传入的班级中", c)
		return
	}
	// 检查修改权限的用户是否在传入的班级中
	_, err = dal.ClassUser.Where(dal.ClassUser.StuID.Eq(req.StuId), dal.ClassUser.ClassID.Eq(req.ClassId)).First()
	if err != nil {
		global.Log.Warnln("修改权限的用户不在传入的班级中", req.ClassId, req.StuId)
		res.FailWithMessage("修改权限的用户不在传入的班级中", c)
		return
	}

	// 修改用户权限
	u := dal.ClassUser
	_, err = u.Where(u.ClassID.Eq(req.ClassId), u.StuID.Eq(req.ModifyStuId)).Update(u.PermissionType, req.Permission)
	if err != nil {
		return
	} else {
		global.Log.Infoln("修改用户权限成功", req.ClassId, req.StuId, req.ModifyStuId)
		res.OkWithMessage("修改用户权限成功", c)
	}
}
