package models

import "snail/teacher_backend/dao"

type Teacher struct {
	ID             int    `json:"id"`
	Name           string `json:"name"`
	Mail           string `json:"mail"`
	Pwd            string `json:"pwd"`
	Gender         int    `json:"gender"`
	Faculty        string `json:"faculty"`
	Major          string `json:"major"`
	Phone          string `json:"phone"`
	Address        string `json:"address"`
	Img            string `json:"img"`
	SocialAccounts string `json:"social_accounts"`
}

func CreateTeacher(teacher *Teacher) (err error) {
	err = dao.DB.Create(&teacher).Error
	return
}

func UpdateTeacher(teacher *Teacher) (err error) {
	err = dao.DB.Save(&teacher).Error
	return
}

func GetAllTeacher() (teacherList *[]Teacher, err error) {
	if err = dao.DB.Find(&teacherList).Error; err != nil {
		return nil, err
	}
	return
}

func GetTeacher(teacher *Teacher) (teacherList *[]Teacher, err error) {
	if err = dao.DB.Where(&teacher).Find(&teacherList).Error; err != nil {
		return nil, err
	}
	return
}

func DeleteTeacher(id string) (err error) {
	err = dao.DB.Where("id=?", id).Error
	return
}
