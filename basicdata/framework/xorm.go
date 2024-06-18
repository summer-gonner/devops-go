package framework

import (
	"devops-go/common/global"
	_ "github.com/go-sql-driver/mysql" //必须导入这个包
	"log"
	"xorm.io/xorm"
	log2 "xorm.io/xorm/log"
)

func InitXorm() *xorm.Engine {
	if global.Application.DataSource.Host == "" {
		log.Fatalf("未配置数据库地址")
	}
	dsn := global.Application.DataSource.Dsn()
	log.Printf("查看数据库连接地址：%s", dsn)
	engine, err := xorm.NewEngine("mysql", dsn)
	engine.SetMaxIdleConns(10)
	engine.SetMaxOpenConns(20)
	engine.ShowSQL(true)
	engine.Logger().SetLevel(log2.LOG_DEBUG)
	if err != nil {
		log.Fatalf("数据库连接失败： %s", err)
	}
	if engine == nil {
		log.Fatalf("数据库连接失败")
	}
	return engine
}
