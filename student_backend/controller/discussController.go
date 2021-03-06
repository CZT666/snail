package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"log"
	"net/http"
	"snail/student_bakcend/logic"
	"snail/student_bakcend/models"
	"snail/student_bakcend/vo"
)

func AddQuestion(c *gin.Context){
	question := new(models.Question)
	if err := c.ShouldBindBodyWith(&question, binding.JSON); err != nil {
		fmt.Printf("Add question bind json failed: %v\n", err)
		c.JSON(http.StatusOK, vo.BadResponse(vo.ParamError))
		return
	}
	org, _ := c.Get("user")
	user, err := models.GetToken(org)
	fmt.Printf("***************************************")
	fmt.Println(user)
	//var user helper.User
	if err != nil {
		log.Printf("question controller get token failed: %v\n", err)
		c.JSON(http.StatusOK, vo.BadResponse(vo.ServerError))
		return
	}
	baseResponse := logic.AddQuestion(question, user)
	c.JSON(http.StatusOK, baseResponse)
	return
}

func GetAllQuestion(c *gin.Context){
	courseID := c.Param("course_id")
	baseResponse := logic.GetAllQuestion(courseID)
	c.JSON(http.StatusOK, baseResponse)
	return
}

func GetSingleQuestion(c *gin.Context){
	questionID := c.Param("question_id")
	baseResponse := logic.GetSingleQuestion(questionID)
	c.JSON(http.StatusOK, baseResponse)
	return
}

func SearchQuestion(c *gin.Context)  {
	searchName := c.Param("name")
	courseID := c.Param("course_id")
	baseResponse := logic.SearchQuestion(searchName,courseID)
	c.JSON(http.StatusOK, baseResponse)
	return
}

func AddAnswer(c *gin.Context){
	answer := new(models.Answer)
	if err := c.ShouldBindBodyWith(&answer, binding.JSON); err != nil {
		log.Printf("Add answer bind json failed: %v\n", err)
		c.JSON(http.StatusOK, vo.BadResponse(vo.ParamError))
		return
	}
	org, _ := c.Get("user")
	user, err := models.GetToken(org)
	//var user helper.User
	if err != nil {
		log.Printf("question controller get token failed: %v\n", err)
		c.JSON(http.StatusOK, vo.BadResponse(vo.ServerError))
		return
	}
	baseResponse := logic.AddAnswer(answer, user)
	c.JSON(http.StatusOK, baseResponse)
	return
}

func GetAnswer(c *gin.Context){
	questionID := c.Param("question_id")
	baseResponse := logic.GetAnswer(questionID)
	c.JSON(http.StatusOK, baseResponse)
	return
}
