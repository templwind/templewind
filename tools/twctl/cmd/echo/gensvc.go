package echo

import (
	_ "embed"
	"fmt"
	"strings"

	"github.com/templwind/templwind/tools/twctl/internal/types"
	"github.com/templwind/templwind/tools/twctl/internal/util"
	"github.com/templwind/templwind/tools/twctl/pkg/site/spec"

	"github.com/zeromicro/go-zero/tools/goctl/config"
	"github.com/zeromicro/go-zero/tools/goctl/util/format"
	"github.com/zeromicro/go-zero/tools/goctl/util/pathx"
)

const contextFilename = "service_context"

//go:embed svc.tpl
var contextTemplate string

func genServiceContext(dir, rootPkg string, cfg *config.Config, site *spec.SiteSpec) error {
	filename, err := format.FileNamingFormat(cfg.NamingFormat, contextFilename)
	if err != nil {
		return err
	}

	var middlewareStr string
	var middlewareAssignment string
	middlewares := util.GetMiddleware(site)

	for _, item := range middlewares {
		middlewareStr += fmt.Sprintf("%s echo.MiddlewareFunc\n", item)
		name := strings.TrimSuffix(item, "Middleware") + "Middleware"
		middlewareAssignment += fmt.Sprintf("%s: %s,\n", item,
			fmt.Sprintf("middleware.New%s().%s", strings.Title(name), "Handle"))
	}

	imports := genSvcImports(rootPkg, len(middlewares) > 0)

	return genFile(fileGenConfig{
		dir:             dir,
		subdir:          types.ContextDir,
		filename:        filename + ".go",
		templateName:    "contextTemplate",
		category:        category,
		templateFile:    contextTemplateFile,
		builtinTemplate: contextTemplate,
		data: map[string]string{
			"imports":              imports,
			"config":               "config.Config",
			"middleware":           middlewareStr,
			"middlewareAssignment": middlewareAssignment,
		},
	})
}

func genSvcImports(rootPkg string, hasMiddlware bool) string {
	imports := []string{}
	imports = append(imports, fmt.Sprintf("\"%s\"", pathx.JoinPackages(rootPkg, types.ConfigDir)))
	if hasMiddlware {
		imports = append(imports, fmt.Sprintf("\"%s\"", pathx.JoinPackages(rootPkg, types.MiddlewareDir)))
	}

	imports = append(imports, "\n\n")
	imports = append(imports, fmt.Sprintf("\"%s\"", "github.com/jmoiron/sqlx"))
	imports = append(imports, fmt.Sprintf("\"%s/db\"", "github.com/templwind/templwind"))

	if hasMiddlware {
		imports = append(imports, fmt.Sprintf("\"%s\"", "github.com/labstack/echo/v4"))
	}

	return strings.Join(imports, "\n\t")
}
