package logic

import (
	"log"
	"reflect"
	"snail/teacher_backend/common"
	"snail/teacher_backend/models"
	"snail/teacher_backend/models/interfaces"
	"snail/teacher_backend/utils"
	"snail/teacher_backend/vo"
	"strconv"
	"time"
)

func AddCourse(course *models.Course, user interfaces.User) (baseResponse *common.BaseResponse) {
	baseResponse = new(common.BaseResponse)
	baseResponse.Code = common.Success
	searchCode := utils.EncodeMD5(user.GetIdentity(), strconv.FormatInt(time.Now().Unix(), 10))
	course.SearchCode = searchCode
	course.CreateBy = user.GetIdentity()
	course.CreateTime = time.Now()
	if err := models.CreateCourse(course); err != nil {
		baseResponse.Code = common.Error
		log.Printf("Course service create course failed: %v\n", err)
		return
	}
	baseResponse.Data = course
	return
}

func QueryCourseList(user interfaces.User, pageRequest *vo.PageRequest) (baseResponse *common.BaseResponse) {
	baseResponse = new(common.BaseResponse)
	baseResponse.Code = common.Success
	userType := user.GetType()
	log.Printf("Type of uer: %v\n", userType)
	// TODO 助教
	if userType == 1 {
		course := new(models.Course)
		course.CreateBy = user.GetIdentity()
		courseList, total, err := models.GetCourse(course, pageRequest)
		if err != nil {
			log.Printf("Query course list failed: %v\n", err)
			baseResponse.Code = common.ServerError
			return
		}
		pageResponse := vo.NewPageResponse(total, courseList)
		baseResponse.Data = pageResponse
	} else {
		assistance := new(models.Assistance)
		assistance.StuID = user.GetIdentity()
		assistanceList, err := models.GetAssistance(assistance)
		if err != nil {
			log.Printf("Course service get assistance failed: %v\n", err)
			baseResponse.Code = common.ServerError
			return
		}
		idList := make([]int, len(assistanceList))
		for index, a := range assistanceList {
			v := reflect.ValueOf(a)
			idList[index] = int(v.FieldByName("CourseID").Int())
		}
		courseList, total, err := models.GetCourseByID(idList, pageRequest)
		if err != nil {
			log.Printf("Course service get assistance by id failed: %v\n", err)
			baseResponse.Code = common.ServerError
			return
		}
		pageResponse := vo.NewPageResponse(total, courseList)
		baseResponse.Data = pageResponse
	}
	return
}

func QueryCourseDetail(course *models.Course) (baseResponse *common.BaseResponse) {
	baseResponse = new(common.BaseResponse)
	baseResponse.Code = common.Success
	err := models.GetSingleCourse(course)
	if err != nil {
		log.Printf("Get single course failed: %v\n", err)
		baseResponse.Code = common.ServerError
		return
	}
	baseResponse.Data = course
	return
}

func UpdateCourse(course *models.Course) (baseResponse *common.BaseResponse) {
	baseResponse = new(common.BaseResponse)
	baseResponse.Code = common.Success
	if err := models.UpdateCourse(course); err != nil {
		baseResponse.Code = common.ServerError
		log.Printf("Update course failed: %v\n", err)
	}
	return
}

func DeleteCourse(course *models.Course) (baseResponse *common.BaseResponse) {
	baseResponse = new(common.BaseResponse)
	baseResponse.Code = common.Success
	if err := models.DeleteCourse(course); err != nil {
		baseResponse.Code = common.ServerError
		log.Printf("Delete course failed: %v\n", err)
	}
	return
}
