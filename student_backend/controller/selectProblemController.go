package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"log"
	"net/http"
	"snail/student_bakcend/logic"
	"snail/student_bakcend/models"
	"snail/student_bakcend/vo"
)

func GetSelect(c *gin.Context) {
	blog := c.Param("blog_id")
	if cast.ToInt64(blog) < 1{
		log.Printf("param error")
		c.JSON(http.StatusOK, vo.BadResponse(vo.ParamError))
		return
	}
	baseResponse := logic.GetSelect(blog)
	c.JSON(http.StatusOK, baseResponse)
	return
}

func SelectScore(c *gin.Context){
	answers := c.PostForm("answers")
	blog := c.PostForm("blog_id")
	org, _ := c.Get("user")
	student, err := models.GetToken(org)
	if err != nil {
		log.Printf("Get token failed: %v\n", err)
		c.JSON(http.StatusOK, vo.BadResponse(vo.ServerError))
		return
	}
	//student := models.Student{
	//	ID: 111,
	//	StudentID: "20171003389",
	//	}
	baseResponse := logic.SelectScore(answers,blog,student.GetIdentity())
	c.JSON(http.StatusOK, baseResponse)
	return
}