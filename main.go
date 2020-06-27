package main

import (
	"ginFast/src/config"
	ginFastDB "ginFast/src/db"
	"ginFast/src/routes"
	"github.com/jinzhu/gorm"
	"github.com/shaohung001/ginFastApp"
	"log"
)

var App *ginFastApp.App
func main() {
	config.InitConfig()
	App = ginFastApp.New(config.ConfigData)
	applyRoutes(App)
	connectDB(App)
}

func connectDB(app *ginFastApp.App)  {
	app.ConnectDB(func(db *gorm.DB, err error) {
		if err != nil {
			panic(err)
		}
		err = ginFastDB.SetupTables(db)
		if err != nil {
			panic(err)
		}
		connectRedis(app)
	})
}

func connectRedis(app *ginFastApp.App)  {
	app.ConnectRedis(func(redisClient *ginFastApp.RedisClient, err error) {
		if err != nil {
			panic(err)
		}
		startApp(app)
	})
}

func startApp(app *ginFastApp.App)  {
	_, err := app.Start()
	if err != nil {
		log.Fatalf("start app fail: %s", err)
	}
}

func applyRoutes(app *ginFastApp.App) {
	for _, route := range routes.PublicRoutes {
		app.AddRoute(route)
	}
}