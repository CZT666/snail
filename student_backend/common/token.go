package common

import (
	"github.com/dgrijalva/jwt-go"
	"student_bakcend/models"
)

type Token struct {
	Student models.Student         `json:"student"`
	jwt.StandardClaims
}
