package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"log"
	"net/http"
	"student_bakcend/logic"
	"student_bakcend/models"
	"student_bakcend/models/helper"
	"student_bakcend/utils"
	"student_bakcend/vo"
)

func JoinCourse(c *gin.Context) {
	org, _ := c.Get("user")
	student, err := utils.GetToken(org)
	if err != nil {
		log.Printf("Get token failed: %v\n", err)
		c.JSON(http.StatusOK, vo.BadResponse(vo.ServerError))
		return
	}
	courseToStudent := new(models.CourseToStudent)
	if err := c.BindJSON(&courseToStudent); err != nil {
		log.Printf("course to student bind json failed: %v\n", err)
		c.JSON(http.StatusOK, vo.BadResponse(vo.ParamError))
		return
	}
	courseToStudent.StudentID = student.ID
	baseResponse := new(vo.BaseResponse)
	baseResponse = logic.JoinCourse(courseToStudent)
	c.JSON(http.StatusOK, baseResponse)
	return
}

func QueryCourseList(c *gin.Context) {
	//org, _ := c.Get("student")
	//student, err := utils.GetToken(org)
	//if err != nil {
	//	log.Printf("Get token failed: %v\n", err)
	//	c.JSON(http.StatusOK, vo.BadResponse(vo.ServerError))
	//	return
	//}
	pageRequest := helper.NewPageRequest()
	if err := c.BindJSON(&pageRequest); err != nil {
		log.Printf("Query course list bind json failed: %v\n", err)
		c.JSON(http.StatusOK, vo.BadResponse(vo.ParamError))
		return
	}

	baseResponse := new(vo.BaseResponse)
	baseResponse = logic.QueryCourseList(pageRequest)
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

func SearchCourse(c *gin.Context)  {
	searchName := c.Param("name")
	pageRequest := helper.NewPageRequest()
	if err := c.BindJSON(&pageRequest); err != nil {
		log.Printf("Query course list bind json failed: %v\n", err)
		c.JSON(http.StatusOK, vo.BadResponse(vo.ParamError))
		return
	}
	baseResponse := logic.SearchCourse(pageRequest,searchName)
	c.JSON(http.StatusOK, baseResponse)
	return
}

