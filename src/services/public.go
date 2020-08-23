package services

import (
	"fmt"
	"ginFast/src/config"
	"ginFast/src/db"
	"ginFast/src/lib/mail"
	"ginFast/src/routes/validate/email"
	"github.com/gin-gonic/gin"
	jsoniter "github.com/json-iterator/go"
	"io/ioutil"
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

func StartFilebeatRecenteUris() gin.HandlerFunc {
	return func(context *gin.Context) {
		token := context.Query("token")
		if token != config.ConfigData.Token {
			context.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"msg": "时间不对", // 给出不对原因(错误误导原因, 防止被攻击)
			})
			return
		}
		ReadNginxAccessLogInRedis()
		context.JSON(http.StatusOK, gin.H{
			"success": true,
		})
	}
}

type ResultObject struct {
	Timestamp string `json:"@timestamp"`
	Message string `json:"message"`
}

var (
	theTargetPattern = []string{
		"/webfig/",
		"/wp-login.php/",
		"robots.txt/",
		"/api/jsonws/invoke",
		"/?XDEBUG_SESSION_START=phpstorm",
		"/invokefunction",
		"/.env/",
		"/cgi-bin/config.exp",
		"/boaform/admin/formLogin",
		"/index.php",
		// more to add
	}
	// 最多一个任务在进行当中
	readNginxAccessLogTaskIsProcessing = false
)

func ReadNginxAccessLogInRedis() {
	if readNginxAccessLogTaskIsProcessing == true {
		return
	}
	readNginxAccessLogTaskIsProcessing = true
	resultJson, err := db.RedisGetRangeByKey("filebeat:nginx:accesslog", 0, 5000)
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
	execResultObject(resObjects)
}

func execResultObject(resObjects []ResultObject)  {
	for _, resObject := range resObjects {
		message := resObject.Message
		messageSplitArray := strings.Split(message, " - - ")
		if len(messageSplitArray) != 2 {
			fmt.Println("获取到的链接分割成的数目不为2！！！")
			break
		}
		ip := messageSplitArray[0]
		msg := messageSplitArray[1]
		shouldContinueOuter := false
		for _, filterPattern := range config.ConfigData.FilterPattern {
			isFilterMsg, _ := regexp.MatchString(filterPattern, msg)
			if isFilterMsg == true {
				shouldContinueOuter = true
				break
			}
		}
		if shouldContinueOuter == true {
			continue
		}
		for _, targetPattern := range theTargetPattern {
			isTargetMsg, _ := regexp.MatchString(targetPattern, msg)
			if isTargetMsg {
				// 调用外部拦截设备
				fmt.Printf("====== block the fucking ip: %s, because reason: %s \n\n", ip, msg)
				go func() {
					res, err := http.Get(fmt.Sprintf("http://%s:9111/api/block/%s/%s", config.ConfigData.IP, "hash:"+ip, ip))
					if err != nil {
						fmt.Println("发送消息至 localhost:9111/api/block 出错了 err: ", err)
						return
					}
					body, err := ioutil.ReadAll(res.Body)
					if err != nil {
						fmt.Println("ioutil.ReadAll err: ", err)
						return
					}
					fmt.Printf("发送消息至 localhost:9111/api/block 状态 %d %v \n", res.StatusCode, string(body))
				}()
				shouldContinueOuter = true
				break
			}
		}
		if shouldContinueOuter == true {
			continue
		}
	}
	readNginxAccessLogTaskIsProcessing = false
}
