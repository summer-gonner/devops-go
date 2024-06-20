package svc

import (
	"devops-go/rpc/model/model/sysmodel"
	"devops-go/rpc/sys/internal/config"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config config.Config

	UserModel sysmodel.SysUserModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	sqlConn := sqlx.NewMysql(c.Mysql.Datasource)

	return &ServiceContext{
		Config: c,

		UserModel: sysmodel.NewSysUserModel(sqlConn),
	}
}
