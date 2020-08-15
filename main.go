package main

import (
	"fmt"
	"ginFast/src/config"
	ginFastDB "ginFast/src/db"
	"ginFast/src/routes"
	customValidate "ginFast/src/routes/validate"
	"github.com/gin-gonic/gin/binding"
	"log"
	"regexp"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/shaohung001/ginFastApp"
)

var App *ginFastApp.App

func main() {
	testFunc()
	config.InitConfig()
	App = ginFastApp.New(config.ConfigData)
	applyRoutes(App)
	ch := make(chan error, 1)
	go connectDBGenerator(App, ch)
	err := <-ch
	if err != nil {
		panic(err)
	}
	ch = make(chan error, 1)
	go connectRedisGenerator(App, ch)
	<- ch
	
	engine, err := App.Start()
	bindValidator()
	startEngine(engine)
}

func testFunc() {
	b := []byte("7683242@163.com")
	pat := `@qq.com$|@163.com`
	reg1 := regexp.MustCompile(pat)
	fmt.Printf("regexp:  %v \n\n", reg1.Match(b))
}

func connectDBGenerator(app *ginFastApp.App, ch chan<- error)  {
	app.ConnectDB(func(db *gorm.DB, err error) {
		if err != nil {
			ch<- err
			return
		}
		err = ginFastDB.SetupTables(db)
		if err != nil {
			ch<- err
			return
		}
		ch<- nil
	})
}
func connectRedisGenerator(app *ginFastApp.App, ch chan<- error) {
	app.ConnectRedis(func(redisClient *ginFastApp.RedisClient, err error) {
		ginFastDB.SetupRedis(redisClient)
		ch<- nil
	})
}

func startEngine(engine *gin.Engine) {
	port := config.ConfigData.Port
	portStr := fmt.Sprintf(":%d", port)
	err := engine.Run(portStr)
	if err != nil {
		log.Fatalf("server run port: %d error: %s", port, err)
	}
	fmt.Println("server is starting now! port : ", port)
}

func applyRoutes(app *ginFastApp.App) {
	for _, route := range routes.PublicRoutes {
		app.AddRoute(route)
	}
}

func bindValidator() {
	binding.Validator = new(customValidate.DefaultValidator)
	//if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
	//	err := v.RegisterValidation("EmailValid", email.EmailValid)
	//	if err != nil {
	//		log.Println("bindValidator err: ", err)
	//	}
	//}
}
