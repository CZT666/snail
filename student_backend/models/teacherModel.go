package models

import "snail/student_bakcend/dao"

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

func (t Teacher) GetIdentity() string {
	return t.Mail
}

func (t Teacher) GetName() string {
	return t.Name
}

func (t Teacher) GetType() int {
	return 1
}


func GetAllTeacher() (teacherList []Teacher, err error) {
	if err = dao.DB.Find(&teacherList).Error; err != nil {
		return nil, err
	}
	return
}

func GetTeacher(teacher *Teacher) (teacherList []Teacher, err error) {
	if err = dao.DB.Where(&teacher).Find(&teacherList).Error; err != nil {
		return nil, err
	}
	return
}


