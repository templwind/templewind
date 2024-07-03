package config

import (
	{{.imports}}
)

type Config struct {
	webserver.WebServerConf
	db.DBConfig
	{{.auth}}
	{{.jwtTrans}}
	Site struct {
		Title string
	}
	Assets Assets
	Menus  Menus
}

type Menus map[string][]MenuEntry
type MenuEntry struct {
	URL        string
	Title      string
	Identifier string
	Children   []MenuEntry `yaml:"Children,omitempty"`
}

type Assets struct {
	CSS []string
	JS  []string
}