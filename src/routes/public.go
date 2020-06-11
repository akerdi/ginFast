package routes

import (
	"ginFast/src/services"
	"github.com/gin-gonic/gin"
)

var PublicRoutes = []*Route{
	{
		Path: "/api/aa",
		Method: "POST",
		Middles: []gin.HandlerFunc{
			services.AA(),
		},
	},
}