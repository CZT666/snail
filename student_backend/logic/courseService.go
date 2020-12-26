package logic

import (
	"fmt"
	"log"
	"student_bakcend/models/helper"
	"student_bakcend/models"
	"student_bakcend/vo"
)

func JoinCourse(courseToStudent *models.CourseToStudent)(baseResponse *vo.BaseResponse){
	baseResponse = new(vo.BaseResponse)
	baseResponse.Code = vo.Success
	if isCoursePrivate(courseToStudent.CourseID){
		courseToStudent.IsValid = 0
	}else{
		courseToStudent.IsValid = 1
	}
	if err := models.CreateCourseToStudent(courseToStudent); err != nil {
		baseResponse.Code = vo.Error
		baseResponse.Msg = "add course to student fail"
		log.Printf("CourseToStudent service create course failed: %v\n", err)
		return
	}

	return
}

func QueryCourseList(pageRequest *helper.PageRequest) (baseResponse *vo.BaseResponse) {
	baseResponse = new(vo.BaseResponse)
	baseResponse.Code = vo.Success
	course := new(models.Course)
	courseList, total, err := models.GetCourse(course, pageRequest)
	if err != nil {
		log.Printf("Query course list failed: %v\n", err)
		baseResponse.Code = vo.ServerError
		return
	}
	pageResponse := helper.NewPageResponse(total, courseList)
	baseResponse.Data = pageResponse
	return
}

func QueryCourseDetail(course *models.Course) (baseResponse *vo.BaseResponse) {
	baseResponse = new(vo.BaseResponse)
	baseResponse.Code = vo.Success
	err := models.GetSingleCourse(course)
	if err != nil {
		log.Printf("Get single course failed: %v\n", err)
		baseResponse.Code = vo.ServerError
		return
	}
	baseResponse.Data = course
	return
}

func SearchCourse(course *models.Course, pageRequest *helper.PageRequest,searchName string) (baseResponse *vo.BaseResponse) {
	baseResponse = new(vo.BaseResponse)
	baseResponse.Code = vo.Success
	courseList, total, err := models.GetSearchCourse(course, pageRequest,searchName)
	if err != nil {
		log.Printf("Query course list failed: %v\n", err)
		baseResponse.Code = vo.ServerError
		return
	}
	pageResponse := helper.NewPageResponse(total, courseList)
	baseResponse.Data = pageResponse
	return
}

func isCoursePrivate(courseID int) bool {
	course := new(models.Course)
	course.ID = courseID
	if err := models.GetSingleCourse(course); err != nil {
		log.Printf("Course operation middle ware get single course failed: %v\n", err)
		return false
	}
	fmt.Printf("course msg is:%v\n",course)
	return course.IsPrivate == 1
}