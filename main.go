package main

import (
	"fmt"
	"ginFast/src/config"
	ginFastDB "ginFast/src/db"
	"ginFast/src/db/entity/email"
	"ginFast/src/routes"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/jinzhu/gorm"
	"github.com/shaohung001/ginFastApp"
	"gopkg.in/go-playground/validator.v8"
)

var App *ginFastApp.App

func main() {
	config.InitConfig()
	App = ginFastApp.New(config.ConfigData)
	applyRoutes(App)
	connectDB(App)
}

func connectDB(app *ginFastApp.App) {
	app.ConnectDB(func(db *gorm.DB, err error) {
		if err != nil {
			panic(err)
		}
		err = ginFastDB.SetupTables(db)
		if err != nil {
			panic(err)
		}

		engine, err := app.Start()
		bindValidator()
		startEngine(engine)
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
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		err := v.RegisterValidation("EmailValid", email.EmailValid)
		if err != nil {
			log.Println("bindValidator err: ", err)
		}
	}
}
