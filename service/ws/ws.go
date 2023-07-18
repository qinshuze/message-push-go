package main

import (
	"ccps.com/service/ws/internal/config"
	"ccps.com/service/ws/internal/message"
	"ccps.com/service/ws/internal/server"
	"ccps.com/service/ws/internal/svc"
	"ccps.com/service/ws/ws"
	"flag"
	"github.com/fatih/color"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net/http"
)

var configFile = flag.String("f", "etc/ws.yaml", "the config file")

func startWebSocketServer(ctx *svc.ServiceContext) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		message.NewClient(ctx, r, w)
	})
	color.Green("Starting websocket server at %s...\n", ctx.Config.WebSocket.ListenOn)
	log.Fatal(http.ListenAndServe(ctx.Config.WebSocket.ListenOn, nil))
}

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	go func() {
		s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
			ws.RegisterWsServer(grpcServer, server.NewWsServer(ctx))
			if c.Mode == service.DevMode || c.Mode == service.TestMode {
				reflection.Register(grpcServer)
			}
		})
		defer s.Stop()

		color.Green("Starting rpc server at %s...\n", c.ListenOn)
		s.Start()
	}()

	startWebSocketServer(ctx)
}
