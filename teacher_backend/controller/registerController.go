package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"snail/teacher_backend/common"
	"snail/teacher_backend/logic"
	"snail/teacher_backend/models"
)

func RegisterTeacher(c *gin.Context) {
	teacher := new(models.Teacher)
	if err := c.BindJSON(&teacher); err != nil {
		baseResponse := new(common.BaseResponse)
		baseResponse.Code = common.ParamError
		fmt.Printf("Register teacher error: %v\n", err)
		c.JSON(http.StatusOK, baseResponse)
		return
	}
	baseResponse := logic.AddTeacher(teacher)
	c.JSON(http.StatusOK, baseResponse)
	return
}

func UpdateTeacherInfo(c *gin.Context) (baseResponse *common.BaseResponse) {
	teacher := new(models.Teacher)
	if err := c.ShouldBind(&teacher); err != nil {
		baseResponse = new(common.BaseResponse)
		baseResponse.Code = common.ParamError
		fmt.Printf("Update teacher error: %v\n", err)
		return
	}
	baseResponse = logic.UpdateTeacher(teacher)
	return
}
