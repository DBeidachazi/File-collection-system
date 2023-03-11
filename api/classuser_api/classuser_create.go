package classuser_api

import "github.com/gin-gonic/gin"

type CreateClassUserRequest struct {
	classId    int32 `json:"classId" bind:"required"`
	stuId      int32 `json:"userId" bind:"required"`
	permission int32 `json:"permission" bind:"required"`
}

func (ClassUserApi) CreateClassUserView(c *gin.Context) {
	var createClassUserRequest CreateClassUserRequest
	err := c.ShouldBindJSON(&createClassUserRequest)
	if err != nil {
		return
	}

}
