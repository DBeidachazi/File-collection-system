package searchRole_api

import (
	"FengfengStudy/global"
	"FengfengStudy/models/orm/dal"
	"FengfengStudy/models/res"
	"github.com/gin-gonic/gin"
)

type SearchRoleRequest struct {
	StuId int32 `json:"stu_id" binding:"required"`
}

func (SearchRoleApi) SearchRoleView(c *gin.Context) {
	var req SearchRoleRequest
	err := c.ShouldBindJSON(&req)
	global.Log.Infoln(req.StuId)
	if err != nil {
		global.Log.Warnln("绑定参数失败")
		global.Log.Warnln(err)
		res.FailWithMessage("绑定参数失败", c)
		return
	}

	findAllRole, err := dal.Role.Where(dal.Role.StuID.Eq(req.StuId)).Find()
	if err != nil {
		global.Log.Warnln("查询任务失败: ", req.StuId)
		global.Log.Errorln(err)
		res.FailWithMessage("查询任务失败", c)
		return
	}

	for k := 0; k < len(findAllRole); { // for range 循环在每次循环时都会创建一个新的迭代变量，这个变量会被重新赋值，而不是更新已有变量的值。 在此处，我们需要更新已有变量的值，因此使用传统的 for 循环。
		v := findAllRole[k]
		if v.Status == 3 {
			// a = append(a[:i], a[i+1:]...) // 删除中间1个元素
			findAllRole = append(findAllRole[:k], findAllRole[k+1:]...)
			break
		}
		k++
	}

	if len(findAllRole) == 0 {
		res.FailWithMessage("没有任务", c)
	} else {
		res.Ok(findAllRole, "查询成功", c)
	}
	// 不包含自己发布的任务
}
