package vo

import "snail/teacher_backend/models"

type LoginRequest struct {
	Account string `json:"account"`
	Pwd     string `json:"pwd"`
}

type LoginRespone struct {
	UserType int `json:"type"`
	Teacher *models.Teacher `json:"teacher"`
	Assistance *models.Student `json:"assistance"`
	Token string `json:"token"`
}
