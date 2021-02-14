package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"log"
	"net/http"
	"student_bakcend/logic"
	"student_bakcend/utils"
	"student_bakcend/vo"
)

func GetSelect(c *gin.Context) {
	blog := c.Param("blogID")
	if cast.ToInt64(blog) < 1{
		log.Printf("param error")
		c.JSON(http.StatusOK, vo.BadResponse(vo.ParamError))
		return
	}
	baseResponse := logic.GetSelect(blog)
	c.JSON(http.StatusOK, baseResponse)
	return
}

func GetSelectScore(c *gin.Context){
	answers := c.PostForm("answers")
	blog := c.PostForm("blog")
	org, _ := c.Get("user")
	student, err := utils.GetToken(org)
	if err != nil {
		log.Printf("Get token failed: %v\n", err)
		c.JSON(http.StatusOK, vo.BadResponse(vo.ServerError))
		return
	}
	//user := models.Student{
	//	ID: 111,
	//	StudentID: "20171003389",
	//	}
	baseResponse := logic.GetSelectScore(answers,blog,student.StudentID)
	c.JSON(http.StatusOK, baseResponse)
	return
}