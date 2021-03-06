package middleware

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"snail/student_bakcend/models"
	"snail/student_bakcend/utils"
	"snail/student_bakcend/vo"
)

func CourseOperationMiddleware(idKey string, readOnly bool) func(c *gin.Context) {
	return func(c *gin.Context) {
		courseID, err := getCourseID(c, idKey)
		fmt.Printf("coureseID msg : %v\n",courseID)
		if err != nil {
			log.Printf("Course operation middle ware get course id failed: %v\n", err)
			c.JSON(http.StatusOK, vo.BadResponse(vo.ParamError))
			c.Abort()
			return
		}
		if !isCoursePrivate(courseID) && readOnly {
			c.Next()
			return
		}else{
			org, _ := c.Get("student")
			student, err := models.GetToken(org)
			if err != nil {
				log.Printf("Get token failed: %v\n", err)
				c.JSON(http.StatusOK, vo.BadResponse(vo.ServerError))
				c.Abort()
				return
			}
			tmp := new(models.CourseToStudent)
			tmp.CourseID = courseID
			tmp.StudentID = student.GetIdentity()
			if err = models.MatchCourseStudent(tmp); err != nil {
				log.Printf("student not allow access course: %v\n", err)
				c.JSON(http.StatusOK, vo.BadResponse(vo.Error))

				c.Abort()
				return
			}
			if tmp.IsValid == 0{
				baseResponse := new(vo.BaseResponse)
				baseResponse.Code = vo.Error
				baseResponse.Msg = "student not allow access"
				c.JSON(http.StatusOK, baseResponse)
				c.Abort()
			}else {
				c.Next()
			}
		}
		return
	}
}

func getCourseID(c *gin.Context, idKey string) (courseID int, err error) {
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		log.Printf("Course operation middle ware get body failed: %v\n", err)
		return
	}
	id, err := utils.GetJsonValue(body, idKey)
	if err != nil {
		log.Printf("Course operation middle ware parse json failed: %v\n", err)
		return
	}
	courseID, err = strconv.Atoi(id)
	if err != nil {
		log.Printf("Course operation middle ware convert string to int failed: %v\n", err)
	}
	c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(body))
	return
}

func isCoursePrivate(courseID int) bool {
	course := new(models.Course)
	course.ID = courseID
	if err := models.GetSingleCourse(course); err != nil {
		log.Printf("Course operation middle ware get single course failed: %v\n", err)
		return false
	}
	return course.IsPrivate == 1
}
