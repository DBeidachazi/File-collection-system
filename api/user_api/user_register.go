package user_api

import (
	"FengfengStudy/global"
	"FengfengStudy/models/orm/dal"
	"FengfengStudy/models/orm/model"
	"FengfengStudy/models/res"
	"FengfengStudy/utils/pwd"
	"github.com/gin-gonic/gin"
)

type UserRegisterRequest struct {
	StuId    int32  `json:"stu_id" binding:"required"`
	Password string `json:"password" binding:"required"`
	Name     string `json:"name" binding:"required"`
}

func (UserApi) UserRegisterView(c *gin.Context) {
	// 参数绑定
	var req UserRegisterRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		return
	}

	_, err = dal.User.Where(dal.User.StuID.Eq(req.StuId)).First()
	if err == nil {
		// 用户已存在
		global.Log.Warnln("用户已存在", req.StuId)
		res.FailWithMessage("用户已注册", c)
		return
	}

	// 入库
	user := model.User{
		StuID:    req.StuId,
		Name:     req.Name,
		Password: pwd.HashPwd(req.Password),
	}
	err = dal.User.Create(&user)
	if err != nil {
		global.Log.Warnln("入库失败", req.StuId, req.Password)
		res.FailWithMessage("注册失败", c)
		return
	} else {
		global.Log.Infoln(req.StuId, "注册成功")
		res.OkWithMessage("注册成功", c)
	}

}
