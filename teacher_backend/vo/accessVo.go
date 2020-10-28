package vo

import (
	"github.com/dgrijalva/jwt-go"
	"snail/teacher_backend/models"
)

type LoginRequest struct {
	Account string `json:"account"`
	Pwd     string `json:"pwd"`
}

type Token struct {
	Type      int             `json:"type"`
	Teacher   *models.Teacher `json:"teacher"`
	Assistant *models.Teacher `json:"assistant"`
	jwt.StandardClaims
}
