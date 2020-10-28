package utils

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"snail/teacher_backend/models"
	"snail/teacher_backend/vo"
	"time"
)

const (
	TokenExpireDuration = time.Hour * 2
	Signature           = "snail"
)

var TokenSecret = []byte("snail")

func GenTeacherToken(teacher *models.Teacher, userType int) (string, error) {
	info := new(vo.Token)
	info.Teacher = teacher
	info.Type = userType
	return genToken(info)
}

func genToken(info *vo.Token) (string, error) {
	info.ExpiresAt = time.Now().Add(TokenExpireDuration).Unix()
	info.Issuer = Signature
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, info)
	return token.SignedString(TokenSecret)
}

func ParseToken(tokenString string) (*vo.Token, error) {
	token, err := jwt.ParseWithClaims(tokenString, &vo.Token{}, func(token *jwt.Token) (interface{}, error) {
		return TokenSecret, nil
	})
	if err != nil {
		return nil, err
	}
	if user, ok := token.Claims.(*vo.Token); ok && token.Valid {
		return user, nil
	}
	return nil, errors.New("invalid token")
}
