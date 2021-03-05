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

func GetProblem(c *gin.Context) {
	blogID := c.Param("blog_id")
	if cast.ToInt64(blogID) < 1{
		log.Printf("param error")
		c.JSON(http.StatusOK, vo.BadResponse(vo.ParamError))
		return
	}
	baseResponse := logic.GetProblem(blogID)
	c.JSON(http.StatusOK, baseResponse)
	return
}

func GetProblemScore(c *gin.Context)  {
	blogID := c.Param("blog_id")
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
	if cast.ToInt64(blogID) < 1{
		log.Printf("param error")
		c.JSON(http.StatusOK, vo.BadResponse(vo.ParamError))
		return
	}
	baseResponse := logic.GetProblemScore(blogID,student.GetIdentity())
	c.JSON(http.StatusOK, baseResponse)
	return
}