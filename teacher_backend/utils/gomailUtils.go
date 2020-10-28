package utils

import (
	"crypto/tls"
	"fmt"
	"log"
	"net"
	"net/smtp"
)

const (
	HOST     = "smtp.qq.com"
	PORT     = 465
	FROM     = "348673210@qq.com"
	PASSWORD = "yczkpzjnnrsocbcf"
)

func SendMail(toMail string, subject string, body string) error {
	header := make(map[string]string)
	header["From"] = "Snail<" + FROM + ">"
	header["To"] = toMail
	header["Subject"] = subject
	header["Content-type"] = "text/html; charset=UTF-8"
	var message string
	for key, value := range header {
		message += fmt.Sprintf("%s: %s\r\n", key, value)
	}
	message += "\r\n" + body
	auth := smtp.PlainAuth("", FROM, PASSWORD, HOST)
	err := SendMailUsingTLS(fmt.Sprintf("%s:%d", HOST, PORT), auth, FROM, []string{toMail}, []byte(message))
	return err
}

// 建立tcp连接
func Dial(address string) (*smtp.Client, error) {
	conn, err := tls.Dial("tcp", address, nil)
	if err != nil {
		log.Printf("Connect tmp fail: %v\n", err)
		return nil, err
	}
	host, _, _ := net.SplitHostPort(address)
	return smtp.NewClient(conn, host)
}

func SendMailUsingTLS(addr string, auth smtp.Auth, from string,
	to []string, msg []byte) (err error) {
	//create smtp client
	c, err := Dial(addr)
	if err != nil {
		log.Println("Create smpt client error:", err)
		return err
	}
	defer c.Close()
	if auth != nil {
		if ok, _ := c.Extension("AUTH"); ok {
			if err = c.Auth(auth); err != nil {
				log.Println("Error during AUTH", err)
				return err
			}
		}
	}
	if err = c.Mail(from); err != nil {
		return err
	}
	for _, addr := range to {
		if err = c.Rcpt(addr); err != nil {
			return err
		}
	}
	w, err := c.Data()
	if err != nil {
		return err
	}
	_, err = w.Write(msg)
	if err != nil {
		return err
	}
	err = w.Close()
	if err != nil {
		return err
	}
	return c.Quit()
}
