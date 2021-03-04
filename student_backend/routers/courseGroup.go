package routers

import (
	"github.com/gin-gonic/gin"
	"student_bakcend/controller"
)

func courseGroup(engine *gin.Engine) {
	group := engine.Group("/course")
	//group.Use(middleware.JWTAuthMiddleware())
	{
		group.POST("/join",controller.JoinCourse)
		group.GET("/list", controller.QueryCourseList)
		group.POST("/detail", controller.QueryCourseDetail)
		group.GET("/search/:name",controller.SearchCourse)
		group.GET("/getStudentCourse",controller.GetStudentCourse)
	}
}
