package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"log"
	"net/http"
	"snail/teacher_backend/common"
	"snail/teacher_backend/logic"
	"snail/teacher_backend/models"
)

func AddAssistance(c *gin.Context) {
	assistance := new(models.Assistance)
	if err := c.ShouldBindBodyWith(&assistance, binding.JSON); err != nil {
		log.Printf("Add assistance bind json failed: %v\n", err)
		c.JSON(http.StatusOK, common.BadResponse(common.ParamError))
		return
	}
	baseResponse := logic.AddAssistance(assistance)
	c.JSON(http.StatusOK, baseResponse)
	return
}

func DeleteAssistance(c *gin.Context) {
	assistance := new(models.Assistance)
	if err := c.ShouldBindBodyWith(&assistance, binding.JSON); err != nil {
		log.Printf("Delete assistance bind json failed: %v\n", err)
		c.JSON(http.StatusOK, common.BadResponse(common.ParamError))
		return
	}
	baseResponse := logic.DeleteAssistance(assistance)
	c.JSON(http.StatusOK, baseResponse)
	return
}
