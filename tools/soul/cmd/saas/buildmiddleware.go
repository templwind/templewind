package saas

import (
	_ "embed"
	"fmt"
	"path/filepath"
	"strings"

	"github.com/templwind/templwind/tools/soul/internal/imports"
	"github.com/templwind/templwind/tools/soul/internal/types"
	"github.com/templwind/templwind/tools/soul/internal/util"
)

func buildMiddleware(builder *SaaSBuilder) error {
	middlewares := util.GetMiddleware(builder.Spec)
	for _, item := range middlewares {

		noCache := false
		if strings.EqualFold(item, "nocache") {
			noCache = true
		}

		middlewareFilename := strings.TrimSuffix(strings.ToLower(item), "middleware")
		// fmt.Println("generating middleware:", middlewareFilename)

		builder.WithRenameFile(
			filepath.Join(
				types.MiddlewareDir,
				"template"),
			filepath.Join(
				types.MiddlewareDir,
				middlewareFilename+".go",
			))

		builder.Data["name"] = util.ToTitle(strings.TrimSuffix(item, "Middleware") + "Middleware")
		builder.Data["imports"] = imports.New(
			imports.WithImport("github.com/labstack/echo/v4"),
		).String()
		builder.Data["isNoCache"] = noCache

		err := builder.genFile(fileGenConfig{
			subdir: types.MiddlewareDir,
			templateFile: filepath.Join(
				"templates",
				types.MiddlewareDir,
				"template.tpl",
			),
			data: builder.Data,
		})
		if err != nil {
			fmt.Println("gen middleware failed:", err)
		}
	}

	return nil
}
