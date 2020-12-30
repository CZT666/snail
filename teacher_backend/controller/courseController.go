package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"log"
	"net/http"
	"snail/teacher_backend/logic"
	"snail/teacher_backend/models"
	"snail/teacher_backend/models/helper"
	"snail/teacher_backend/vo"
)

func AddCourse(c *gin.Context) {
	course := new(models.Course)
	if err := c.BindJSON(&course); err != nil {
		log.Printf("Add course bind json failed: %v\n", err)
		c.JSON(http.StatusOK, vo.BadResponse(vo.ParamError))
		return
	}
	org, _ := c.Get("user")
	user, err := models.GetToken(org)
	if err != nil {
		log.Printf("Course controller get token failed: %v\n", err)
		c.JSON(http.StatusOK, vo.BadResponse(vo.ServerError))
		return
	}
	var baseResponse *vo.BaseResponse
	baseResponse = logic.AddCourse(course, user)
	c.JSON(http.StatusOK, baseResponse)
	return
}

func QueryCourseList(c *gin.Context) {
	org, _ := c.Get("user")
	user, err := models.GetToken(org)
	if err != nil {
		log.Printf("Get token failed: %v\n", err)
		c.JSON(http.StatusOK, vo.BadResponse(vo.ServerError))
		return
	}
	pageRequest := helper.NewPageRequest()
	if err = c.BindJSON(&pageRequest); err != nil {
		log.Printf("Query course list bind json failed: %v\n", err)
		c.JSON(http.StatusOK, vo.BadResponse(vo.ParamError))
		return
	}

	baseResponse := new(vo.BaseResponse)
	baseResponse = logic.QueryCourseList(user, pageRequest)
	c.JSON(http.StatusOK, baseResponse)
	return
}

func QueryCourseDetail(c *gin.Context) {
	course := new(models.Course)
	if err := c.ShouldBindBodyWith(&course, binding.JSON); err != nil {
		log.Printf("Query course bind json failed: %v\n", err)
		c.JSON(http.StatusOK, vo.BadResponse(vo.ParamError))
		return
	}
	baseResponse := logic.QueryCourseDetail(course)
	c.JSON(http.StatusOK, baseResponse)
	return
}

func UpdateCourse(c *gin.Context) {
	course := new(models.Course)
	if err := c.ShouldBindBodyWith(&course, binding.JSON); err != nil {
		log.Printf("Update course bind json failed: %v\n", err)
		c.JSON(http.StatusOK, vo.BadResponse(vo.ParamError))
		return
	}
	baseResponse := logic.UpdateCourse(course)
	c.JSON(http.StatusOK, baseResponse)
	return
}

func DeleteCourse(c *gin.Context) {
	course := new(models.Course)
	if err := c.ShouldBindBodyWith(&course, binding.JSON); err != nil {
		log.Printf("Delete course bind json failed: %v\n", err)
		c.JSON(http.StatusOK, vo.BadResponse(vo.ParamError))
		return
	}
	baseResponse := logic.DeleteCourse(course)
	c.JSON(http.StatusOK, baseResponse)
	return
}
