package logic

import (
	"encoding/json"
	"fmt"
	"log"
	"snail/teacher_backend/common"
	"snail/teacher_backend/dao"
	"snail/teacher_backend/models"
	"snail/teacher_backend/utils"
	"snail/teacher_backend/vo"
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
	}
	return
}

func TeacherLogin(user *vo.LoginRequest) (baseResponse *common.BaseResponse) {
	baseResponse = new(common.BaseResponse)
	baseResponse.Code = common.Success
	mail := user.Account
	pwd := user.Pwd
	teacher, ok := isTeacher(mail, pwd)
	if ok {
		tokenString, err := utils.GenTeacherToken(teacher, 0)
		if err != nil {
			fmt.Printf("Generate token error: %v\n", err)
			baseResponse.Code = common.TokenError
			return
		}
		baseResponse.Data = tokenString
		return
	}
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
		return
	}
	// 将邮箱写入redis
	// 将发送邮件请求发向消息队列
	err := sendResetReqToNSQ(mail)
	if err != nil {
		log.Printf("Send reset pwd req to nsq error: %v\n", err)
		// TODO 回滚redis
		baseResponse.Code = common.ServerError
	}
	return
}

func sendResetReqToNSQ(mail string) error {
	proofString, err := utils.GenResetProof(mail)
	if err != nil {
		return err
	}
	req := &common.ResetPwdRequest{
		Mail:  mail,
		Proof: proofString,
	}
	reqJson, _ := json.Marshal(req)
	return dao.ResetPwdNSQProducer.Publish("reset_pwd", reqJson)
}

func UpdatePwd(newPwd string, proof string, mail string) (baseResponse *common.BaseResponse) {
	baseResponse = new(common.BaseResponse)
	baseResponse.Code = common.Success
	ok, err := utils.ParseMailProof(proof, mail)
	if err != nil {
		log.Printf("Parse mail proof error: %v\n", err)
		baseResponse.Code = common.ServerError
		return
	}
	if !ok {
		baseResponse.Code = common.Error
		return
	}
	teacher := new(models.Teacher)
	teacher.Mail = mail
	teacherList, err := models.GetTeacher(teacher)
	if len(teacherList) != 1 {
		baseResponse.Code = common.Error
		return
	}
	teacherList[0].Pwd = newPwd
	if err = models.UpdateTeacher(&teacherList[0]); err != nil {
		baseResponse.Code = common.ServerError
	}
	return
}
