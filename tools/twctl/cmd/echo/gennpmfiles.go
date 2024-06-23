package echo

import (
	_ "embed"
	"fmt"
	"path/filepath"

	"github.com/templwind/templwind/tools/twctl/internal/util"
	"github.com/templwind/templwind/tools/twctl/pkg/site/spec"
)

//go:embed readme.tpl
var readmeTemplate string

//go:embed package.tpl
var packageTemplate string

//go:embed postcss.tpl
var postcssConfigTemplate string

//go:embed tailwind.tpl
var tailwindConfigTemplate string

//go:embed tsconfig.tpl
var tsconfigTemplate string

//go:embed vite.tpl
var viteConfigTemplate string

//go:embed vite-env.tpl
var viteEnvTemplate string

//go:embed styles.tpl
var stylesTemplate string

//go:embed main.ts.tpl
var mainTsTemplate string

//go:embed error.ts.tpl
var errorTsTemplate string

//go:embed gitignore.tpl
var gitignoreTemplate string

var (
	srcDir       = "src"
	componentDir = filepath.Join(srcDir, "components")
)

func genNpmFiles(dir string, site *spec.SiteSpec) error {
	filename := util.ToCamel(site.Name)
	// fmt.Println("filename:", filename)

	files := []struct {
		filename        string
		subdir          string
		template        string
		templateName    string
		templateFile    string
		builtinTemplate string
		data            map[string]string
	}{
		{
			filename:        "README.md",
			templateName:    "readmeTemplate",
			templateFile:    readmeTemplateFile,
			builtinTemplate: readmeTemplate,
			data: map[string]string{
				"serviceName": filename,
			},
		},
		{
			filename:        "package.json",
			templateName:    "packageTemplate",
			templateFile:    packageTemplateFile,
			builtinTemplate: packageTemplate,
			data: map[string]string{
				"serviceName": filename,
			},
		},
		{
			filename:        "postcss.config.js",
			templateName:    "postcssConfigTemplate",
			templateFile:    postcssConfigTemplateFile,
			builtinTemplate: postcssConfigTemplate,
			data: map[string]string{
				"serviceName": filename,
			},
		},
		{
			filename:        "tailwind.config.js",
			templateName:    "tailwindConfigTemplate",
			templateFile:    tailwindConfigTemplateFile,
			builtinTemplate: tailwindConfigTemplate,
			data: map[string]string{
				"serviceName": filename,
			},
		},
		{
			filename:        "tsconfig.json",
			templateName:    "tsconfigTemplate",
			templateFile:    tsconfigTemplateFile,
			builtinTemplate: tsconfigTemplate,
			data: map[string]string{
				"serviceName": filename,
			},
		},
		{
			filename:        "vite.config.js",
			templateName:    "viteConfigTemplate",
			templateFile:    viteConfigTemplateFile,
			builtinTemplate: viteConfigTemplate,
			data: map[string]string{
				"serviceName": filename,
			},
		},
		{
			filename:        "vite-env.d.ts",
			subdir:          srcDir,
			templateName:    "viteEnvTemplate",
			templateFile:    viteEnvTemplateFile,
			builtinTemplate: viteEnvTemplate,
		},
		{
			filename:        "styles.scss",
			subdir:          srcDir,
			templateName:    "stylesTemplate",
			templateFile:    stylesTemplateFile,
			builtinTemplate: stylesTemplate,
		},
		{
			filename:        "main.ts",
			subdir:          srcDir,
			templateName:    "mainTsTemplate",
			templateFile:    mainTsTemplateFile,
			builtinTemplate: mainTsTemplate,
		},
		{
			filename:        "error.ts",
			subdir:          componentDir,
			templateName:    "errorTsTemplate",
			templateFile:    errorTsTemplateFile,
			builtinTemplate: errorTsTemplate,
		},
		{
			filename:        ".gitignore",
			templateName:    "gitignoreTemplate",
			templateFile:    gitignoreTemplateFile,
			builtinTemplate: gitignoreTemplate,
		},
	}

	for _, file := range files {
		if err := genFile(fileGenConfig{
			dir:             dir,
			subdir:          file.subdir,
			filename:        file.filename,
			templateName:    file.templateName,
			category:        category,
			templateFile:    file.templateFile,
			builtinTemplate: file.builtinTemplate,
			data: map[string]string{
				"serviceName": filename,
			},
		}); err != nil {
			fmt.Println(err.Error())
		}
	}

	return nil
}
