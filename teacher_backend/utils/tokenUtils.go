package utils

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"reflect"
	"snail/teacher_backend/models"
	"snail/teacher_backend/models/interfaces"
	"snail/teacher_backend/vo"
	"time"
)

const (
	TokenExpireDuration = time.Hour * 2
	Signature           = "snail"
)

var TokenSecret = []byte("snail")

func GenToken(user interface{}, userType int) (string, error) {
	info := new(vo.Token)
	info.User = user
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

func GetToken(org interface{}) (user interfaces.User, err error) {
	//t := reflect.TypeOf(org)
	v := reflect.ValueOf(org)
	typeValue := v.Elem().FieldByName("Type").Int()
	userIno := v.Elem().FieldByName("User").Interface()
	jsonString := genJson(userIno)
	if typeValue == 1 {
		teacher := new(models.Teacher)
		err = json.Unmarshal([]byte(jsonString), &teacher)
		user = teacher
	} else {
		assistance := new(models.Student)
		err = json.Unmarshal([]byte(jsonString), &assistance)
		user = assistance
	}
	return
}

func genJson(x interface{}) string {
	v := reflect.ValueOf(x)
	stringBuffer := new(bytes.Buffer)
	stringBuffer.WriteString("{")
	for index, val := range v.MapKeys() {
		stringBuffer.WriteString("\"")
		stringBuffer.WriteString(val.String())
		stringBuffer.WriteString("\":")
		if v.MapIndex(val).Elem().Kind() == reflect.Float64 {
			stringBuffer.WriteString(fmt.Sprintf("%v", v.MapIndex(val).Elem().Float()))
		} else {
			stringBuffer.WriteString("\"")
			stringBuffer.WriteString(fmt.Sprintf("%v", v.MapIndex(val)))
			stringBuffer.WriteString("\"")
		}
		if index != len(v.MapKeys())-1 {
			stringBuffer.WriteString(",")
		}
	}
	stringBuffer.WriteString("}")
	return stringBuffer.String()
}
