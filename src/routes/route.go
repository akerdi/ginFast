package routes

import "github.com/gin-gonic/gin"

type Route struct {
	Path string
	Method string
	Middles []gin.HandlerFunc
}

func (route *Route) GetPath() string {
	return route.Path
}

func (route *Route) GetMethod() string {
	return route.Method
}

func (route *Route) GetMiddles() []gin.HandlerFunc {
	return route.Middles
}
