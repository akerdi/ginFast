package mail

import (
	"fmt"
	"ginFast/src/config"
	"net/smtp"
	"strings"
)

func SendMail_QQ(to []string, subject, body string) error {
	authToken := config.ConfigData.QQMail.Token
	auth := smtp.PlainAuth("", config.ConfigData.QQMail.Sender, authToken, "smtp.qq.com")
	nickname := config.ConfigData.QQMail.Nickname
	user := config.ConfigData.QQMail.Sender
	content_type := "Content-Type text/plain; charset=UTF-8"
	msg := []byte("To:" + strings.Join(to, ",") + "\r\nFrom:" + nickname +
		"<" + user + ">\r\nSubject:" + subject + "\r\n" + content_type + "\r\n\r\n" + body)
	err := smtp.SendMail("smtp.qq.com:25", auth, user, to, msg)
	if err != nil {
		return err
	}
	return nil
}

const (
	HOST        = "smtp.163.com"
	SERVER_ADDR = "smtp.163.com:25"
	USER = "18516636217@163.com"
	PASSWORD = ""
)

type Email struct {
	to      string "to"
	subject string "subject"
	msg     string "msg"
}

func NewEmail(to, subject, msg string) *Email {
	return &Email{to: to, subject: subject, msg: msg}
}
func SendMail_AA(to, subject, msg string) error {
	email := NewEmail(to, subject, msg)
	auth := smtp.PlainAuth("", USER, PASSWORD, HOST)
	sendTo := strings.Split(email.to, ";")
	done := make(chan error, 1024)
	go func() {
		defer close(done)
		for _, v := range sendTo {
			str := strings.Replace("From: "+USER+"~To: "+v+"~Subject: "+email.subject+"~~", "~", "\r\n", -1) + email.msg
			err := smtp.SendMail(
				SERVER_ADDR,
				auth,
				USER,
				[]string{v},
				[]byte(str),
			)
			fmt.Println("????? errrr::: ", err)
			done <- err
		}
	}()
	for i := 0; i < len(sendTo); i++ {
		<-done
	}
	return nil
}

// 发送163邮件参考
// https://studygolang.com/articles/2745
func SendMail_163(to []string, subject, body string) error {
	return nil
}
