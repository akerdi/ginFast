package services

import (
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