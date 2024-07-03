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

//go:embed templates/handler.tpl
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

	// fmt.Println("server", server)
	// os.Exit(0)

	filename, err := format.FileNamingFormat(cfg.NamingFormat, handlerName)
	if err != nil {
		return err
	}

	subDir := getHandlerFolderPath(server)
	handlerFile := path.Join(dir, subDir, filename+".go")
	os.Remove(handlerFile)

	methods := []MethodConfig{}
	for _, method := range handler.Methods {
		// fmt.Println("method:", method.Route)

		if method.IsStatic {
			continue
		}

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

		topicsFromClient := []Topic{}
		topicsFromServer := []Topic{}
		if method.IsSocket {

			for _, topic := range method.SocketNode.Topics {
				var reqType, resType string
				var hasReqType, hasResType bool
				if topic.RequestType != nil && len(topic.RequestType.GetName()) > 0 {
					hasReqType = true
					reqType = util.ToTitle(topic.RequestType.GetName())
				}
				if topic.ResponseType != nil && len(topic.ResponseType.GetName()) > 0 {
					hasResType = true
					resType = util.ToTitle(topic.ResponseType.GetName())
				}

				if !topic.InitiatedByClient {
					topicsFromServer = append(topicsFromServer, Topic{
						RawTopic:     strings.TrimSpace(topic.Topic),
						Topic:        "Topic" + util.ToPascal(topic.Topic),
						Name:         topic.GetName(),
						RequestType:  reqType,
						HasReqType:   hasReqType,
						ResponseType: resType,
						HasRespType:  hasResType,
						Call:         util.ToPascal(util.ToTitle(topic.Topic)),
					})
				} else {
					topicsFromClient = append(topicsFromClient, Topic{
						RawTopic:     strings.TrimSpace(topic.Topic),
						Topic:        "Topic" + util.ToPascal(topic.Topic),
						Name:         topic.GetName(),
						RequestType:  reqType,
						HasReqType:   hasReqType,
						ResponseType: resType,
						HasRespType:  hasResType,
						Call:         util.ToPascal(util.ToTitle(topic.Topic)),
					})
				}
			}
		}

		// fmt.Println("handlerName:", handlerName)
		methods = append(methods, MethodConfig{
			HandlerName:      handlerName,
			RequestType:      requestType,
			ResponseType:     responseType,
			HasResp:          hasResp,
			HasReq:           hasReq,
			HasDoc:           method.Doc != nil,
			HasPage:          method.Page != nil,
			ControllerName:   controllerName,
			ControllerType:   strings.Title(getControllerName(handler)),
			Call:             strings.Title(strings.TrimSuffix(handlerName, "Handler")),
			IsSocket:         method.IsSocket,
			TopicsFromClient: topicsFromClient,
			TopicsFromServer: topicsFromServer,
		})
	}

	// b, _ := json.MarshalIndent(methods, "", "  ")
	// fmt.Println("methods", string(b))

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
	theme := server.GetAnnotation("theme")
	if len(theme) == 0 {
		theme = "themes/templwind"
	} else {
		theme = "themes/" + theme
	}

	hasTypes := false
	hasTypesFromSocket := false
	requiresEvents := false
	for _, method := range handler.Methods {
		if method.RequestType != nil && len(method.RequestType.GetName()) > 0 {
			hasTypes = true
			break
		}
		if method.IsSocket {
			for _, topic := range method.SocketNode.Topics {
				if topic.RequestType != nil && len(topic.RequestType.GetName()) > 0 {
					hasTypesFromSocket = true
				}
				if topic.ResponseType != nil && len(topic.ResponseType.GetName()) > 0 {
					hasTypesFromSocket = true
				}
				if !topic.InitiatedByClient {
					requiresEvents = true
				}
			}
		}
	}

	hasSocket := false
	hasView := false
	for _, method := range handler.Methods {
		if method.IsSocket {
			hasSocket = true
			continue
		}
		if method.Method == "GET" || method.ReturnsPartial {
			hasView = true
			break
		}
	}

	imports := []string{}

	if hasSocket {
		imports = append(imports, fmt.Sprintf("\"%s\"", "context"))
		imports = append(imports, fmt.Sprintf("\"%s\"", "encoding/json"))
		imports = append(imports, fmt.Sprintf("\"%s\"", "log"))
	} else {
		imports = append(imports, fmt.Sprintf("\"%s\"", "net/http"))
	}

	if hasView {
		imports = append(imports, fmt.Sprintf("\"%s\"", "strconv"))
		imports = append(imports, fmt.Sprintf("\"%s\"", "time"))
	}

	imports = append(imports, "\n\n")

	if requiresEvents {
		imports = append(imports, fmt.Sprintf("\"%s\"", pathx.JoinPackages(parentPkg, types.EventsDir)))
	}
	imports = append(imports, fmt.Sprintf("\"%s\"", pathx.JoinPackages(parentPkg, getControllerFolderPath(server, handler))))
	imports = append(imports, fmt.Sprintf("\"%s\"", pathx.JoinPackages(parentPkg, types.ContextDir)))

	if hasTypes || hasTypesFromSocket {
		imports = append(imports, fmt.Sprintf("\"%s\"", pathx.JoinPackages(parentPkg, types.TypesDir)))
	}

	if hasView {
		imports = append(imports, fmt.Sprintf("\n\nbaseof \"%s\"", pathx.JoinPackages(parentPkg, theme, "layouts/baseof")))
		imports = append(imports, fmt.Sprintf("error500 \"%s\"", pathx.JoinPackages(parentPkg, theme, "error500")))
		imports = append(imports, fmt.Sprintf("footer \"%s\"", pathx.JoinPackages(parentPkg, theme, "partials", "footer")))
		imports = append(imports, fmt.Sprintf("head \"%s\"", pathx.JoinPackages(parentPkg, theme, "partials", "head")))
		imports = append(imports, fmt.Sprintf("header \"%s\"", pathx.JoinPackages(parentPkg, theme, "partials", "header")))
	}

	imports = append(imports, "\n\n")

	if hasSocket {
		imports = append(imports, fmt.Sprintf("gobwasWs \"%s\"", "github.com/gobwas/ws"))
		imports = append(imports, fmt.Sprintf("\"%s\"", "github.com/gobwas/ws/wsutil"))
	}

	imports = append(imports, fmt.Sprintf("\"%s\"", "github.com/labstack/echo/v4"))
	if hasTypes {
		imports = append(imports, fmt.Sprintf("\"%s\"", "github.com/zeromicro/go-zero/rest/httpx"))
	}

	if hasView {
		imports = append(imports, fmt.Sprintf("\"%s\"", "github.com/templwind/templwind"))
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
		routeName := getRouteName(handler, method)
		return baseName + strings.Title(strings.ToLower(method.Method)) + routeName + "Handler"
	}

	return baseName + "Handler"
}

// getRouteName returns the sanitized part of the route for naming.
func getRouteName(handler spec.Handler, method *spec.Method) string {
	baseRoute := handler.Methods[0].Route // Assuming the first method's route is the base route
	trimmedRoute := strings.TrimPrefix(method.Route, baseRoute)
	routeName := titleCaseRoute(trimmedRoute)

	// fmt.Println("RouteName", method.Route, baseRoute, trimmedRoute, routeName)

	return routeName
}

// titleCaseRoute converts the route to a title case format suitable for naming.
func titleCaseRoute(route string) string {
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
