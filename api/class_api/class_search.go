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

}
