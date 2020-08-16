module ginFast

go 1.13

require (
	github.com/gin-gonic/gin v1.6.3
	github.com/go-playground/validator/v10 v10.3.0
	github.com/go-redis/redis v6.15.8+incompatible
	github.com/jinzhu/gorm v1.9.12
	github.com/json-iterator/go v1.1.9
	github.com/shaohung001/ginFastApp v0.0.1
)

replace github.com/shaohung001/ginFastApp => ../ginFastApp
