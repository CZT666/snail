package models

import (
	"fmt"
	"snail/student_bakcend/dao"
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
	ID        int    `json:"id"`
	CourseID  int    `json:"course_id"`
	StudentID string `json:"student_id"`
	IsValid   int 	 `json:"is_valid"`
}

func CreateCourseToStudent(courseToStudent *CourseToStudent) (err error) {
	err = dao.DB.Create(&courseToStudent).Error
	return
}

func GetCourse(course *Course) (courseList []Course, err error) {
	if err = dao.DB.Where(&course).Find(&courseList).Error; err != nil {
		return nil, err
	}
	fmt.Printf("course list msg:%v\n", courseList)
	return
}

func GetSearchCourse(searchName string) (courseList []Course, err error) {
	if err = dao.DB.Where("course_title like ?", "%"+searchName+"%").Find(&courseList).Error; err != nil {
		return nil, err
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

func GetCourseStudent(course *CourseToStudent)(courseList []CourseToStudent, err error){
	if err = dao.DB.Where(&course).Find(&courseList).Error; err != nil {
		return nil, err
	}
	fmt.Printf("course list msg:%v\n", courseList)
	return
}