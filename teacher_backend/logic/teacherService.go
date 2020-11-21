package logic

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis"
	"log"
	"snail/teacher_backend/common"
	"snail/teacher_backend/dao"
	"snail/teacher_backend/models"
	"snail/teacher_backend/utils"
	"snail/teacher_backend/vo"
	"time"
)

const (
	resetKeyPreFix = "mail.reset."
)

func AddTeacher(teacher *models.Teacher) (baseResponse *common.BaseResponse) {
	baseResponse = new(common.BaseResponse)
	baseResponse.Code = common.Success
	if isMailExist(teacher.Mail) {
		baseResponse.Code = common.AccountExist
		return
	}
	if err := models.CreateTeacher(teacher); err != nil {
		baseResponse.Code = common.Error
		baseResponse.Msg = "添加失败"
		log.Printf("Teacher service create teacher failed: %v\n", err)
	}
	return
}

func TeacherLogin(user *vo.LoginRequest) (baseResponse *common.BaseResponse) {
	baseResponse = new(common.BaseResponse)
	baseResponse.Code = common.Success
	account := user.Account
	pwd := user.Pwd
	log.Printf("account: %v pwd: %v\n", account, pwd)
	teacher, ok := isTeacher(account, pwd)
	if ok {
		log.Printf("Teacher login: %v\n", account)
		tokenString, err := utils.GenToken(teacher, 1)
		if err != nil {
			fmt.Printf("Generate token error: %v\n", err)
			baseResponse.Code = common.TokenError
			return
		}
		baseResponse.Data = tokenString
		return
	} else {
		student, ok := isAssistance(account, pwd)
		if ok {
			log.Printf("Assistance login: %v\n", account)
			tokenString, err := utils.GenToken(student, 2)
			if err != nil {
				fmt.Printf("Generate token error: %v\n", err)
				baseResponse.Code = common.TokenError
				return
			}
			baseResponse.Data = tokenString
			return
		}
	}
	baseResponse.Code = common.Error
	baseResponse.Code = "账号或密码错误"
	return
}

func isTeacher(mail string, pwd string) (user *models.Teacher, ok bool) {
	teacher := new(models.Teacher)
	teacher.Mail = mail
	teacher.Pwd = pwd
	res, err := models.GetTeacher(teacher)
	if err != nil {
		fmt.Printf("Teacher judge error: %v\n", err)
	}
	fmt.Printf("Len of res: %v\n", len(res))
	if len(res) > 0 {
		return &res[0], true
	} else {
		return &models.Teacher{}, false
	}
}

func isAssistance(studentID string, pwd string) (user *models.Student, ok bool) {
	assistance := new(models.Assistance)
	assistance.StuID = studentID
	assistanceList, err := models.GetAssistance(assistance)
	if err != nil {
		log.Printf("Teacher service get assistance failed: %v\n", err)
		return nil, false
	}
	if len(assistanceList) < 1 {
		log.Printf("User Invalid")
		return nil, false
	}
	student := new(models.Student)
	student.StudentID = studentID
	student.Pwd = pwd
	if err = models.GetSingleStudent(student); err != nil {
		log.Printf("Teacher service get single student failed: %v\n", err)
		return nil, false
	}
	return student, true
}

func isMailExist(mail string) bool {
	teacher := new(models.Teacher)
	teacher.Mail = mail
	teacherList, err := models.GetTeacher(teacher)
	if err != nil {
		log.Printf("Find mail error: %v\n", err)
		return false
	}
	return len(teacherList) > 0
}

func ResetPwdReq(mail string) (baseResponse *common.BaseResponse) {
	baseResponse = new(common.BaseResponse)
	baseResponse.Code = common.Success
	if !isMailExist(mail) {
		baseResponse.Code = common.MailNotExist
		log.Printf("Mail already exists: %v\n", mail)
		return
	}
	proofString, err := utils.GenResetProof(mail)
	if err != nil {
		baseResponse.Code = common.ServerError
		log.Printf("Generate mail proof failed: %v\n", err)
		return
	}
	// 将邮箱写入redis
	err = addResetReqToRedis(mail, proofString)
	if err != nil {
		baseResponse.Code = common.ServerError
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
		baseResponse.Code = common.ServerError
	}
	return
}

func sendResetReqToNSQ(mail string, proof string) error {
	req := &common.ResetPwdRequest{
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

func UpdatePwd(newPwd string, proof string, mail string) (baseResponse *common.BaseResponse) {
	baseResponse = new(common.BaseResponse)
	baseResponse.Code = common.Success
	redisKey := resetKeyPreFix + mail
	cacheInfo, err := dao.RedisDB.Get(context.Background(), redisKey).Result()
	if err == redis.Nil {
		baseResponse.Code = common.ProofInvalid
		log.Printf("Reset Mail Proof Invalid")
		return
	} else if err != nil {
		baseResponse.Code = common.ServerError
		log.Printf("Redis get key error: %v\n", err)
		return
	}
	ok := cacheInfo == proof
	if !ok {
		baseResponse.Code = common.ProofInvalid
		return
	}
	teacher := new(models.Teacher)
	teacher.Mail = mail
	teacherList, err := models.GetTeacher(teacher)
	if len(teacherList) != 1 {
		baseResponse.Code = common.Error
		baseResponse.Msg = "用户不存在"
		return
	}
	teacherList[0].Pwd = newPwd
	if err = models.UpdateTeacher(&teacherList[0]); err != nil {
		baseResponse.Code = common.ServerError
	}
	_ = redisDeleteKey(mail)
	return
}
