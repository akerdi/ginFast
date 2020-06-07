package main

import (
	"ginFast/src"
	"github.com/shaohung001/ginFastApp"
	"log"
)

func main() {
	src.InitConfig()
	app := ginFastApp.New(src.ConfigData)
	_, err := app.Start()
	if err != nil {
		log.Fatalf("start app fail: %s", err)
	}
}
