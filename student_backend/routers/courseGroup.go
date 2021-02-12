package routers

import (
	"github.com/gin-gonic/gin"
	"student_bakcend/controller"
	"student_bakcend/middleware"
)

func courseGroup(engine *gin.Engine) {
	courseID := "id"
	group := engine.Group("/course")
	//group.Use(middleware.JWTAuthMiddleware())
	{
		group.POST("/join",controller.JoinCourse)
		group.GET("/list", controller.QueryCourseList)
		group.GET("/detail", middleware.CourseOperationMiddleware(courseID, true), controller.QueryCourseDetail)
		group.GET("/search/:name",controller.SearchCourse)
	}
}
