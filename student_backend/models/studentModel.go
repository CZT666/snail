package models

import (
	"snail/student_bakcend/dao"
)

type Student struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	StudentID string `json:"student_id"`
	Mail      string `json:"mail"`
	Pwd       string `json:"pwd"`
	Gender    int    `json:"gender"`
	Faculty   string `json:"faculty"`
	Major     string `json:"major"`
	Img       string `json:"img"`
}

func (s Student) GetIdentity() string {
	return s.StudentID
}

func (s Student) GetName() string {
	return s.Name
}

func (s Student) GetType() int {
	return 2
}

func GetSingleStudent(student *Student) (err error) {
	err = dao.DB.Where(&student).First(&student).Error
	return
}
func CreateStudent(student *Student) (err error) {
	err = dao.DB.Create(&student).Error
	return
}

func GetStudent(student *Student) (studentList []Student, err error) {
	if err = dao.DB.Where(&student).Find(&studentList).Error; err != nil {
		return nil, err
	}
	return
}

func UpdateStudent(student *Student) (err error) {
	err = dao.DB.Save(&student).Error
	return
}