package echo

import (
	_ "embed"
	"fmt"
	"sort"
	"strings"

	"github.com/templwind/templwind/tools/twctl/internal/types"
	"github.com/templwind/templwind/tools/twctl/internal/util"
	"github.com/templwind/templwind/tools/twctl/pkg/site/spec"

	"github.com/zeromicro/go-zero/core/collection"
	"github.com/zeromicro/go-zero/tools/goctl/config"
	"github.com/zeromicro/go-zero/tools/goctl/util/format"
)

//go:embed middleware.tpl
var middlewareImplementCode string

func genMiddleware(dir string, cfg *config.Config, site *spec.SiteSpec) error {
	middlewares := util.GetMiddleware(site)
	for _, item := range middlewares {
		middlewareFilename := strings.TrimSuffix(strings.ToLower(item), "middleware") + "_middleware"
		filename, err := format.FileNamingFormat(cfg.NamingFormat, middlewareFilename)
		if err != nil {
			return err
		}

		imports := genMiddlewareImports()
		// fmt.Println(imports)

		name := strings.TrimSuffix(item, "Middleware") + "Middleware"
		err = genFile(fileGenConfig{
			dir:             dir,
			subdir:          types.MiddlewareDir,
			filename:        filename + ".go",
			templateName:    "contextTemplate",
			category:        category,
			templateFile:    middlewareImplementCodeFile,
			builtinTemplate: middlewareImplementCode,
			data: map[string]string{
				"name":           util.ToTitle(name),
				"ImportPackages": imports,
			},
		})
		if err != nil {
			return err
		}
	}

	return nil
}

func genMiddlewareImports() string {
	importSet := collection.NewSet()
	// importSet.AddStr(fmt.Sprintf("\"%s\"", pathx.JoinPackages(parentPkg, types.ContextDir)))
	imports := importSet.KeysStr()
	sort.Strings(imports)
	projectSection := strings.Join(imports, "\n\t")
	depSection := `"github.com/labstack/echo/v4"`
	return fmt.Sprintf("%s\n\n\t%s", projectSection, depSection)
}
