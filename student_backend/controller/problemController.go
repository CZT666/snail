package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"log"
	"net/http"
	"student_bakcend/logic"
	"student_bakcend/vo"
)

func GetProblem(c *gin.Context) {
	blog := c.Param("blogID")
	if cast.ToInt64(blog) < 1{
		log.Printf("param error")
		c.JSON(http.StatusOK, vo.BadResponse(vo.ParamError))
		return
	}
	baseResponse := logic.GetProblem(blog)
	c.JSON(http.StatusOK, baseResponse)
	return
}