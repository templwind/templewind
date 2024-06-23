package config

import {{.imports}}

type Config struct {
	webserver.WebServerConf
	{{.auth}}
	{{.jwtTrans}}
}
