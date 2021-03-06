package logic

import (
	"fmt"
	"log"
	"snail/student_bakcend/models"
	"snail/student_bakcend/vo"
)

func JoinCourse(courseToStudent *models.CourseToStudent)(baseResponse *vo.BaseResponse){
	baseResponse = new(vo.BaseResponse)
	baseResponse.Code = vo.Success
	if !isCourseExists(courseToStudent.CourseID){
		baseResponse.Code = vo.Error
		baseResponse.Msg = "course not exist"
		return
	}
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

func QueryCourseList() (baseResponse *vo.BaseResponse) {
	baseResponse = new(vo.BaseResponse)
	baseResponse.Code = vo.Success
	course := new(models.Course)
	courseList, err := models.GetCourse(course)
	if err != nil {
		log.Printf("Query course list failed: %v\n", err)
		baseResponse.Code = vo.ServerError
		return
	}
	baseResponse.Data = courseList
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

func SearchCourse(searchName string) (baseResponse *vo.BaseResponse) {
	baseResponse = new(vo.BaseResponse)
	baseResponse.Code = vo.Success
	courseList, err := models.GetSearchCourse(searchName)
	if err != nil {
		log.Printf("Query course list failed: %v\n", err)
		baseResponse.Code = vo.ServerError
		return
	}
	baseResponse.Data = courseList
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

func isCourseExists(courseID int)bool{
	course := new(models.Course)
	course.ID = courseID
	if err := models.GetSingleCourse(course); err!=nil{
		return false
	}
	return true
}

func GetStudentCourse(studentID string)  (baseResponse *vo.BaseResponse) {
	baseResponse = new(vo.BaseResponse)
	baseResponse.Code = vo.Success
	studentCourse := models.CourseToStudent{
		StudentID: studentID,
	}
	res,tmpErr := models.GetCourseStudent(&studentCourse)
	if tmpErr != nil{
		log.Printf("get course student failed: %v\n", tmpErr)
		baseResponse.Code = vo.ServerError
		return
	}
	fmt.Printf("result ***************:%v",res)
	var course []models.Course
	for i := range res{
		tmp := models.Course{ID: res[i].CourseID}
		if err:= models.GetSingleCourse(&tmp);err!=nil{
			log.Printf("get single course failed: %v\n", err)
			baseResponse.Code = vo.ServerError
			return
		}
		course = append(course,tmp)
	}
	baseResponse.Data = res
	return
}