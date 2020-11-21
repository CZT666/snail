package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"log"
	"net/http"
	"snail/teacher_backend/common"
	"snail/teacher_backend/logic"
	"snail/teacher_backend/models"
	"snail/teacher_backend/utils"
	"snail/teacher_backend/vo"
)

func AddBlog(c *gin.Context) {
	blog := new(models.Blog)
	if err := c.ShouldBindBodyWith(&blog, binding.JSON); err != nil {
		log.Printf("Add Blog bind json failed: %v\n", err)
		c.JSON(http.StatusOK, common.BadResponse(common.ParamError))
		return
	}
	org, _ := c.Get("user")
	user, err := utils.GetToken(org)
	if err != nil {
		log.Printf("Blog controller get token failed: %v\n", err)
		c.JSON(http.StatusOK, common.BadResponse(common.ServerError))
		return
	}
	baseResponse := logic.AddBlog(blog, user)
	c.JSON(http.StatusOK, baseResponse)
	return
}

func QueryBlogList(c *gin.Context) {
	blogListRequest := new(vo.BlogListRequest)
	if err := c.ShouldBindBodyWith(&blogListRequest, binding.JSON); err != nil {
		log.Printf("Query blog list bind json failed: %v\n", err)
		c.JSON(http.StatusOK, common.BadResponse(common.ParamError))
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
		c.JSON(http.StatusOK, common.BadResponse(common.ParamError))
		return
	}
	baseResponse := logic.QueryBlogDetail(blog)
	c.JSON(http.StatusOK, baseResponse)
	return
}

func UpdateBlog(c *gin.Context) {
	blog := new(models.Blog)
	if err := c.ShouldBindBodyWith(&blog, binding.JSON); err != nil {
		log.Printf("Update blog bind json failed: %v\n", err)
		c.JSON(http.StatusOK, common.BadResponse(common.ParamError))
		return
	}
	baseResponse := logic.UpdateBlog(blog)
	c.JSON(http.StatusOK, baseResponse)
	return
}

func DeleteBlog(c *gin.Context) {
	blog := new(models.Blog)
	if err := c.ShouldBindBodyWith(&blog, binding.JSON); err != nil {
		log.Printf("Delete blog bind json failed: %v\n", err)
		c.JSON(http.StatusOK, common.BadResponse(common.ParamError))
		return
	}
	baseResponse := logic.DeleteBlog(blog)
	c.JSON(http.StatusOK, baseResponse)
	return
}
