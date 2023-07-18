package config

import (
	"github.com/zeromicro/go-zero/core/discov"
	"github.com/zeromicro/go-zero/rest"
)

type AuthConf struct {
	AccessSecret string
	AccessTTL    int
}

type Config struct {
	Lang string
	rest.RestConf
	Auth       AuthConf
	DataSource string
	WsRpc      discov.EtcdConf
}
