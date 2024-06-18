package main

import (
	"devops-go/controllers"
	"log"
)

func main() {

	routes := controllers.InitRoutes()
	err := routes.Run(":8020")
	if err != nil {
		log.Fatalf("程序启动失败： %s", err.Error())
	}
}
