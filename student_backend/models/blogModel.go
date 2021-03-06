package models

import (
	"snail/student_bakcend/dao"
	"snail/student_bakcend/models/helper"
	"time"
)

type Blog struct {
	ID         int       `json:"id"`
	CourseID   int       `json:"course_id"`
	BlogTitle  string    `json:"blog_title"`
	Content    string    `json:"content"`
	Author     string    `json:"author"`
	Tag        string    `json:"tag"`
	Type       int       `json:"type"`
	CreateTime time.Time `json:"create_time"`
	UpdateTime time.Time `json:"update_time"`
}

func GetBlog(blog *Blog) (blogList []Blog, total int, err error) {
	if err := dao.DB.Where(&blog).Find(&blogList).Count(&total).Error; err != nil {
		return nil, 0, err
	}
	return
}

func GetSingleBlog(blog *Blog) (err error) {
	err = dao.DB.Where(&blog).First(&blog).Error
	return
}

func GetSearchBlog(pageRequest *helper.PageRequest, searchName string) (blogList []Blog, total int, err error) {
	page := pageRequest.Page
	pageSize := pageRequest.PageSize
	if err = dao.DB.Where("blog_title like ?", "%"+searchName+"%").Limit(pageSize).Offset((page - 1) * pageSize).Find(&blogList).Count(&total).Error; err != nil {
		return nil, 0, err
	}
	return
}