package routers

import (
	"github.com/gin-gonic/gin"
	"student_bakcend/controller"
)

func accessGroup(engine *gin.Engine) {
	engine.POST("/register", controller.StudentRegister)
	engine.GET("/login", controller.StudentLogin)
	engine.GET("/resetPwd/:mail", controller.ResetPwdReq)
	engine.POST("/resetPwd", controller.UpdatePwd)
}
