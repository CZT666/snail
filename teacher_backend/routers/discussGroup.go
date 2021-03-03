package routers

import (
	"github.com/gin-gonic/gin"
	"snail/teacher_backend/controller"
	"snail/teacher_backend/middleware"
)

func discussGroup(engine *gin.Engine){
	groupDiscuss := engine.Group("/discuss")
	groupDiscuss.Use(middleware.JWTAuthMiddleware())
	{
		groupDiscuss.GET("getRedPoint",controller.GetRedPoint)
		groupDiscuss.GET("getAllQuestion",controller.GetAllQuestion)
		groupDiscuss.GET("getSingleQuestion/:question_id",controller.GetSingleQuestion)
		groupDiscuss.POST("addAnswer",controller.AddAnswer)
		groupDiscuss.GET("getAnswer/:question_id",controller.GetAnswer)
	}
}