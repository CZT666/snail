package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"student_bakcend/vo"
	"student_bakcend/logic"
	"student_bakcend/models"
)

// TODO 邮箱校验
// TODO 密码加密
func StudentRegister(c *gin.Context) {
	student := new(models.Student)
	if err := c.BindJSON(&student); err != nil {
		log.Printf("Student register bind json error: %v\n", err)
		c.JSON(http.StatusOK, vo.BadResponse(vo.ParamError))
		return
	}
	baseResponse := logic.AddStudent(student)
	c.JSON(http.StatusOK, baseResponse)
	return
}

func StudentLogin(c *gin.Context) {
	student := new(models.Student)
	if err := c.BindJSON(&student); err != nil {
		fmt.Printf("Student login error: %v\n", err)
		c.JSON(http.StatusOK, vo.BadResponse(vo.ParamError))
		return
	}
	baseResponse := logic.StudentLogin(student)
	c.JSON(http.StatusOK, baseResponse)
	return
}

func ResetPwdReq(c *gin.Context) {
	mail := c.Param("mail")
	baseResponse := logic.ResetPwdReq(mail)
	c.JSON(http.StatusOK, baseResponse)
	return
}

func UpdatePwd(c *gin.Context) {
	newPwd := c.PostForm("pwd")
	proof := c.PostForm("proof")
	mail := c.PostForm("mail")
	baseResponse := logic.UpdatePwd(newPwd, proof, mail)
	c.JSON(http.StatusOK, baseResponse)
	return
}
