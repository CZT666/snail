package common

import "github.com/dgrijalva/jwt-go"

type Token struct {
	Type int         `json:"type"`
	User interface{} `json:"user"`
	jwt.StandardClaims
}
