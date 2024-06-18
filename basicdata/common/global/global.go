package global

import (
	"devops-go/config"
	"xorm.io/xorm"
)

var (
	Application *config.Application
	DB          *xorm.Engine
)
