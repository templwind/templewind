package echo

import (
	"fmt"
	"os"
	"path"
	"sort"
	"strings"
	"text/template"
	"time"

	"github.com/templwind/templwind/tools/twctl/internal/types"
	"github.com/templwind/templwind/tools/twctl/internal/util"
	"github.com/templwind/templwind/tools/twctl/pkg/site/spec"

	"github.com/zeromicro/go-zero/core/collection"
	"github.com/zeromicro/go-zero/tools/goctl/config"
	"github.com/zeromicro/go-zero/tools/goctl/util/format"
	"github.com/zeromicro/go-zero/tools/goctl/util/pathx"
)

const (
	routesFilename = "routes"
	routesTemplate = `// Code generated by goctl. DO NOT EDIT.
package handler

import (
	{{if .hasTimeout}}
	"time"
	
	{{end}}

	{{.importPackages}}
)

type jwtCustomClaims struct {
	Name  string ` + "`json:\"name\"`" + `
	Admin bool   ` + "`json:\"admin\"`" + `
	jwt.RegisteredClaims
}

func RegisterHandlers(server *echo.Echo, svcCtx *svc.ServiceContext) {
	{{.routesAdditions}}
}
`
	routesAdditionTemplate = `
	{{.groupName}} := server.Group(
		"{{.prefix}}",{{if .middlewares}},
		[]echo.MiddlewareFunc{
			{{.middlewares}}
		}...,
		{{end}}
	)

	{{.routes}}
`
	timeoutThreshold = time.Millisecond
)

var mapping = map[string]string{
	"delete":  "DELETE",
	"get":     "GET",
	"head":    "HEAD",
	"post":    "POST",
	"put":     "PUT",
	"patch":   "PATCH",
	"connect": "CONNECT",
	"options": "OPTIONS",
	"trace":   "TRACE",
}

type (
	group struct {
		name             string
		routes           []route
		jwtEnabled       bool
		signatureEnabled bool
		authName         string
		timeout          string
		middlewares      []string
		prefix           string
		jwtTrans         string
		maxBytes         string
	}
	route struct {
		method  string
		path    string
		handler string
		doc     map[string]string
	}
)

func genRoutes(dir, rootPkg string, cfg *config.Config, site *spec.SiteSpec) error {
	var builder strings.Builder
	groups, err := getRoutes(site)
	if err != nil {
		return err
	}

	templateText, err := pathx.LoadTemplate(category, routesAdditionTemplateFile, routesAdditionTemplate)
	if err != nil {
		return err
	}

	var hasTimeout bool
	gt := template.Must(template.New("groupTemplate").Parse(templateText))
	for _, g := range groups {
		var routesBuilder strings.Builder
		for _, r := range g.routes {
			if len(r.doc) > 0 {
				routesBuilder.WriteString(fmt.Sprintf("\n%s\n", util.GetDoc(r.doc)))
			}
			routesBuilder.WriteString(fmt.Sprintf(
				`%s.%s("%s", %s)
	`,
				toPrefix(g.name)+"Group",
				mapping[strings.ToLower(r.method)],
				r.path,
				r.handler,
			))
		}

		for i, _ := range g.middlewares {
			g.middlewares[i] = "svcCtx." + util.ToTitle(g.middlewares[i]) + ","
		}

		if g.jwtEnabled {
			g.middlewares = append(g.middlewares, `			echojwt.WithConfig(echojwt.Config{
				NewClaimsFunc: func(c echo.Context) jwt.Claims { return new(jwtCustomClaims) },
				SigningKey: []byte(svcCtx.Config.`+g.authName+`.AccessSecret),
			}),`)
		}

		if err := gt.Execute(&builder, map[string]string{
			"groupName":   toPrefix(g.name) + "Group",
			"middlewares": strings.Join(g.middlewares, "\n"),
			"routes":      routesBuilder.String(),
			"prefix":      g.prefix,
		}); err != nil {
			return err
		}

		if len(g.timeout) > 0 {
			hasTimeout = true
		}
	}

	routeFilename, err := format.FileNamingFormat(cfg.NamingFormat, routesFilename)
	if err != nil {
		return err
	}

	routeFilename = routeFilename + ".go"
	filename := path.Join(dir, types.HandlerDir, routeFilename)
	os.Remove(filename)

	return genFile(fileGenConfig{
		dir:             dir,
		subdir:          types.HandlerDir,
		filename:        routeFilename,
		templateName:    "routesTemplate",
		category:        category,
		templateFile:    routesTemplateFile,
		builtinTemplate: routesTemplate,
		data: map[string]any{
			"hasTimeout":      hasTimeout,
			"importPackages":  genRouteImports(rootPkg, site),
			"routesAdditions": strings.TrimSpace(builder.String()),
		},
	})
}

func genRouteImports(parentPkg string, site *spec.SiteSpec) string {
	importSet := collection.NewSet()
	importSet.AddStr(fmt.Sprintf("\"%s\"", pathx.JoinPackages(parentPkg, types.ContextDir)))
	// importSet.AddStr(fmt.Sprintf("\"%s\"", pathx.JoinPackages(parentPkg, types.MiddlewareDir)))
	hasJwt := false
	for _, server := range site.Servers {
		folder := server.GetAnnotation(types.GroupProperty)
		importSet.AddStr(fmt.Sprintf("%s \"%s\"", toPrefix(folder),
			pathx.JoinPackages(parentPkg, types.HandlerDir, folder)))

		jwt := server.GetAnnotation("jwt")
		if len(jwt) > 0 {
			hasJwt = true
		}

	}
	imports := importSet.KeysStr()
	sort.Strings(imports)
	projectSection := strings.Join(imports, "\n\t")
	depSection := []string{`"github.com/golang-jwt/jwt/v5"`}
	if hasJwt {
		depSection = append(depSection, `"github.com/labstack/echo-jwt/v4"`)
	}
	depSection = append(depSection, `"github.com/labstack/echo/v4"`)
	return fmt.Sprintf("%s\n\n\t%s", projectSection, strings.Join(depSection, "\n\t"))
}

func getRoutes(site *spec.SiteSpec) ([]group, error) {
	var routes []group

	for _, server := range site.Servers {
		var groupedRoutes group
		folder := server.GetAnnotation(types.GroupProperty)
		groupedRoutes.name = folder
		for _, s := range server.Services {
			for _, r := range s.Handlers {
				handler := getHandlerName(r)
				handler = handler + "(svcCtx)"
				if len(folder) > 0 {
					handler = toPrefix(folder) + "." + strings.ToUpper(handler[:1]) + handler[1:]
				}

				groupedRoutes.routes = append(groupedRoutes.routes, route{
					method:  mapping[r.Method],
					path:    r.Route,
					handler: handler,
					doc:     r.DocAnnotation.Properties,
				})
			}
		}

		groupedRoutes.timeout = server.GetAnnotation("timeout")
		groupedRoutes.maxBytes = server.GetAnnotation("maxBytes")

		jwt := server.GetAnnotation("jwt")
		if len(jwt) > 0 {
			groupedRoutes.authName = jwt
			groupedRoutes.jwtEnabled = true
		}
		jwtTrans := server.GetAnnotation(types.JwtTransKey)
		if len(jwtTrans) > 0 {
			groupedRoutes.jwtTrans = jwtTrans
		}

		signature := server.GetAnnotation("signature")
		if signature == "true" {
			groupedRoutes.signatureEnabled = true
		}
		middleware := server.GetAnnotation("middleware")
		if len(middleware) > 0 {
			groupedRoutes.middlewares = append(groupedRoutes.middlewares, strings.Split(middleware, ",")...)
		}
		prefix := server.GetAnnotation("prefix")
		prefix = strings.ReplaceAll(prefix, `"`, "")
		prefix = strings.TrimSpace(prefix)
		if len(prefix) > 0 {
			prefix = path.Join("/", prefix)
			groupedRoutes.prefix = prefix
		}
		routes = append(routes, groupedRoutes)
	}

	return routes, nil
}

func toPrefix(folder string) string {
	return strings.ReplaceAll(folder, "/", "")
}