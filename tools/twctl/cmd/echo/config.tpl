package config

import {{.imports}}

type Config struct {
	webserver.WebServerConf
	db.DBConfig
	{{.auth}}
	{{.jwtTrans}}
}
