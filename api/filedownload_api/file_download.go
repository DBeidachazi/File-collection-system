package filedownload_api

import (
	"FengfengStudy/global"
	"FengfengStudy/models/orm/dal"
	"FengfengStudy/models/res"
	"FengfengStudy/utils"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"os"
	"strconv"
)

type FileDownloadRequest struct {
	CourseId int32 `json:"courseId" bind:"required"`
	StuId    int32 `json:"stuId" bind:"required"`
	ClassId  int32 `json:"classId" bind:"required"`
}

func (FileDownloadApi) FileDownloadApi(c *gin.Context) {
	var req FileDownloadRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		global.Log.Warnln("参数错误", err)
		return
	}
	findPermission, err := dal.ClassUser.Where(dal.ClassUser.ClassID.Eq(req.ClassId), dal.ClassUser.StuID.Eq(req.StuId)).First()
	if err != nil {
		// 在该班级中没有找到该用户
		global.Log.Warnln("在班级中没有找到该用户")
		res.FailWithMessage("没有权限下载", c)
		return
	}
	if findPermission.PermissionType == 0 {
		global.Log.Warnln("权限不足")
		res.FailWithMessage("权限不足", c)
		return
	}

	// 压缩后文件删除
	dir := "./files/" + trans(req.ClassId) + "/" + trans(req.CourseId) + "/"
	dst := "./files/" + trans(req.ClassId) + "/" + trans(req.CourseId) + "/" + trans(req.StuId) + ".zip"
	if err := utils.Zip(dir, dst); err != nil {
		global.Log.Warnln("压缩文件失败", err)
		res.FailWithMessage("压缩文件失败", c)
		return
	}
	global.Log.Infof("压缩文件成功")
	// 传输文件
	c.File(dst)
	// 删除压缩文件
	err = os.Remove(dst)
	logrus.Infof("删除压缩文件成功")
	if err != nil {
		global.Log.Warnln("删除压缩文件失败", err)
		return
	}
	return
}

func trans(i int32) string {
	return strconv.Itoa(int(i))
}
