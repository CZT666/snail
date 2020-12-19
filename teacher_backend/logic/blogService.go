package logic

import (
	"log"
	"snail/teacher_backend/models"
	"snail/teacher_backend/models/helper"
	"snail/teacher_backend/models/interfaces"
	"snail/teacher_backend/vo"
	"time"
)

func AddBlog(blog *models.Blog, user interfaces.User) (baseResponse *vo.BaseResponse) {
	baseResponse = new(vo.BaseResponse)
	baseResponse.Code = vo.Success
	blog.Author = user.GetIdentity()
	blog.Type = user.GetType()
	now := time.Now()
	blog.CreateTime = now
	blog.UpdateTime = now
	if err := models.CreateBlog(blog); err != nil {
		log.Printf("Bolg service add blog failed: %v\n", err)
		baseResponse.Code = vo.Error
		baseResponse.Msg = "添加失败"
	}
	return
}

func QueryBlogList(request *vo.BlogListRequest) (baseResponse *vo.BaseResponse) {
	baseResponse = new(vo.BaseResponse)
	baseResponse.Code = vo.Success
	pageRequest := new(helper.PageRequest)
	pageRequest.Page = request.Page
	pageRequest.PageSize = request.PageSize
	blog := new(models.Blog)
	blog.CourseID = request.CourseID
	blogList, total, err := models.GetBlog(blog, pageRequest)
	if err != nil {
		log.Printf("Bolg service query blog list failed: %v\n", err)
		baseResponse.Code = vo.Error
		baseResponse.Msg = "查询失败"
	}
	pageResponse := new(helper.PageResponse)
	pageResponse.Data = blogList
	pageResponse.Total = total
	baseResponse.Data = pageResponse
	return
}

func QueryBlogDetail(blog *models.Blog) (baseResponse *vo.BaseResponse) {
	baseResponse = new(vo.BaseResponse)
	baseResponse.Code = vo.Success
	if err := models.GetSingleBlog(blog); err != nil {
		log.Printf("Blog service get single blog failed: %v\n", err)
		baseResponse.Code = vo.Error
		baseResponse.Msg = "查询失败"
		return
	}
	baseResponse.Data = blog
	return
}

func UpdateBlog(blog *models.Blog) (baseResponse *vo.BaseResponse) {
	baseResponse = new(vo.BaseResponse)
	baseResponse.Code = vo.Success
	blog.UpdateTime = time.Now()
	if err := models.UpdateBlog(blog); err != nil {
		log.Printf("Blog service update blog failed: %v\n", err)
		baseResponse.Code = vo.Error
		baseResponse.Msg = "更新失败"
	}
	return
}

func DeleteBlog(blog *models.Blog) (baseResponse *vo.BaseResponse) {
	baseResponse = new(vo.BaseResponse)
	baseResponse.Code = vo.Success

	if err := models.DeleteBlog(blog); err != nil {
		log.Printf("Blog service delete blog failed: %v\n", err)
		baseResponse.Code = vo.Error
		baseResponse.Msg = "删除失败"
	}
	return
}
