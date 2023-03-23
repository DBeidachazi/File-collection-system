package class_api

import (
	"FengfengStudy/global"
	"FengfengStudy/models/orm/dal"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type ClassSearchRequest struct {
	StuId int32 `json:"stuId bind:required"`
}

func (ClassApi) SearchClassView(c *gin.Context) {
	var req ClassSearchRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		return
	}

	findClassUser, err := dal.ClassUser.Where(dal.ClassUser.StuID.Eq(req.StuId)).Find()
	if err != nil {
		global.Log.Warnln("查询班级失败: ", req.StuId)
		logrus.Warnln("err: ", err)
		return
	}
	_ = findClassUser
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
