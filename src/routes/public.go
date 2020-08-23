package routes

import (
	"ginFast/src/services"
	"github.com/gin-gonic/gin"
)

var PublicRoutes = []*Route{
	{
		Path: "/api/sendMail",
		Method: "GET",
		Middles: []gin.HandlerFunc{
			services.SendMail(),
		},
	},
	{
		Path: "/api/startFilebeat", // TODO 增加限流方案
		Method: "GET",
		Middles: []gin.HandlerFunc{
			services.StartFilebeatRecenteUris(),
		},
	},
}