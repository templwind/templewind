package echo

import (
	"fmt"

	"github.com/zeromicro/go-zero/tools/goctl/util/pathx"
)

const (
	category                    = "api"
	configTemplateFile          = "config.tpl"
	contextTemplateFile         = "context.tpl"
	etcTemplateFile             = "etc.tpl"
	handlerTemplateFile         = "handler.tpl"
	controllerTemplateFile      = "controller.tpl"
	mainTemplateFile            = "main.tpl"
	middlewareImplementCodeFile = "middleware.tpl"
	routesTemplateFile          = "routes.tpl"
	routesAdditionTemplateFile  = "route-addition.tpl"
	templTemplateFile           = "templ.tpl"
	typesTemplateFile           = "types.tpl"
	airTemplateFile             = "air.tpl"
	packageTemplateFile         = "package.tpl"
	readmeTemplateFile          = "README.md"
	postcssConfigTemplateFile   = "postcss.tpl"
	tailwindConfigTemplateFile  = "tailwind.tpl"
	tsconfigTemplateFile        = "tsconfig.tpl"
	viteConfigTemplateFile      = "vite.tpl"
	viteEnvTemplateFile         = "vite-env.tpl"
	stylesTemplateFile          = "styles.tpl"
	mainTsTemplateFile          = "main.ts.tpl"
	errorTsTemplateFile         = "error.ts.tpl"
	gitignoreTemplateFile       = "gitignore.tpl"
	schemaTemplateFile          = "1_schema.tpl"
	localEnvTemplateFile        = "local.env.tpl"
	makefileTemplateFile        = "makefile.tpl"
)

var templates = map[string]string{
	configTemplateFile:          configTemplate,
	contextTemplateFile:         contextTemplate,
	etcTemplateFile:             etcTemplate,
	handlerTemplateFile:         handlerTemplate,
	controllerTemplateFile:      controllerTemplate,
	mainTemplateFile:            mainTemplate,
	middlewareImplementCodeFile: middlewareImplementCode,
	routesTemplateFile:          routesTemplate,
	routesAdditionTemplateFile:  routesAdditionTemplate,
	templTemplateFile:           templTemplate,
	typesTemplateFile:           typesTemplate,
	airTemplateFile:             airTemplate,
	packageTemplateFile:         packageTemplate,
	readmeTemplateFile:          readmeTemplate,
	postcssConfigTemplateFile:   postcssConfigTemplate,
	tailwindConfigTemplateFile:  tailwindConfigTemplate,
	tsconfigTemplateFile:        tsconfigTemplate,
	viteConfigTemplateFile:      viteConfigTemplate,
	viteEnvTemplateFile:         viteEnvTemplate,
	stylesTemplateFile:          stylesTemplate,
	mainTsTemplateFile:          mainTsTemplate,
	errorTsTemplateFile:         errorTsTemplate,
	gitignoreTemplateFile:       gitignoreTemplate,
	schemaTemplateFile:          schemaTemplate,
	localEnvTemplateFile:        localEnvTemplate,
	makefileTemplateFile:        makefileTemplate,
}

// Category returns the category of the api files.
func Category() string {
	return category
}

// Clean cleans the generated deployment files.
func Clean() error {
	return pathx.Clean(category)
}

// GenTemplates generates api template files.
func GenTemplates() error {
	return pathx.InitTemplates(category, templates)
}

// RevertTemplate reverts the given template file to the default value.
func RevertTemplate(name string) error {
	content, ok := templates[name]
	if !ok {
		return fmt.Errorf("%s: no such file name", name)
	}
	return pathx.CreateTemplate(category, name, content)
}

// Update updates the template files to the templates built in current goctl.
func Update() error {
	err := Clean()
	if err != nil {
		return err
	}

	return pathx.InitTemplates(category, templates)
}
