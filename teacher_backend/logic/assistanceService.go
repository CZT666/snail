package logic

import (
	"log"
	"snail/teacher_backend/models"
	"snail/teacher_backend/vo"
)

// TODO 助教无权限
func AddAssistance(assistance *models.Assistance) (baseResponse *vo.BaseResponse) {
	baseResponse = new(vo.BaseResponse)
	baseResponse.Code = vo.Success
	if isStudent(assistance.StuID) {
		if isAssistanceExist(assistance.StuID, assistance.CourseID) {
			log.Printf("Assistance exist.")
			baseResponse.Code = vo.Error
			baseResponse.Msg = "该学生已是课程助教"
			return
		}
		if err := models.CreateAssistance(assistance); err != nil {
			log.Printf("Assistance Server create assistance failed: %v\n", err)
			baseResponse.Code = vo.ServerError
		}
	} else {
		baseResponse.Code = vo.Error
		baseResponse.Msg = "添加助教失败"
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

func isAssistanceExist(stuID string, courseID int) bool {
	assistance := new(models.Assistance)
	assistance.StuID = stuID
	assistance.CourseID = courseID
	assistanceList, err := models.GetAssistance(assistance)
	if err != nil {
		log.Printf("Assistance service get assistance failed: %v\n", err)
		return false
	}
	return len(assistanceList) != 0
}

func DeleteAssistance(assistance *models.Assistance) (baseResponse *vo.BaseResponse) {
	baseResponse = new(vo.BaseResponse)
	baseResponse.Code = vo.Success
	if err := models.DeleteAssistance(assistance); err != nil {
		log.Printf("Assistance service delete assistance failed: %v\n", err)
		baseResponse.Code = vo.Error
		baseResponse.Msg = "删除失败"
	}
	return baseResponse
}

func FindStudent(student *models.Student) (baseResponse *vo.BaseResponse) {
	baseResponse = new(vo.BaseResponse)
	baseResponse.Code = vo.Success
	if err := models.GetSingleStudent(student); err != nil {
		log.Printf("Assistance service find student failed: %v\n", err)
		baseResponse.Code = vo.Error
		baseResponse.Msg = "查找失败"
	}
	baseResponse.Data = student
	return
}

func GetAllAssistance(assistance *models.Assistance) (baseResponse *vo.BaseResponse) {
	baseResponse = new(vo.BaseResponse)
	baseResponse.Code = vo.Success
	assistanceList, err := models.GetAssistance(assistance)
	if err != nil {
		log.Printf("get assistance failed: %v\n", err)
		baseResponse.Code = vo.Error
		return baseResponse
	}
	var stuList []*models.Student
	for _, stu := range assistanceList {
		tmp := new(models.Student)
		tmp.StudentID = stu.StuID
		if err := models.GetSingleStudent(tmp); err != nil {
			log.Printf("find student failed: %v\n", err)
			continue
		}
		stuList = append(stuList, tmp)
	}
	baseResponse.Data = stuList
	return baseResponse
}
