package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	SysRpc zrpc.RpcClientConf

	Mysql struct {
		Datasource string
	}
}