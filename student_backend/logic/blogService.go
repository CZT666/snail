package logic

import (
	"log"
	"snail/student_bakcend/models"
	"snail/student_bakcend/models/helper"
	"snail/student_bakcend/vo"
)


func QueryBlogList(request *vo.BlogListRequest) (baseResponse *vo.BaseResponse) {
	baseResponse = new(vo.BaseResponse)
	baseResponse.Code = vo.Success
	blog := new(models.Blog)
	blog.CourseID = request.CourseID
	blogList, total, err := models.GetBlog(blog)
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

