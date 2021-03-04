package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"log"
	"net/http"
	"student_bakcend/logic"
	"student_bakcend/models"
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
	//student := new(models.Student)
	courseToStudent := new(models.CourseToStudent)
	if err := c.BindJSON(&courseToStudent); err != nil {
		log.Printf("course to student bind json failed: %v\n", err)
		c.JSON(http.StatusOK, vo.BadResponse(vo.ParamError))
		return
	}
	courseToStudent.StudentID = student.StudentID
	baseResponse := new(vo.BaseResponse)
	baseResponse = logic.JoinCourse(courseToStudent)
	c.JSON(http.StatusOK, baseResponse)
	return
}

func QueryCourseList(c *gin.Context) {
	baseResponse := logic.QueryCourseList()
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
	baseResponse := logic.SearchCourse(searchName)
	c.JSON(http.StatusOK, baseResponse)
	return
}

func GetStudentCourse(c *gin.Context) {
	org, _ := c.Get("user")
	student, err := utils.GetToken(org)
	if err != nil {
		log.Printf("Get token failed: %v\n", err)
		c.JSON(http.StatusOK, vo.BadResponse(vo.ServerError))
		return
	}
	//student := models.Student{
	//	ID: 16,
	//	StudentID: "20171003389",
	//	}
	baseResponse := logic.GetStudentCourse(student.StudentID)
	c.JSON(http.StatusOK, baseResponse)
	return
}

