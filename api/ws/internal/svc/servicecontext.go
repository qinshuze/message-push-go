package svc

import (
	"ccps.com/api/ws/internal/config"
	"ccps.com/api/ws/internal/lang"
	"ccps.com/api/ws/internal/middleware"
	"ccps.com/api/ws/internal/types"
	"ccps.com/internal/utils"
	"ccps.com/internal/utils/language"
	"ccps.com/internal/utils/response"
	"ccps.com/service/ws/wsclient"
	_ "github.com/lib/pq"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/rest/httpx"
	"github.com/zeromicro/go-zero/zrpc"
	"net/http"
)

type Middleware struct {
	Authorize rest.Middleware // 认证中间件
	Signature rest.Middleware // 接口签名中间件
}

type Rpc struct {
	Ws wsclient.Ws
}

type ServiceContext struct {
	Config     config.Config    // 全局配置
	LangLoader *language.Loader // 语言加载器
	AuthInfo   types.AuthInfo   // 登陆认证信息
	Db         sqlx.SqlConn     // 数据库实例
	Rpc                         // Rpc 服务实例
	Middleware                  // 注册中间件
}

func NewServiceContext(c config.Config) *ServiceContext {
	langLoader := language.Load(lang.AllLanguageMap)

	serviceContext := &ServiceContext{Config: c, LangLoader: langLoader}
	serviceContext.Db = sqlx.NewSqlConn("postgres", c.DataSource)
	serviceContext.Middleware.Signature = middleware.NewSignatureMiddleware(serviceContext.Db, langLoader).Handle
	serviceContext.Middleware.Authorize = middleware.NewAuthorizeMiddleware(
		c.Auth, langLoader, func(info types.AuthInfo) {
			serviceContext.AuthInfo = info
		},
	).Handle

	conn := zrpc.MustNewClient(zrpc.RpcClientConf{Etcd: c.WsRpc})
	serviceContext.Ws = wsclient.NewWs(conn)
	return serviceContext
}

func (c *ServiceContext) Response(w http.ResponseWriter) *response.JsonResponse {
	jsonResponse := response.NewJson(w)
	jsonResponse.SetLangLoader(c.LangLoader)
	jsonResponse.SetSendHandler(func(w http.ResponseWriter, data any, next func()) {
		httpx.WriteJson(w, utils.EmptyInt(jsonResponse.GetStatusCode(), http.StatusOK), data)
	})

	return jsonResponse
}
