package user_api

import (
	"FengfengStudy/global"
	"FengfengStudy/models/orm/dal"
	"FengfengStudy/models/res"
	"FengfengStudy/utils/pwd"
	"github.com/gin-gonic/gin"
)

type UserLoginRequest struct {
	StuId    int32  `json:"stu_id" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (UserApi) UserLoginView(c *gin.Context) {
	// 参数绑定
	var req UserLoginRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		return
	}

	searchId, err := dal.User.Where(dal.User.StuID.Eq(req.StuId)).First()
	if err != nil {
		// 用户不存在
		global.Log.Warnln("用户不存在", req.StuId)
		res.FailWithMessage("用户名或密码错误", c)
	}
	// 校验密码
	isCheck := pwd.CheckPwd(searchId.Password, req.Password)
	if !isCheck {
		// 密码错误
		global.Log.Warnln("密码错误", req.StuId, req.Password)
		res.FailWithMessage("用户名或密码错误", c)
		return
	}
	// 登陆成功
	global.Log.Infoln(req.StuId, "登陆成功")
	res.OkWithMessage("登录成功", c)
}
