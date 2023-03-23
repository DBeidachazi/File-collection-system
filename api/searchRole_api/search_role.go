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
		global.Log.Errorln(err)
		return
	}

	findAll, err := dal.Role.Where(dal.Role.StuID.Eq(req.StuId)).Find()
	if err != nil {
		global.Log.Errorln(err)
		return
	}

	for k, v := range findAll {
		global.Log.Infoln(v)
		if v.Status == 3 {
			// fixme a = append(a[:i], a[i+1:]...) // 删除中间1个元素
			findAll = append(findAll[:k], findAll[k+1:]...)
			k--
		}
	}

	if len(findAll) == 0 {
		res.FailWithMessage("没有任务", c)
	} else {
		res.Ok(findAll, "查询成功", c)
	}
	// 不包含自己发布的任务
}
