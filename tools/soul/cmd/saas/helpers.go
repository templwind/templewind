package saas

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/templwind/templwind/tools/soul/internal/util"
	"github.com/templwind/templwind/tools/soul/pkg/site/spec"
)

// getHandlerName constructs the handler name based on the handler and method details.
func getHandlerName(handler spec.Handler, method *spec.Method) string {
	baseName, err := getHandlerBaseName(handler)
	if err != nil {
		panic(err)
	}

	if method != nil {
		routeName := getRouteName(handler, method)
		// return baseName + strings.Title(strings.ToLower(method.Method)) + routeName + "Handler"
		return strings.Title(strings.ToLower(method.Method)) + routeName + "Handler"
	}

	return util.ToPascal(baseName + "Handler")
}

// getRouteName returns the sanitized part of the route for naming.
func getRouteName(handler spec.Handler, method *spec.Method) string {
	baseRoute := handler.Methods[0].Route // Assuming the first method's route is the base route
	trimmedRoute := strings.TrimPrefix(method.Route, baseRoute)
	routeName := titleCaseRoute(trimmedRoute)

	// fmt.Println("RouteName", method.Route, baseRoute, trimmedRoute, routeName)

	return routeName
}

func getHandlerBaseName(route spec.Handler) (string, error) {
	name := route.Name
	name = strings.TrimSpace(name)
	name = strings.TrimSuffix(name, "handler")
	name = strings.TrimSuffix(name, "Handler")

	return name, nil
}

func getLogicName(handler spec.Handler) string {
	baseName, err := getHandlerBaseName(handler)
	if err != nil {
		panic(err)
	}

	return baseName + "Logic"
}

func getPropsName(handler spec.Handler) string {
	baseName, err := getHandlerBaseName(handler)
	if err != nil {
		panic(err)
	}

	return baseName
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

func loadFile(file string) (string, error) {
	if !fileExists(file) {
		return "", nil
	}

	content, err := os.ReadFile(file)
	if err != nil {
		return "", err
	}

	return string(content), nil
}

// FileExists returns true if the specified file is exists.
func fileExists(file string) bool {
	_, err := os.Stat(file)
	return err == nil
}

// FileNameWithoutExt returns a file name without suffix.
func fileNameWithoutExt(file string) string {
	return strings.TrimSuffix(file, filepath.Ext(file))
}
