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
		{{if .HasResp}}resp, {{else}}content, {{end}}err := l.{{.Call}}({{if .HasReq}}&req, {{end}}e)
		if err != nil {
			// Log the error and send a generic error message to the client
			e.Logger().Error(err)
			{{if .HasResp}}
			// Send a JSON error response
			return e.JSON(http.StatusInternalServerError, map[string]interface{}{
				"error": "Internal Server Error",
			}){{else}}
			// Send an HTML error response
			return templwind.Render(e, http.StatusInternalServerError,
				baseof.New(
					baseof.WithLTRDir("ltr"),
					baseof.WithLangCode("en"),
					baseof.WithHead(head.New(
						head.WithSiteTitle(svcCtx.Config.Site.Title),
						head.WithIsHome(true),
						head.WithCSS(
							svcCtx.Config.Assets.CSS...,
						),
					)),
					baseof.WithContent(error500.New(
						error500.WithErrors(
							"Internal Server Error",
						),
					)),
				),
			){{end}}
		}
		{{if .HasResp}}
		return e.JSON(http.StatusOK, resp){{else}}// Assemble the page
		return templwind.Render(e, http.StatusOK,
			baseof.New(
				baseof.WithLTRDir("ltr"),
				baseof.WithLangCode("en"),
				baseof.WithHead(head.New(
					head.WithSiteTitle(svcCtx.Config.Site.Title),
					head.WithIsHome(true),
					head.WithCSS(
						svcCtx.Config.Assets.CSS...,
					),
					head.WithJS(
						svcCtx.Config.Assets.JS...,
					),
				)),
				baseof.WithHeader(header.New(
					header.WithBrandName(svcCtx.Config.Site.Title),
					header.WithLoginURL("/auth/login"),
					header.WithLoginTitle("Log in"),
					header.WithMenus(svcCtx.Menus),
				)),
				baseof.WithFooter(footer.New(
					footer.WithYear(strconv.Itoa(time.Now().Year())),
				)),
				baseof.WithContent(content),
			),
		){{end}}
	}
}
{{end}}
