package config

import "github.com/zeromicro/go-zero/zrpc"

type WebSocketConf struct {
	ListenOn   string
	ClientAuth string
	WhiteList  []string
}

type Config struct {
	WebSocket WebSocketConf
	zrpc.RpcServerConf
}
