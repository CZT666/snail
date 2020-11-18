package logic

import (
	"log"
	"snail/teacher_backend/common"
	"snail/teacher_backend/models"
)

func AddAssistance(assistance *models.Assistance) (baseResponse *common.BaseResponse) {
	baseResponse = new(common.BaseResponse)
	baseResponse.Code = common.Success
	if isStudent(assistance.StuID) {
		if err := models.CreateAssistance(assistance); err != nil {
			log.Printf("Assistance Server create assistance failed: %v\n", err)
			baseResponse.Code = common.ServerError
		}
	} else {
		baseResponse.Code = common.Error
	}
	return
}

func isStudent(stuID string) bool {
	student := new(models.Student)
	student.StudentID = stuID
	err := models.GetSingleStudent(student)
	if err != nil {
		log.Printf("Assistance Service get single student failed: %v\n", err)
		return false
	}
	if student.ID < 1 {
		log.Printf("User Invalid: %v\n", student.StudentID)
		return false
	}
	return true
}
