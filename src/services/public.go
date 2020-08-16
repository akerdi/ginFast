package services

import (
	"fmt"
	"ginFast/src/db"
	"ginFast/src/lib/mail"
	"ginFast/src/routes/validate/email"
	"github.com/gin-gonic/gin"
	jsoniter "github.com/json-iterator/go"
	"log"
	"net/http"
	"regexp"
	"strings"
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

type ResultObject struct {
	Timestamp string `json:"@timestamp"`
	Message string `json:"message"`
}

func ReadNginxAccessLogInRedis() {
	resultJson, err := db.RedisGetRangeByKey("filebeat:nginx:accesslog", 0, 100)
	if err != nil {
		panic(err)
	}
	
	var resObjects []ResultObject
	for _, str := range resultJson {
		var resObject ResultObject
		err = jsoniter.Unmarshal([]byte(str), &resObject)
		if err != nil {
			panic(err)
		}
		resObjects = append(resObjects, resObject)
	}
	
	
	exec(resObjects)
}

func exec(resObjects []ResultObject)  {
	for _, resObject := range resObjects {
		message := resObject.Message
		messageSplitArray := strings.Split(message, " - - ")
		ip := messageSplitArray[0]
		msg := messageSplitArray[1]
		isVPNMsg, _ := regexp.MatchString("/58ff4ec7/", msg)
		if isVPNMsg == true {
			continue
		}
		fmt.Println("222222==== ", ip, "     ip]]] [[[[msg: ", msg)
	}
}
