package echo

import (
	_ "embed"
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/templwind/templwind/tools/twctl/internal/types"
	"github.com/templwind/templwind/tools/twctl/internal/util"
	"github.com/templwind/templwind/tools/twctl/pkg/site/spec"

	"github.com/zeromicro/go-zero/tools/goctl/config"
	"github.com/zeromicro/go-zero/tools/goctl/util/format"
	"github.com/zeromicro/go-zero/tools/goctl/util/pathx"
)

const defaultControllerPackage = "controller"

//go:embed handler.tpl
var handlerTemplate string

func genHandlers(dir, rootPkg string, cfg *config.Config, site *spec.SiteSpec) error {
	for _, server := range site.Servers {
		for _, service := range server.Services {
			for _, handler := range service.Handlers {
				if err := genHandler(dir, rootPkg, cfg, server, handler); err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func genHandler(dir, rootPkg string, cfg *config.Config, server spec.Server, handler spec.Handler) error {
	handlerName := getHandlerName(handler, nil)
	handlerPath := getHandlerFolderPath(server)
	pkgName := handlerPath[strings.LastIndex(handlerPath, "/")+1:]
	// controllerName := defaultControllerPackage
	if handlerPath != types.HandlerDir {
		handlerName = strings.Title(handlerName)
		// controllerName = pkgName
	}

	controllerName := strings.ToLower(util.ToCamel(handler.Name))

	// fmt.Println("controllerName:", controllerName, handler.Name)
	// fmt.Println("handlerPath:", filepath.Join(handlerPath, util.ToKebab(handler.Name)))
	// os.Exit(0)

	filename, err := format.FileNamingFormat(cfg.NamingFormat, handlerName)
	if err != nil {
		return err
	}

	subDir := getHandlerFolderPath(server)
	handlerFile := path.Join(dir, subDir, filename+".go")
	os.Remove(handlerFile)

	type MethodConfig struct {
		RequestType    string
		ResponseType   string
		HasResp        bool
		HasReq         bool
		HandlerName    string
		HasDoc         bool
		Doc            string
		HasPage        bool
		ControllerName string
		ControllerType string
		Call           string
	}

	methods := []MethodConfig{}
	for _, method := range handler.Methods {
		hasResp := method.ResponseType != nil && len(method.ResponseType.GetName()) > 0
		hasReq := method.RequestType != nil && len(method.RequestType.GetName()) > 0

		requestType := ""
		if hasReq {
			requestType = util.ToTitle(method.RequestType.GetName())
		}
		responseType := ""
		if hasResp {
			responseType = util.ToTitle(method.ResponseType.GetName())
		}

		handlerName := util.ToTitle(getHandlerName(handler, &method))

		// fmt.Println("handlerName:", handlerName)
		methods = append(methods, MethodConfig{
			HandlerName:    handlerName,
			RequestType:    requestType,
			ResponseType:   responseType,
			HasResp:        hasResp,
			HasReq:         hasReq,
			HasDoc:         method.Doc != nil,
			HasPage:        method.Page != nil,
			ControllerName: controllerName,
			ControllerType: strings.Title(getControllerName(handler)),
			Call:           strings.Title(strings.TrimSuffix(handlerName, "Handler")),
		})
	}

	imports := genHandlerImports(server, handler, rootPkg)

	return genFile(fileGenConfig{
		dir:             dir,
		subdir:          subDir,
		filename:        filename + ".go",
		templateName:    "handlerTemplate",
		category:        category,
		templateFile:    handlerTemplateFile,
		builtinTemplate: handlerTemplate,
		data: map[string]any{
			"PkgName": pkgName,
			"Imports": imports,
			"Methods": methods,
		},
	})
}

func genHandlerImports(server spec.Server, handler spec.Handler, parentPkg string) string {
	hasTypes := false
	for _, method := range handler.Methods {
		if method.RequestType != nil && len(method.RequestType.GetName()) > 0 {
			hasTypes = true
			break
		}
	}

	imports := []string{
		fmt.Sprintf("\"%s\"\n\n", "net/http"),
		fmt.Sprintf("\"%s\"", pathx.JoinPackages(parentPkg, getControllerFolderPath(server, handler))),
		fmt.Sprintf("\"%s\"", pathx.JoinPackages(parentPkg, types.ContextDir)),
	}
	if hasTypes {
		imports = append(imports, fmt.Sprintf("\"%s\"", pathx.JoinPackages(parentPkg, types.TypesDir)))
	}

	imports = append(imports, fmt.Sprintf("\n\n\"%s\"", "github.com/labstack/echo/v4"))
	if hasTypes {
		imports = append(imports, fmt.Sprintf("\"%s\"", "github.com/zeromicro/go-zero/rest/httpx"))
	}

	return strings.Join(imports, "\n\t")
}

func getHandlerBaseName(route spec.Handler) (string, error) {
	name := route.Name
	name = strings.TrimSpace(name)
	name = strings.TrimSuffix(name, "handler")
	name = strings.TrimSuffix(name, "Handler")

	return name, nil
}

func getHandlerFolderPath(server spec.Server) string {
	folder := server.GetAnnotation(types.GroupProperty)
	if len(folder) == 0 || folder == "/" {
		return types.HandlerDir
	}

	folder = strings.TrimPrefix(folder, "/")
	folder = strings.TrimSuffix(folder, "/")

	return path.Join(types.HandlerDir, folder)
}

// getHandlerName constructs the handler name based on the handler and method details.
func getHandlerName(handler spec.Handler, method *spec.Method) string {
	baseName, err := getHandlerBaseName(handler)
	if err != nil {
		panic(err)
	}

	if method != nil {
		routePart := getRoutePart(handler, method)
		return baseName + strings.Title(strings.ToLower(method.Method)) + routePart + "Handler"
	}

	return baseName + "Handler"
}

// getRoutePart returns the sanitized part of the route for naming.
func getRoutePart(handler spec.Handler, method *spec.Method) string {
	baseRoute := handler.Methods[0].Route // Assuming the first method's route is the base route
	trimmedRoute := strings.TrimPrefix(method.Route, baseRoute)
	sanitizedRoute := sanitizeRoute(trimmedRoute)
	return sanitizedRoute
}

// sanitizeRoute converts the route to a title case format suitable for naming.
func sanitizeRoute(route string) string {
	// Remove leading and trailing slashes
	route = strings.Trim(route, "/")

	// Split the route by '/' and process each part
	parts := strings.Split(route, "/")
	for i, part := range parts {
		if part != "" {
			// Handle route parameters
			if strings.HasPrefix(part, ":") {
				parts[i] = "By" + strings.Title(strings.TrimPrefix(part, ":"))
			} else {
				parts[i] = strings.Title(part)
			}
		}
	}

	return strings.Join(parts, "")
}

func getControllerName(handler spec.Handler) string {
	baseName, err := getHandlerBaseName(handler)
	if err != nil {
		panic(err)
	}

	return baseName + "Controller"
}

func getPropsName(handler spec.Handler) string {
	baseName, err := getHandlerBaseName(handler)
	if err != nil {
		panic(err)
	}

	return baseName
}
