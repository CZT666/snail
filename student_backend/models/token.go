package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"log"
	"reflect"
	"strconv"
	"strings"
	"student_bakcend/models/helper"
	"time"
)

type Token struct {
	User interface{} `json:"user"`
	jwt.StandardClaims
}

const (
	TokenExpireDuration = time.Hour * 2
	Signature           = "snail"
)

var TokenSecret = []byte("snail")

func GenToken(student *Student) (string, error) {
	info := new(Token)
	info.User = student
	return genToken(info)
}

func genToken(info *Token) (string, error) {
	info.ExpiresAt = time.Now().Add(TokenExpireDuration).Unix()
	info.Issuer = Signature
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, info)
	return token.SignedString(TokenSecret)
}

func ParseToken(tokenString string) (*Token, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Token{}, func(token *jwt.Token) (interface{}, error) {
		return TokenSecret, nil
	})
	if err != nil {
		return nil, err
	}
	if user, ok := token.Claims.(*Token); ok && token.Valid {
		return user, nil
	}
	return nil, errors.New("invalid token")
}


func GetToken(org interface{}) (user helper.User, err error) {
	//t := reflect.TypeOf(org)
	v := reflect.ValueOf(org)
	userIno := v.Elem().FieldByName("Student").Interface()
	fmt.Printf("user info: %v\n", userIno)
	jsonString := genJson(userIno)
	student := new(Student)
	err = json.Unmarshal([]byte(jsonString), &student)
	user = student
	return
}

func genJson(x interface{}) string {
	userInfo := fmt.Sprintf("%v", x)
	userInfo = strings.Trim(userInfo, "{")
	userInfo = strings.Trim(userInfo, "}")
	list := strings.Split(userInfo, " ")
	log.Printf("list: %v\n", list)
	student := new(Student)
	student.ID, _ = strconv.Atoi(list[0])
	student.Name = list[1]
	student.StudentID = list[2]
	student.Mail = list[3]
	student.Pwd = list[4]
	student.Gender, _ = strconv.Atoi(list[5])
	student.Faculty = list[6]
	student.Major = list[7]
	jsonString, err := json.Marshal(student)
	if err != nil {
		log.Printf("convert json failed: %v\n", err)
	}
	log.Printf("res: %v\n", string(jsonString))
	return string(jsonString)
}