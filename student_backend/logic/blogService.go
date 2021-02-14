package logic

import (
	"log"
	"student_bakcend/models"
	"student_bakcend/models/helper"
	"student_bakcend/vo"
)


func QueryBlogList(request *vo.BlogListRequest) (baseResponse *vo.BaseResponse) {
	baseResponse = new(vo.BaseResponse)
	baseResponse.Code = vo.Success
	pageRequest := new(helper.PageRequest)
	pageRequest.Page = request.Page
	pageRequest.PageSize = request.PageSize
	blog := new(models.Blog)
	blog.ID = request.BlogID
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

