package logic

import (
	"encoding/json"
	"fmt"
	"github.com/nsqio/go-nsq"
	"log"
	"snail/teacher_backend/settings"
	"snail/teacher_backend/utils"
	"snail/teacher_backend/vo"
	"time"
)

const (
	resetSubject = "Snail-重置密码"
)

type MyHandler struct {
	Title string
}

func (myHandler *MyHandler) HandleMessage(msg *nsq.Message) (err error) {
	request := new(vo.ResetPwdRequest)
	err = json.Unmarshal(msg.Body, &request)
	if err != nil {
		return
	}
	if request.Mail != "" {
		err = sendResetMail(request)
	}
	return
}

func sendResetMail(request *vo.ResetPwdRequest) error {
	mail := request.Mail
	proof := request.Proof
	err := utils.SendMail(mail, resetSubject, genContent(proof, mail))
	return err
}

func InitResetPwdConsumer(cfg *settings.ResetPwdConsumerConfig) (err error) {
	config := nsq.NewConfig()
	config.LookupdPollInterval = 15 * time.Second
	fmt.Printf("cfg: %v\n", cfg)
	conn, err := nsq.NewConsumer(cfg.Topic, cfg.Channel, config)
	if err != nil {
		log.Printf("Init reset pwd comsumer error: %v\n", err)
		return
	}
	consumer := &MyHandler{
		Title: "snail_teacher_backend",
	}
	conn.AddHandler(consumer)
	address := fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)
	if err = conn.ConnectToNSQLookupd(address); err != nil {
		return
	}
	return nil
}

func genContent(proof string, mail string) (content string) {
	url := fmt.Sprintf("http://%s:%d/teacherResetPwd?mail=%s?proof=%s", settings.Conf.WorkHost, settings.Conf.WorkPort, mail, proof)
	fmt.Printf("Reset pwd url: %v", url)
	content = fmt.Sprintf("<a href='%s' target='_blank'>请点击重置密码，有效期24小时</a>", url)
	return
}
