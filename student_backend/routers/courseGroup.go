package routers

import (
	"github.com/gin-gonic/gin"
	"student_bakcend/controller"
	"student_bakcend/middleware"
)

func courseGroup(engine *gin.Engine) {
	group := engine.Group("/course")
	group.Use(middleware.JWTAuthMiddleware())
	{
		group.POST("/join",controller.JoinCourse)
		group.GET("/list", controller.QueryCourseList)
		group.GET("/detail", controller.QueryCourseDetail)
		group.GET("/search/:name",controller.SearchCourse)
		group.GET("/getStudentCourse",controller.GetStudentCourse)
	}
}
