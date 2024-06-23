package echo

import (
	_ "embed"
	"fmt"
	"path"
	"strconv"
	"strings"

	"github.com/templwind/templwind/tools/twctl/internal/types"
	"github.com/templwind/templwind/tools/twctl/internal/util"
	"github.com/templwind/templwind/tools/twctl/pkg/site/spec"

	"github.com/zeromicro/go-zero/tools/goctl/api/parser/g4/gen/api"
	"github.com/zeromicro/go-zero/tools/goctl/config"
	"github.com/zeromicro/go-zero/tools/goctl/util/format"
	"github.com/zeromicro/go-zero/tools/goctl/util/pathx"
	"github.com/zeromicro/go-zero/tools/goctl/vars"
)

//go:embed logic.tpl
var logicTemplate string

//go:embed templ.tpl
var templTemplate string

func genLogic(dir, rootPkg string, cfg *config.Config, site *spec.SiteSpec) error {
	for _, server := range site.Servers {
		for _, service := range server.Services {
			for _, handler := range service.Handlers {
				err := genLogicByHandler(dir, rootPkg, cfg, server, handler)
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func genLogicByHandler(dir, rootPkg string, cfg *config.Config, server spec.Server, handler spec.Handler) error {
	logic := getLogicName(handler)
	goFile, err := format.FileNamingFormat(cfg.NamingFormat, logic)
	if err != nil {
		return err
	}

	imports := genLogicImports(handler, rootPkg)
	var responseString string
	var returnString string
	var requestString string
	if handler.ResponseType != nil && len(handler.ResponseType.GetName()) > 0 {
		resp := util.ResponseGoTypeName(handler, types.TypesPacket)
		responseString = "(resp " + resp + ", err error)"
		returnString = "return"
	} else {
		responseString = "error"
		returnString = fmt.Sprintf("return templwind.Render(c, 200, %s(c, l.svcCtx))", strings.Title(handler.Name))
	}
	if handler.RequestType != nil && len(handler.RequestType.GetName()) > 0 {
		requestString = "req *" + util.RequestGoTypeName(handler, types.TypesPacket)
	}

	requestStringParts := []string{
		requestString,
		"c echo.Context",
	}
	requestString = func(parts []string) string {
		rParts := make([]string, 0)
		for _, part := range parts {
			if len(part) == 0 {
				continue
			}
			rParts = append(rParts, strings.TrimSpace(part))
		}
		return strings.Join(rParts, ", ")
	}(requestStringParts)

	subDir := getLogicFolderPath(server)

	templImports := genTemplImports(rootPkg)
	// fmt.Println("templImports", templImports)
	// templ file first
	genFile(fileGenConfig{
		dir:             dir,
		subdir:          subDir,
		filename:        strings.ToLower(handler.Name) + ".templ",
		templateName:    "templTemplate",
		category:        category,
		templateFile:    templTemplateFile,
		builtinTemplate: templTemplate,
		data: map[string]any{
			"pkgName":      subDir[strings.LastIndex(subDir, "/")+1:],
			"templImports": templImports,
			"templName":    strings.Title(handler.Name),
		},
	})

	return genFile(fileGenConfig{
		dir:             dir,
		subdir:          subDir,
		filename:        goFile + ".go",
		templateName:    "logicTemplate",
		category:        category,
		templateFile:    logicTemplateFile,
		builtinTemplate: logicTemplate,
		data: map[string]any{
			"pkgName":      subDir[strings.LastIndex(subDir, "/")+1:],
			"imports":      imports,
			"logic":        strings.Title(logic),
			"function":     strings.Title(strings.TrimSuffix(logic, "Logic")),
			"responseType": responseString,
			"returnString": returnString,
			"request":      requestString,
			"hasDoc":       len(handler.DocAnnotation.Properties) > 0,
			"doc":          util.GetDoc(handler.DocAnnotation.Properties),
		},
	})
}

func getLogicFolderPath(server spec.Server) string {
	folder := server.GetAnnotation(types.GroupProperty)
	if len(folder) == 0 || folder == "/" {
		return types.LogicDir
	}
	folder = strings.TrimPrefix(folder, "/")
	folder = strings.TrimSuffix(folder, "/")
	return path.Join(types.LogicDir, folder)
}

func genTemplImports(parentPkg string) string {
	var imports []string
	// imports = append(imports, `"net/http"`+"\n")
	imports = append(imports, fmt.Sprintf("\"%s\"", pathx.JoinPackages(parentPkg, types.ContextDir)))
	// if shallImportTypesPackage(handler) {
	// 	imports = append(imports, fmt.Sprintf("\"%s\"\n", pathx.JoinPackages(parentPkg, types.TypesDir)))
	// }
	imports = append(imports, "\n\t\"github.com/labstack/echo/v4\"")
	return strings.Join(imports, "\n\t")
}

func genLogicImports(handler spec.Handler, parentPkg string) string {
	var imports []string
	imports = append(imports, `"context"`)
	// imports = append(imports, `"net/http"`+"\n")
	imports = append(imports, fmt.Sprintf("\"%s\"", pathx.JoinPackages(parentPkg, types.ContextDir)))
	// if handler.ResponseType == nil {
	// 	imports = append(imports, fmt.Sprintf("\"%s\"", pathx.JoinPackages(parentPkg, types.UtilsDir)))
	// }
	if shallImportTypesPackage(handler) {
		imports = append(imports, fmt.Sprintf("\"%s\"\n", pathx.JoinPackages(parentPkg, types.TypesDir)))
	}
	imports = append(imports, "\n\n\"github.com/labstack/echo/v4\"")
	if handler.ResponseType == nil {
		imports = append(imports, "\"github.com/templwind/templwind\"")
	}
	imports = append(imports, fmt.Sprintf("\"%s/core/logx\"", vars.ProjectOpenSourceURL))
	return strings.Join(imports, "\n\t")
}

func onlyPrimitiveTypes(val string) bool {
	fields := strings.FieldsFunc(val, func(r rune) bool {
		return r == '[' || r == ']' || r == ' '
	})

	for _, field := range fields {
		if field == "map" {
			continue
		}
		// ignore array dimension number, like [5]int
		if _, err := strconv.Atoi(field); err == nil {
			continue
		}
		if !api.IsBasicType(field) {
			return false
		}
	}

	return true
}

func shallImportTypesPackage(handler spec.Handler) bool {

	if handler.RequestType != nil && len(handler.RequestType.GetName()) > 0 {
		return true
	}

	// fmt.Println("handler.RequestType.GetName()", handler.RequestType.GetName())

	respTypeName := handler.ResponseType
	if handler.ResponseType == nil || len(respTypeName.GetName()) == 0 {
		return false
	}

	if onlyPrimitiveTypes(respTypeName.GetName()) {
		return false
	}

	return true
}
