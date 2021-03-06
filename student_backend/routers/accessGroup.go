package routers

import (
	"github.com/gin-gonic/gin"
	"snail/student_bakcend/controller"
)

func accessGroup(engine *gin.Engine) {
	engine.POST("/register", controller.StudentRegister)
	engine.POST("/login", controller.StudentLogin)
	engine.GET("/resetPwd/:mail", controller.ResetPwdReq)
	engine.POST("/resetPwd", controller.UpdatePwd)
}
