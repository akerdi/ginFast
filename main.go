package main

import (
	"fmt"
	"github.com/shaohung001/ginFastApp"
	"log"
)

func main() {
	fmt.Println("????????", ginFastApp.InitApp())
	
	fmt.Println("kkkkkk" , ginFastApp.AhuangName())
	err := ginFastApp.Start(ginFastApp.Init(), ":9300")
	if err != nil {
		log.Fatal("????", err)
	}
}
