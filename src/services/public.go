package services

import (
	"fmt"
	"ginFast/src/lib/mail"
	"ginFast/src/routes/validate/email"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"regexp"
)

var (
	emailQQPat = "@qq.com$"
	email163Pat = "@163.com$"
	emailSupportPat = fmt.Sprintf("%s|%s", emailQQPat, email163Pat)
	emailQQRegex = regexp.MustCompile(emailQQPat)
	email163Regex = regexp.MustCompile(email163Pat)
	emailSupportRegex = regexp.MustCompile(emailSupportPat)
)

func SendMail() gin.HandlerFunc  {
	return func(context *gin.Context) {
		var emailData struct {
			Email string
			Error string
		}
		var emailObj email.Email
		//if err := context.ShouldBind(&emailObj); err != nil {
		//	context.JSON(http.StatusBadRequest, nil)
		if err := context.BindQuery(&emailObj); err != nil {
			emailData.Error = err.Error()
			context.JSON(http.StatusBadRequest, emailData)
			return
		}
		log.Printf("emailObj emailObj: %v \n\n", emailObj)
		// TODO bEmail email 是否支持多个？ `"939713120@qq.com;767838865@qq.com"`
		// 当前支持一个email
		bEmail := []byte(emailObj.Email)
		if !emailSupportRegex.Match(bEmail) {
			emailData.Error = "Email only support for @qq.com and @163.com"
			context.JSON(http.StatusBadRequest, emailData)
			return
		}
		emailData.Email = emailObj.Email
		if emailQQRegex.Match(bEmail) {
			_ = mail.SendMailQQ([]string{emailObj.Email}, emailObj.Subject, emailObj.Msg)
		} else {
			_ = mail.SendMail163([]string{emailObj.Email}, emailObj.Subject, emailObj.Msg)
			//if err != nil {
			//	emailData.Error = err.Error()
			//	context.JSON(http.StatusBadRequest, emailData)
			//	return
			//}
		}
		context.JSON(http.StatusOK, emailData)
	}
}
