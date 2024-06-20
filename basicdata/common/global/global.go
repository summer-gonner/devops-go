package global

import (
	"devops-go/basicdata/config"
	"xorm.io/xorm"
)

var (
	Application *config.Application
	DB          *xorm.Engine
)
