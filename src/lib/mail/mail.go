package mail

import (
	"ginFast/src/config"
	"net/smtp"
	"strings"
)

func SendMailQQ(to []string, subject, body string) error {
	authToken := config.ConfigData.MailQQ.Token
	host := "smtp.qq.com"

	nickname := config.ConfigData.MailQQ.Nickname
	user := config.ConfigData.MailQQ.Sender
	return sendEmail(to, subject, body, user, nickname, authToken, host)
}

func sendEmail(to []string, subject, body, user, nickname, authToken, host string) error {
	content_type := "Content-Type text/plain; charset=UTF-8"
	msg := []byte("To:" + strings.Join(to, ",") + "\r\nFrom:" + nickname +
		"<" + user + ">\r\nSubject:" + subject + "\r\n" + content_type + "\r\n\r\n" + body)
	auth := smtp.PlainAuth("", user, authToken, host)
	err := smtp.SendMail(host+":25", auth, user, to, msg)
	if err != nil {
		return err
	}
	return nil
}

func SendMail163(to []string, subject, body string) error {
	authToken := config.ConfigData.MailQQ.Token
	user := config.ConfigData.MailQQ.Sender
	nickname := config.ConfigData.MailQQ.Nickname
	host := "smtp.163.com"
	return sendEmail(to, subject, body, user, nickname, authToken, host)
}

// 发送163邮件参考
// https://studygolang.com/articles/2745
