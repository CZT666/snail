package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"log"
	"net/http"
	"snail/teacher_backend/logic"
	"snail/teacher_backend/models"
	"snail/teacher_backend/vo"
)

func AddAssistance(c *gin.Context) {
	assistance := new(models.Assistance)
	if err := c.ShouldBindBodyWith(&assistance, binding.JSON); err != nil {
		log.Printf("Add assistance bind json failed: %v\n", err)
		c.JSON(http.StatusOK, vo.BadResponse(vo.ParamError))
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
		c.JSON(http.StatusOK, vo.BadResponse(vo.ParamError))
		return
	}
	baseResponse := logic.DeleteAssistance(assistance)
	c.JSON(http.StatusOK, baseResponse)
	return
}

func FindStudent(ctx *gin.Context) {
	req := new(models.Student)
	if err := ctx.ShouldBindBodyWith(&req, binding.JSON); err != nil {
		log.Printf("Add assistance bind json failed: %v\n", err)
		ctx.JSON(http.StatusOK, vo.BadResponse(vo.ParamError))
		return
	}
	baseResponse := logic.FindStudent(req)
	ctx.JSON(http.StatusOK, baseResponse)
	return
}

func GetAllAssistance(c *gin.Context) {
	req := new(models.Assistance)
	if err := c.ShouldBindBodyWith(&req, binding.JSON); err != nil {
		log.Printf("Add assistance bind json failed: %v\n", err)
		c.JSON(http.StatusOK, vo.BadResponse(vo.ParamError))
		return
	}
	baseResponse := logic.GetAllAssistance(req)
	c.JSON(http.StatusOK, baseResponse)
	return
}
