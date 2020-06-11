package main

import (
	"ginFast/src/config"
	ginFastDB "ginFast/src/db"
	"ginFast/src/routes"
	"github.com/jinzhu/gorm"
	"github.com/shaohung001/ginFastApp"
	"log"
)

func main() {
	config.InitConfig()
	app := ginFastApp.New(config.ConfigData)
	applyRoutes(app)
	app.ConnectDB(func(db *gorm.DB, err error) {
		if err != nil {
			panic(err)
		}
		err = ginFastDB.SetupTables(db)
		if err != nil {
			panic(err)
		}
		
		_, err = app.Start()
		if err != nil {
			log.Fatalf("start app fail: %s", err)
		}
	})
	
}

func applyRoutes(app *ginFastApp.App) {
	for _, route := range routes.PublicRoutes {
		app.AddRoute(route)
	}
}