package dao

import (
	"fmt"
	"github.com/nsqio/go-nsq"
	"log"
	"snail/student_bakcend/settings"
)

var ResetPwdNSQProducer *nsq.Producer

func InitResetPwdNSQ(cfg *settings.NSQConfig) (err error) {
	address := fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)
	config := nsq.NewConfig()
	ResetPwdNSQProducer, err = nsq.NewProducer(address, config)
	if err != nil {
		log.Printf("Init reset pwd nsq error: %v\n", err)
		return
	}
	return
}
