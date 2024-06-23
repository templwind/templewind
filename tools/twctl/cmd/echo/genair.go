package echo

import (
	_ "embed"
)

//go:embed air.tpl
var airTemplate string

func genAir(dir string) error {
	return genFile(fileGenConfig{
		dir:             dir,
		filename:        ".air.toml",
		templateName:    "airTemplate",
		category:        category,
		templateFile:    airTemplateFile,
		builtinTemplate: airTemplate,
		data:            map[string]string{},
	})
}
