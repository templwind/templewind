package {{.PkgName}}

import (
	"github.com/labstack/echo/v4"
	"github.com/zeromicro/go-zero/rest/httpx"
	{{.ImportPackages}}
)

{{if .HasDoc}}{{.Doc}}{{end}}
func {{.HandlerName}}(svcCtx *svc.ServiceContext) echo.HandlerFunc {
	return func(e echo.Context) error {
		{{if .HasRequest}}var req types.{{.RequestType}}
		if err := httpx.Parse(e.Request(), &req); err != nil {
			httpx.ErrorCtx(e.Request().Context(), e.Response(), err)
			return err
		}

		{{end}}l := {{.LogicName}}.New{{.LogicType}}(e.Request().Context(), svcCtx)
		{{if .HasResp}}resp, {{end}}err := l.{{.Call}}({{if .HasRequest}}&req, {{end}}e)
		if err != nil {
			httpx.ErrorCtx(e.Request().Context(), e.Response(), err)
		} else {
			{{if .HasResp}}httpx.OkJsonCtx(e.Request().Context(), e.Response(), resp){{else}}httpx.Ok(e.Response()){{end}}
		}
		return nil
	}
}

