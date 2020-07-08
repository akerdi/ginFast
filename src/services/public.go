package services

import (
	"fmt"
	"ginFast/src/lib/mail"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AA() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"happy": "aker",
		})
	}
}

func SendMail() gin.HandlerFunc  {
	return func(context *gin.Context) {
		//err := mail.SendMail_QQ([]string{"939713120@qq.com"}, "``tettetessss111==", "bobohdy")
		err := mail.SendMail_AA("939713120@qq.com;767838865@qq.com", "tetetetststetest", "godgodgodgodgodgodgod")
		if err != nil {
			fmt.Printf("send mail error: %v", err)
			context.Status(http.StatusBadRequest)
			return
		}
		context.JSON(http.StatusOK, gin.H{
			"success": true,
		})
	}
}
