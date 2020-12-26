package models

import (
	"fmt"
	"student_bakcend/dao"
	"student_bakcend/models/helper"
	"time"
)

type Course struct {
	ID          int       `json:"id"`
	CourseTitle string    `json:"course_title"`
	IsPrivate   int       `json:"is_private"`
	Intro       string    `json:"intro"`
	Goal        string    `json:"goal"`
	UserFor     string    `json:"user_for"`
	SearchCode  string    `json:"search_code"`
	CreateBy    string    `json:"create_by"`
	CreateTime  time.Time `json:"create_time"`
}
type CourseToStudent struct {
	ID        int `json:"id"`
	CourseID  int `json:"course_id"`
	StudentID int `json:"student_id"`
	IsValid   int `json:"is_valid"`
}

func CreateCourseToStudent(courseToStudent *CourseToStudent) (err error) {
	err = dao.DB.Create(&courseToStudent).Error
	return
}

func GetCourse(course *Course, pageRequest *helper.PageRequest) (courseList []Course, total int, err error) {
	page := pageRequest.Page
	pageSize := pageRequest.PageSize
	if err = dao.DB.Where(&course).Limit(pageSize).Offset((page - 1) * pageSize).Find(&courseList).Count(&total).Error; err != nil {
		return nil, 0, err
	}
	fmt.Printf("course list msg:%v\n", courseList)
	return
}

func GetSearchCourse(course *Course, pageRequest *helper.PageRequest, searchName string) (courseList []Course, total int, err error) {
	page := pageRequest.Page
	pageSize := pageRequest.PageSize
	if err = dao.DB.Where("course_title like ?", "%"+searchName+"%").Limit(pageSize).Offset((page - 1) * pageSize).Find(&courseList).Count(&total).Error; err != nil {
		return nil, 0, err
	}
	return
}

func GetSingleCourse(course *Course) (err error) {
	err = dao.DB.Where(&course).First(&course).Error
	return
}
func MatchCourseStudent(course *CourseToStudent) (err error) {
	err = dao.DB.Where(&course).First(&course).Error
	return
}
