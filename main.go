package main

import (
	"devops-go/common/global"
	"devops-go/controllers"
	"devops-go/framework"
	"log"
)

func main() {
	framework.InitApplicationYaml()
	global.DB = framework.InitXorm() //初始化数据库
	routes := controllers.InitRoutes()
	addr := global.Application.Server.Addr()
	err := routes.Run(addr)
	if err != nil {
		log.Fatalf("程序启动失败： %s", err.Error())
	}
}
