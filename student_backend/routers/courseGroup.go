package routers

import (
	"github.com/gin-gonic/gin"
	"snail/student_bakcend/controller"
	"snail/student_bakcend/middleware"
)

func courseGroup(engine *gin.Engine) {
	group := engine.Group("/course")
	group.GET("/list", controller.QueryCourseList)
	group.GET("/search/:name",controller.SearchCourse)
	group.Use(middleware.JWTAuthMiddleware())
	{
		group.POST("/join",controller.JoinCourse)
		group.POST("/detail", controller.QueryCourseDetail)
		group.GET("/getStudentCourse",controller.GetStudentCourse)
	}

}
