package {{.PkgName}}

import (
	{{.Imports}}
)

{{range .Methods}}
{{if .HasDoc}}{{.Doc}}{{end}}
func {{.HandlerName}}(svcCtx *svc.ServiceContext) echo.HandlerFunc {
	return func(e echo.Context) error {
		{{if .HasReq}}var req types.{{.RequestType}}
		if err := httpx.Parse(e.Request(), &req); err != nil {
			// Log the error and send a generic error message to the client
			e.Logger().Error(err)
			// Send a JSON error response
			return e.JSON(http.StatusInternalServerError, map[string]interface{}{
				"error": "Internal Server Error",
			})
		}
		
		{{end}}l := {{.ControllerName}}.New{{.ControllerType}}(e.Request().Context(), svcCtx)
		{{if .HasResp}}resp, {{end}}err := l.{{.Call}}({{if .HasReq}}&req, {{end}}e)
		if err != nil {
			// Log the error and send a generic error message to the client
			e.Logger().Error(err)
			{{if .HasResp}}
			// Send a JSON error response
			return e.JSON(http.StatusInternalServerError, map[string]interface{}{
				"error": "Internal Server Error",
			}){{else}}
			// Send an HTML error response
			return e.HTML(http.StatusInternalServerError, "<h1>Internal Server Error</h1>"){{end}}
		}
		{{if .HasResp}}
		return e.JSON(http.StatusOK, resp){{else}}return nil{{end}}
	}
}
{{end}}
