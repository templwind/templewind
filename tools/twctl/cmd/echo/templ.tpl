package {{.pkgName}}

import (
	{{.templImports}}
)

templ {{.templName}}(c echo.Context, svcCtx *svc.ServiceContext){
    <div>
        <h1>{{.templName}}</h1>
    </div>
}