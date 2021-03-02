package logic

import (
	"log"
	"reflect"
	"snail/teacher_backend/models"
	"snail/teacher_backend/models/helper"
	"snail/teacher_backend/utils"
	"snail/teacher_backend/vo"
	"strconv"
	"time"
)

func AddCourse(course *models.Course, user helper.User) (baseResponse *vo.BaseResponse) {
	baseResponse = new(vo.BaseResponse)
	baseResponse.Code = vo.Success
	searchCode := utils.EncodeMD5(user.GetIdentity(), strconv.FormatInt(time.Now().Unix(), 10))
	course.SearchCode = searchCode
	course.CreateBy = user.GetIdentity()
	course.CreateTime = time.Now()
	if err := models.CreateCourse(course); err != nil {
		baseResponse.Code = vo.Error
		baseResponse.Msg = "添加失败"
		log.Printf("Course service create course failed: %v\n", err)
		return
	}
	baseResponse.Data = course
	return
}

func QueryCourseList(user helper.User) (baseResponse *vo.BaseResponse) {
	baseResponse = new(vo.BaseResponse)
	baseResponse.Code = vo.Success
	userType := user.GetType()
	log.Printf("Type of uer: %v\n", userType)
	// TODO 助教
	if userType == 1 {
		course := new(models.Course)
		course.CreateBy = user.GetIdentity()
		courseList, total, err := models.GetCourse(course)
		if err != nil {
			log.Printf("Query course list failed: %v\n", err)
			baseResponse.Code = vo.ServerError
			return
		}
		pageResponse := helper.NewPageResponse(total, courseList)
		baseResponse.Data = pageResponse
	} else {
		assistance := new(models.Assistance)
		assistance.StuID = user.GetIdentity()
		assistanceList, err := models.GetAssistance(assistance)
		if err != nil {
			log.Printf("Course service get assistance failed: %v\n", err)
			baseResponse.Code = vo.ServerError
			return
		}
		idList := make([]int, len(assistanceList))
		for index, a := range assistanceList {
			v := reflect.ValueOf(a)
			idList[index] = int(v.FieldByName("CourseID").Int())
		}
		courseList, total, err := models.GetCourseByID(idList)
		if err != nil {
			log.Printf("Course service get assistance by id failed: %v\n", err)
			baseResponse.Code = vo.ServerError
			return
		}
		pageResponse := helper.NewPageResponse(total, courseList)
		baseResponse.Data = pageResponse
	}
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

func UpdateCourse(course *models.Course) (baseResponse *vo.BaseResponse) {
	baseResponse = new(vo.BaseResponse)
	baseResponse.Code = vo.Success
	if err := models.UpdateCourse(course); err != nil {
		baseResponse.Code = vo.ServerError
		log.Printf("Update course failed: %v\n", err)
	}
	return
}

func DeleteCourse(course *models.Course) (baseResponse *vo.BaseResponse) {
	baseResponse = new(vo.BaseResponse)
	baseResponse.Code = vo.Success
	if err := models.DeleteCourse(course); err != nil {
		baseResponse.Code = vo.ServerError
		log.Printf("Delete course failed: %v\n", err)
	}
	return
}
