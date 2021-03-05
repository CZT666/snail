package routers

import (
	"github.com/gin-gonic/gin"
	"student_bakcend/controller"
	"student_bakcend/middleware"
)

func discussGroup(engine *gin.Engine){
	groupDiscuss := engine.Group("/discuss")
	groupDiscuss.Use(middleware.JWTAuthMiddleware())
	{
		groupDiscuss.POST("/addQuestion",controller.AddQuestion)
		groupDiscuss.GET("/getAllQuestion/:course_id",controller.GetAllQuestion)
		groupDiscuss.GET("/getSingleQuestion/:question_id",controller.GetSingleQuestion)
		groupDiscuss.GET("/search/:course_id/:name",controller.SearchQuestion)
		groupDiscuss.POST("/addAnswer",controller.AddAnswer)
		groupDiscuss.GET("/getAnswer/:question_id",controller.GetAnswer)

	}

}