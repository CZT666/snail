package logic

import (
	"encoding/json"
	"fmt"
	"github.com/nsqio/go-nsq"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"snail/teacher_backend/common"
	"snail/teacher_backend/settings"
	"snail/teacher_backend/utils"
	"strings"
	"time"
)

const (
	resetSubject = "Snail-重置密码"
)

type MyHandler struct {
	Title string
}

func (myHandler *MyHandler) HandleMessage(msg *nsq.Message) (err error) {
	request := new(common.ResetPwdRequest)
	err = json.Unmarshal(msg.Body, &request)
	if err != nil {
		return
	}
	if request.Mail != "" {
		err = sendResetMail(request)
	}
	return
}

func sendResetMail(request *common.ResetPwdRequest) error {
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
	basePath, err := os.Getwd()
	if err != nil {
		log.Printf("Reset mail base path error: %v\n", basePath)
	}
	fileName := filepath.Join(basePath, "statics", "template", "ResetMail.html")
	log.Printf("file name: %v\n", fileName)
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Printf("read file error: %v\n", err)
	}
	content = string(data)
	content = strings.Replace(content, "-workHost", settings.Conf.WorkHost, -1)
	content = strings.Replace(content, "-workPort", settings.Conf.WorkPort, -1)
	content = strings.Replace(content, "-mailString", mail, -1)
	content = strings.Replace(content, "-proofString", proof, -1)
	fmt.Printf("html content:\n%v", content)
	return
}
