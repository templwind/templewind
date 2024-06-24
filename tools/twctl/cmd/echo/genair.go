package echo

import (
	_ "embed"

	"github.com/templwind/templwind/tools/twctl/internal/util"
	"github.com/templwind/templwind/tools/twctl/pkg/site/spec"
)

//go:embed air.tpl
var airTemplate string

func genAir(dir string, site *spec.SiteSpec) error {
	serviceName := util.ToCamel(site.Name)

	return genFile(fileGenConfig{
		dir:             dir,
		filename:        ".air.toml",
		templateName:    "airTemplate",
		category:        category,
		templateFile:    airTemplateFile,
		builtinTemplate: airTemplate,
		data: map[string]string{
			"serviceName": serviceName,
		},
	})
}
