package middleware

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
	"snail/teacher_backend/common"
	"snail/teacher_backend/models"
	"snail/teacher_backend/utils"
	"strconv"
)

func CourseOperationMiddleware(idKey string, readOnly bool) func(c *gin.Context) {
	return func(c *gin.Context) {
		courseID, err := getCourseID(c, idKey)
		if err != nil {
			log.Printf("Course operation middle ware get course id failed: %v\n", err)
			c.JSON(http.StatusOK, common.BadResponse(common.ParamError))
			c.Abort()
			return
		}
		if !isCoursePrivate(courseID) && readOnly {
			c.Next()
			return
		}
		org, _ := c.Get("user")
		user, err := utils.GetToken(org)
		if err != nil {
			log.Printf("Get token failed: %v\n", err)
			c.JSON(http.StatusOK, common.BadResponse(common.ServerError))
			c.Abort()
			return
		}
		userType := user.GetType()
		if userType == 1 {
			tmp := new(models.Course)
			tmp.ID = courseID
			if err = models.GetSingleCourse(tmp); err != nil {
				log.Printf("Course operation middle ware get single course failed: %v\n", err)
				c.JSON(http.StatusOK, common.BadResponse(common.ServerError))
				c.Abort()
				return
			}
			// 课程不存在或者无权限
			if tmp.SearchCode == "" || tmp.CreateBy != user.GetIdentity() {
				log.Printf("User: %v has no access right to the course %v or course no exist.\n", user.GetIdentity(), courseID)
				c.JSON(http.StatusOK, common.BadResponse(common.Error))
				c.Abort()
				return
			}
			c.Next()
		} else {
			tmp := new(models.Assistance)
			tmp.StuID = user.GetIdentity()
			assistanceList, err := models.GetAssistance(tmp)
			if err != nil {
				log.Printf("Course operation middle ware get assistance failed: %v\n", err)
				c.JSON(http.StatusOK, common.BadResponse(common.ServerError))
				c.Abort()
				return
			}
			for _, assistance := range assistanceList {
				if assistance.CourseID == courseID {
					c.Next()
					return
				}
			}
			log.Printf("User: %v has no access right to the course %v or course no exist.\n", user.GetIdentity(), courseID)
			c.JSON(http.StatusOK, common.BadResponse(common.Error))
			c.Abort()
			return
		}
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
