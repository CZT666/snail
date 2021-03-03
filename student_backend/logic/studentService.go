package logic

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis"
	"log"
	"student_bakcend/vo"
	"student_bakcend/dao"
	"student_bakcend/models"
	"student_bakcend/utils"
	"time"
)

const (
	resetKeyPreFix = "mail.reset."
)
type StudentToken struct {
	Student models.Student
	Token string
}
func AddStudent(student *models.Student) (baseResponse *vo.BaseResponse) {
	baseResponse = new(vo.BaseResponse)
	baseResponse.Code = vo.Success
	if isMailExist(student.Mail) {
		baseResponse.Code = vo.AccountExist
		return
	}
	if err := models.CreateStudent(student); err != nil {
		baseResponse.Code = vo.Error
		baseResponse.Msg = "添加失败"
		log.Printf("Student service create teacher failed: %v\n", err)
	}
	return
}

func StudentLogin(student *models.Student) (baseResponse *vo.BaseResponse) {
	baseResponse = new(vo.BaseResponse)
	baseResponse.Code = vo.Success
	studentID := student.StudentID
	pwd := student.Pwd
	log.Printf("studentID: %v pwd: %v\n", studentID, pwd)
	student, ok := isStudent(studentID, pwd)
	if ok {
		log.Printf("Student login: %v\n", studentID)
		tokenString, err := utils.GenToken(student)
		if err != nil {
			fmt.Printf("Generate token error: %v\n", err)
			baseResponse.Code = vo.TokenError
			return
		}
		var studentInfo models.Student
		studentInfo.StudentID = student.StudentID
		if _, err := models.GetStudent(&studentInfo);err != nil{
			fmt.Printf("get student error: %v\n", err)
			baseResponse.Code = vo.Error
			return
		}
		result := StudentToken{Student: studentInfo,Token: tokenString}
		baseResponse.Data = result
		return
	} else {
		baseResponse.Code = vo.Error
		baseResponse.Msg = "账号或密码错误"
		return
	}

}

func isStudent(studentID string, pwd string) (user *models.Student, ok bool) {
	student := new(models.Student)
	student.StudentID = studentID
	student.Pwd = pwd
	res, err := models.GetStudent(student)
	if err != nil {
		fmt.Printf("Student judge error: %v\n", err)
	}
	fmt.Printf("Len of res: %v\n", len(res))
	if len(res) > 0 {
		return &res[0], true
	} else {
		return &models.Student{}, false
	}
}


func isMailExist(mail string) bool {
	student := new(models.Student)
	student.Mail = mail
	studentList, err := models.GetStudent(student)
	if err != nil {
		log.Printf("Find mail error: %v\n", err)
		return false
	}
	return len(studentList) > 0
}

func ResetPwdReq(mail string) (baseResponse *vo.BaseResponse) {
	baseResponse = new(vo.BaseResponse)
	baseResponse.Code = vo.Success
	if !isMailExist(mail) {
		baseResponse.Code = vo.MailNotExist
		log.Printf("Mail already exists: %v\n", mail)
		return
	}
	proofString, err := utils.GenResetProof(mail)
	if err != nil {
		baseResponse.Code = vo.ServerError
		log.Printf("Generate mail proof failed: %v\n", err)
		return
	}
	// 将邮箱写入redis
	err = addResetReqToRedis(mail, proofString)
	if err != nil {
		baseResponse.Code = vo.ServerError
		log.Printf("Add reset req into redis failed: %v\n", err)
		return
	}
	// 将发送邮件请求发向消息队列
	err = sendResetReqToNSQ(mail, proofString)
	if err != nil {
		log.Printf("Send reset pwd req to nsq error: %v\n", err)
		// 回滚redis
		err = redisDeleteKey(mail)
		if err != nil {
			log.Printf("Mail reset req redis rollback failed: %v\n", err)
		}
		baseResponse.Code = vo.ServerError
	}
	return
}

func sendResetReqToNSQ(mail string, proof string) error {
	req := &vo.ResetPwdRequest{
		Mail:  mail,
		Proof: proof,
	}
	reqJson, _ := json.Marshal(req)
	return dao.ResetPwdNSQProducer.Publish("reset_pwd", reqJson)
}

func addResetReqToRedis(mail string, proof string) error {
	key := resetKeyPreFix + mail
	num, err := dao.RedisDB.Set(context.Background(), key, proof, 24*time.Hour).Result()
	log.Printf("Reset mail request add into redis, total: %v", num)
	return err
}

func redisDeleteKey(mail string) error {
	key := resetKeyPreFix + mail
	num, err := dao.RedisDB.Del(context.Background(), key).Result()
	log.Printf("Delete reset mail request redis, total: %v", num)
	return err
}

func UpdatePwd(newPwd string, proof string, mail string) (baseResponse *vo.BaseResponse) {
	baseResponse = new(vo.BaseResponse)
	baseResponse.Code = vo.Success
	redisKey := resetKeyPreFix + mail
	cacheInfo, err := dao.RedisDB.Get(context.Background(), redisKey).Result()
	if err == redis.Nil {
		baseResponse.Code = vo.ProofInvalid
		log.Printf("Reset Mail Proof Invalid")
		return
	} else if err != nil {
		baseResponse.Code = vo.ServerError
		log.Printf("Redis get key error: %v\n", err)
		return
	}
	ok := cacheInfo == proof
	if !ok {
		baseResponse.Code = vo.ProofInvalid
		return
	}
	student := new(models.Student)
	student.Mail = mail
	studentList, err := models.GetStudent(student)
	if len(studentList) != 1 {
		baseResponse.Code = vo.Error
		baseResponse.Msg = "用户不存在"
		return
	}
	studentList[0].Pwd = newPwd
	if err = models.UpdateStudent(&studentList[0]); err != nil {
		baseResponse.Code = vo.ServerError
	}
	_ = redisDeleteKey(mail)
	return
}
