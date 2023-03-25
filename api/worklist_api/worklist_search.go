package worklist_api

import (
	"FengfengStudy/global"
	"FengfengStudy/models/orm/dal"
	"FengfengStudy/models/res"
	"github.com/gin-gonic/gin"
)

type WorklistRequest struct {
	StuId int32 `json:"stu_id" binding:"required"`
}

func (WorklistApi) WorklistSearchView(c *gin.Context) {
	var req WorklistRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		global.Log.Warnln("绑定参数失败")
		global.Log.Warnln(err)
		res.FailWithMessage("绑定参数失败", c)
		return
	}
	_, err = dal.Worklist.Where(dal.Worklist.StuID.Eq(req.StuId)).First()
	if err != nil {
		global.Log.Warnln("没有worklist: ", req.StuId)
		res.FailWithMessage("没有发布任务", c)
		return
	}
	findAllWorklist, err := dal.Worklist.Where(dal.Worklist.StuID.Eq(req.StuId)).Find()
	if err != nil {
		global.Log.Warnln("查询工作列表失败: ", req.StuId)
		global.Log.Warnln(err)
		res.FailWithMessage("没有发布任务", c)
		return
	}

	global.Log.Infoln("查询工作列表成功: ", req.StuId)
	res.Ok(findAllWorklist, "查询工作列表成功", c)
}
