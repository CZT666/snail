package vo

import (
	"github.com/dgrijalva/jwt-go"
)

type LoginRequest struct {
	Account string `json:"account"`
	Pwd     string `json:"pwd"`
}

type Token struct {
	Type int         `json:"type"`
	User interface{} `json:"user"`
	jwt.StandardClaims
}
