package setting_api

import (
	"FengfengStudy/models/orm/dal"
	"FengfengStudy/models/res"
	"github.com/gin-gonic/gin"
)

func (SettingsApi) SettingsInfoView(c *gin.Context) {
	// todo 样例
	//res.Ok(map[string]string{}, "成功", c)
	//res.OkWithData(map[string]string{
	//	"test":  "test",
	//	"test2": "test2",
	//}, c)
	stuQId, err := dal.ClassUser.Where(dal.ClassUser.StuID.Eq(20212319)).First()
	if err != nil {
		return
	}
	res.OkWithData(stuQId, c)
	//res.FailWithCode(res.SystemError, c)

}
