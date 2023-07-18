package {{.PkgName}}

import (
	"net/http"

    "github.com/go-playground/validator/v10"
	"github.com/zeromicro/go-zero/rest/httpx"
	{{.ImportPackages}}
)

func {{.HandlerName}}(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		{{if .HasRequest}}var req types.{{.RequestType}}
        _ = httpx.Parse(r, &req)
		if err := validator.New().Struct(req); err != nil {
			svcCtx.Response(w).Error(err)
			return
		}

		{{end}}l := {{.LogicName}}.New{{.LogicType}}(r.Context(), svcCtx)
		{{if .HasResp}}resp, {{end}}err := l.{{.Call}}({{if .HasRequest}}&req{{end}})
		if err != nil {
		    svcCtx.Response(w).Error(err)
		    return
		}

		{{if .HasResp}}svcCtx.Response(w).R(resp){{else}}svcCtx.Response(w).K(){{end}}
	}
}
