package config

import (
	"github.com/zeromicro/go-zero/core/discov"
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	Lang string
	rest.RestConf
	Auth       AuthConf
	DataSource string
	WsRpc      discov.EtcdConf
}
