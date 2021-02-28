package controller

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"snail/teacher_backend/models/helper"
	"snail/teacher_backend/logic"
	"snail/teacher_backend/vo"
	"snail/teacher_backend/models"
	"github.com/gin-gonic/gin/binding"
)

func GetRedPoint(c *gin.Context)  {
	//org, _ := c.Get("user")
	//user, err := models.GetToken(org)
	var user helper.User
	//if err != nil {
	//	log.Printf("question controller get token failed: %v\n", err)
	//	c.JSON(http.StatusOK, vo.BadResponse(vo.ServerError))
	//	return
	//}
	baseResponse := logic.GetRedPoint(user)
	c.JSON(http.StatusOK, baseResponse)
	return
}

func GetAllQuestion(c *gin.Context)  {
	//org, _ := c.Get("user")
	//user, err := models.GetToken(org)
	var user helper.User
	//if err != nil {
	//	log.Printf("question controller get token failed: %v\n", err)
	//	c.JSON(http.StatusOK, vo.BadResponse(vo.ServerError))
	//	return
	//}
	pageRequest := helper.NewPageRequest()
	if err := c.BindJSON(&pageRequest); err != nil {
		log.Printf("get all question page request bind json failed: %v\n", err)
		c.JSON(http.StatusOK, vo.BadResponse(vo.ParamError))
		return
	}
	baseResponse := logic.GetAllQuestion(user,pageRequest)
	c.JSON(http.StatusOK, baseResponse)
	return
}

func GetSingleQuestion(c *gin.Context)  {
	questionID := c.Param("question_id")
	//org, _ := c.Get("user")
	//user, err := models.GetToken(org)
	var user helper.User
	//if err != nil {
	//	log.Printf("question controller get token failed: %v\n", err)
	//	c.JSON(http.StatusOK, vo.BadResponse(vo.ServerError))
	//	return
	//}
	baseResponse := logic.GetSingleQuestion(questionID,user)
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
	//org, _ := c.Get("user")
	//user, err := models.GetToken(org)
	var user helper.User
	//if err != nil {
	//	log.Printf("question controller get token failed: %v\n", err)
	//	c.JSON(http.StatusOK, vo.BadResponse(vo.ServerError))
	//	return
	//}
	baseResponse := logic.AddAnswer(answer, user)
	c.JSON(http.StatusOK, baseResponse)
	return
}

func GetAnswer(c *gin.Context){
	questionID := c.Param("question_id")
	pageRequest := helper.NewPageRequest()
	if err := c.BindJSON(&pageRequest); err != nil {
		log.Printf("get all answer page request bind json failed: %v\n", err)
		c.JSON(http.StatusOK, vo.BadResponse(vo.ParamError))
		return
	}
	baseResponse := logic.GetAnswer(questionID,pageRequest)
	c.JSON(http.StatusOK, baseResponse)
	return
}