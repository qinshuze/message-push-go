package svc

import (
	"ccps.com/api/admin/internal/config"
	"ccps.com/api/admin/internal/lang"
	"ccps.com/api/admin/internal/middleware"
	"ccps.com/api/admin/internal/types"
	"ccps.com/internal/utils"
	"ccps.com/internal/utils/language"
	"ccps.com/internal/utils/response"
	"ccps.com/service/ws/wsclient"
	"github.com/fatih/color"
	_ "github.com/lib/pq"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/rest/httpx"
	"github.com/zeromicro/go-zero/zrpc"
	"net/http"
	"runtime/debug"
)

type Middleware struct {
	Authorize rest.Middleware // 认证中间件
	Signature rest.Middleware // 接口签名中间件
}

type Rpc struct {
	Ws wsclient.Ws
}

type ServiceContext struct {
	Config     config.Config
	LangLoader *language.Loader // 语言加载器
	Db         sqlx.SqlConn     // 数据库实例
	Rpc                         // Rpc 服务实例
	Middleware                  // 注册中间件
	LoginUser  types.LoginUser  // 登陆用户
}

func NewServiceContext(c config.Config) *ServiceContext {
	langLoader := language.Load(lang.AllLanguageMap)
	conn := zrpc.MustNewClient(zrpc.RpcClientConf{Etcd: c.WsRpc})

	serviceContext := &ServiceContext{Config: c, LangLoader: langLoader}
	serviceContext.Db = sqlx.NewSqlConn("postgres", c.DataSource)
	serviceContext.Ws = wsclient.NewWs(conn)

	serviceContext.Middleware.Signature = middleware.NewSignatureMiddleware(serviceContext.Db, langLoader).Handle
	serviceContext.Middleware.Authorize = middleware.NewAuthorizeMiddleware(
		c.Auth, langLoader, func(info types.LoginUser) {
			serviceContext.LoginUser = info
		},
	).Handle

	return serviceContext
}

func (c *ServiceContext) Response(w http.ResponseWriter) *response.JsonResponse {
	jsonResponse := response.NewJson(w)
	jsonResponse.SetLangLoader(c.LangLoader)
	jsonResponse.SetErrorHandler(func(w http.ResponseWriter, err error, next func()) {
		color.Red(err.Error())
		color.Red(string(debug.Stack()))
		next()
	})
	jsonResponse.SetSendHandler(func(w http.ResponseWriter, data any, next func()) {
		statusCode := utils.EmptyInt(jsonResponse.GetStatusCode(), http.StatusOK)
		httpx.WriteJson(w, statusCode, data)
	})

	return jsonResponse
}
