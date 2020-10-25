package logic

import (
	"snail/teacher_backend/common"
	"snail/teacher_backend/models"
)

func AddTeacher(teacher *models.Teacher) (baseResponse *common.BaseResponse) {
	baseResponse = new(common.BaseResponse)
	baseResponse.Code = common.Success
	if err := models.CreateTeacher(teacher); err != nil {
		baseResponse.Code = common.Error
	}
	return
}

func UpdateTeacher(teacher *models.Teacher) (baseResponse *common.BaseResponse) {
	baseResponse = new(common.BaseResponse)
	baseResponse.Code = common.Success
	if err := models.UpdateTeacher(teacher); err != nil {
		baseResponse.Code = common.Error
	}
	return
}

func GetAllTeacher() (baseResponse *common.BaseResponse) {
	baseResponse = new(common.BaseResponse)
	baseResponse.Code = common.Success
	teacherList, err := models.GetAllTeacher()
	if err != nil {
		baseResponse.Code = common.Error
	}
	baseResponse.Data = teacherList
	return
}

func GetTeacher(teacher *models.Teacher) (baseResponse *common.BaseResponse) {
	baseResponse = new(common.BaseResponse)
	baseResponse.Code = common.Success
	teacherList, err := models.GetTeacher(teacher)
	if err != nil {
		baseResponse.Code = common.Error
	}
	baseResponse.Data = teacherList
	return
}

func DeleteTeacher(id string) (baseResponse *common.BaseResponse) {
	baseResponse = new(common.BaseResponse)
	baseResponse.Code = common.Success
	if err := models.DeleteTeacher(id); err != nil {
		baseResponse.Code = common.Error
	}
	return
}
