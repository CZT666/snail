package models

import (
	"snail/teacher_backend/dao"
	"snail/teacher_backend/models/helper"
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

func CreateBlog(blog *Blog) (err error) {
	err = dao.DB.Create(&blog).Error
	return
}

func GetBlog(blog *Blog, request *helper.PageRequest) (blogList []Blog, total int, err error) {
	page := request.Page
	pageSize := request.PageSize
	if err := dao.DB.Where(&blog).Limit(pageSize).Offset((page - 1) * pageSize).Find(&blogList).Count(&total).Error; err != nil {
		return nil, 0, err
	}
	return
}

func GetSingleBlog(blog *Blog) (err error) {
	err = dao.DB.Where(&blog).First(&blog).Error
	return
}

func UpdateBlog(blog *Blog) (err error) {
	err = dao.DB.Model(&Blog{}).Updates(&blog).Error
	return
}

func DeleteBlog(blog *Blog) (err error) {
	err = dao.DB.Delete(&blog).Error
	return
}
