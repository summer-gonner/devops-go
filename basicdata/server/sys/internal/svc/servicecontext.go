package svc

import (
	"devops-go/basicdata/server/model/model/sysmodel"
	"devops-go/basicdata/server/sys/internal/config"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config      config.Config
	RedisClient *redis.Redis
	UserModel   sysmodel.SysUserModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	sqlConn := sqlx.NewMysql(c.Mysql.Datasource)

	conf := redis.RedisConf{
		Host: c.RedisConf.Host,
		Type: c.RedisConf.Type,
		Pass: c.RedisConf.Pass,
		Tls:  c.RedisConf.Tls,
	}

	return &ServiceContext{
		Config:      c,
		RedisClient: redis.MustNewRedis(conf),
		UserModel:   sysmodel.NewSysUserModel(sqlConn),
	}
}
