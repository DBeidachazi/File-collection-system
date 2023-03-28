package class_api

import (
	"FengfengStudy/global"
	"FengfengStudy/models/orm/dal"
	"FengfengStudy/models/orm/model"
	"FengfengStudy/models/res"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type ClassSearchRequest struct {
	StuId int32 `json:"stu_id" bind:"required"`
}
type ClassResponse struct {
	ClassUser *model.ClassUser `json:"class_user"`
	ClassName string           `json:"class_name"`
}

// 仅查询我创建的班级
func (ClassApi) SearchClassView(c *gin.Context) {
	var req ClassSearchRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		return
	}
	// 找到所有创建的班级
	findClassUser, err := dal.ClassUser.Where(dal.ClassUser.StuID.Eq(req.StuId), dal.ClassUser.PermissionType.Eq(2)).Find()
	if err != nil || len(findClassUser) == 0 {
		global.Log.Warnln("查询班级失败: ", req.StuId)
		logrus.Warnln("err: ", err)
		res.FailWithMessage("没有创建班级", c)
		return
	}
	var classResponse []ClassResponse
	for _, v := range findClassUser {
		logrus.Infoln(v)
		findClass, err := dal.Class.Where(dal.Class.ClassID.Eq(v.ClassID)).First()
		if err != nil {
			global.Log.Warnln("查询班级失败: ", v.ClassID)
			logrus.Warnln("err: ", err)
			res.FailWithMessage("查询班级失败", c)
			return
		}
		classResponse = append(classResponse, ClassResponse{
			ClassUser: v,
			ClassName: findClass.ClassName,
		})
	}
	res.Ok(classResponse, "查询班级成功", c)
	/* todo 班级查询
	1. 查询班级用户表，获取全部班级id
	2. 管理权限 0:普通用户 1:管理员 2: 班级创建者 ,
	    根据管理权限将页面分为我管理的和我加入的,还有一个创建班级
	3. 我管理的班级中
		1. 如果我是班级创建者，我可以设置管理员
		2. 如果我是班级创建者或管理员，我可以发布作业

	*/
	// 班级管理
	// 我管理的班级 创建班级
	// 创建班级（班级管理员设定）

}
