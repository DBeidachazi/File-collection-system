package fileupload_api

import (
	"FengfengStudy/global"
	"FengfengStudy/models/orm/dal"
	"FengfengStudy/models/res"
	"github.com/gin-gonic/gin"
	"strconv"
)

// todo role.status 修改状态
func (FileUploadApi) FileUploadView(c *gin.Context) {
	classId := c.Param("classid")
	stuId := c.Param("stuid")
	courseId := c.Param("courseid")

	searchFileType, err := dal.Course.Where(dal.Course.ClassID.Eq(trans(classId)), dal.Course.CourseID.Eq(trans(courseId))).First()
	if err != nil {
		global.Log.Warnln("没有找到该班级下的课程")
		res.FailWithMessage("没有找到该班级下的课程", c)
		return
	}
	_, err = dal.ClassUser.Where(dal.ClassUser.ClassID.Eq(trans(classId)), dal.ClassUser.StuID.Eq(trans(stuId))).First()
	if err != nil {
		global.Log.Warnln("没有找到该班级中的学生")
		res.FailWithMessage("没有找到该班级中的学生", c)
		return
	}
	fileType := searchFileType.FileType
	file, _ := c.FormFile("file") // 上传请求参数名？ form-key
	fileName := classId + "_" + stuId
	dst := "./files/" + classId + "/" + courseId + "/" + fileName + "." + fileType
	err = c.SaveUploadedFile(file, dst)
	if err != nil {
		global.Log.Warnln("文件保存失败", err)
		return
	}

	global.Log.Infoln("文件上传成功", classId, courseId)
	res.OkWithMessage("文件上传成功", c)
	return
}

func trans(s string) int32 {
	atoi, err := strconv.Atoi(s)
	if err != nil {
		global.Log.Warnln("Atoi失败", err)
		return 0
	}
	return int32(atoi)
}
