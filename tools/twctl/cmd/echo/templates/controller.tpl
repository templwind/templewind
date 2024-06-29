package {{.pkgName}}

import (
	{{.imports}}
)

type {{.controllerType}} struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func New{{.controllerType}}(ctx context.Context, svcCtx *svc.ServiceContext) *{{.controllerType}} {
	return &{{.controllerType}}{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

{{- range .methods}}
{{if .HasDoc}}{{.Doc}}{{end}}
func (l *{{.ControllerType}}) {{.Call}}({{.Request}}) {{.ResponseType}} {
	// todo: add your logic here and delete this line

	{{.ReturnString}}
}
{{end}}