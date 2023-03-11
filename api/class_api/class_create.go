package class_api

import "github.com/gin-gonic/gin"

type ClassCreateRequest struct {
	name    string `json:"name" bind:"required"`
	classId int32  `json:"classId" bind:"required"`
}

func (ClassApi) CreateClassView(c *gin.Context) {
	var classCreateRequest ClassCreateRequest
	err := c.ShouldBindJSON(&classCreateRequest)
	if err != nil {
		return
	}

}

// TODO 提取个函数让创建班级的用户使用
