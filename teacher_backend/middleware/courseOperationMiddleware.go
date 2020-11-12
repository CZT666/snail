package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"log"
	"net/http"
	"snail/teacher_backend/common"
	"snail/teacher_backend/models"
	"snail/teacher_backend/utils"
)

func CourseOperationMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		course := new(models.Course)
		if err := c.ShouldBindBodyWith(&course, binding.JSON); err != nil || course.ID < 1 {
			log.Printf("Course operation middle ware bind json failed: %v\n", err)
			c.JSON(http.StatusOK, common.BadResponse(common.ParamError))
			c.Abort()
			return
		}
		org, _ := c.Get("user")
		// TODO 助教
		user, err := utils.GetToken(org)
		if err != nil {
			log.Printf("Get token failed: %v\n", err)
			c.JSON(http.StatusOK, common.BadResponse(common.ServerError))
			return
		}
		userType := user.GetType()
		if userType == 1 {
			tmp := new(models.Course)
			tmp.ID = course.ID
			if err = models.GetSingleCourse(tmp); err != nil {
				log.Printf("Course operation middle ware get single course failed: %v\n", err)
				c.JSON(http.StatusOK, common.BadResponse(common.ServerError))
				c.Abort()
				return
			}
			// 课程不存在或者无权限
			if tmp.SearchCode == "" || tmp.CreateBy != user.GetIdentity() {
				log.Printf("User has no access right to the course or course no exist: %v\n", course.ID)
				c.JSON(http.StatusOK, common.BadResponse(common.Error))
				c.Abort()
				return
			}
		}
		c.Next()
	}
}