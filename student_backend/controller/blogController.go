package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"log"
	"net/http"
	"student_bakcend/logic"
	"student_bakcend/models"
	"student_bakcend/vo"
)


func QueryBlogList(c *gin.Context) {
	blogListRequest := new(vo.BlogListRequest)
	if err := c.ShouldBindBodyWith(&blogListRequest, binding.JSON); err != nil {
		log.Printf("Query blog list bind json failed: %v\n", err)
		c.JSON(http.StatusOK, vo.BadResponse(vo.ParamError))
		return
	}
	baseResponse := logic.QueryBlogList(blogListRequest)
	c.JSON(http.StatusOK, baseResponse)
	return
}

func QueryBlogDetail(c *gin.Context) {
	blog := new(models.Blog)
	if err := c.ShouldBindBodyWith(&blog, binding.JSON); err != nil {
		log.Printf("Query blog detail bind json failed: %v\n", err)
		c.JSON(http.StatusOK, vo.BadResponse(vo.ParamError))
		return
	}
	baseResponse := logic.QueryBlogDetail(blog)
	c.JSON(http.StatusOK, baseResponse)
	return
}

