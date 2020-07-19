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
}