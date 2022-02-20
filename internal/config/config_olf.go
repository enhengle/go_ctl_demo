package config

import "github.com/tal-tech/go-zero/rest"

/**
goctl
配置文件
*/
type ConfigOlf struct {
	rest.RestConf
	Logger   Logger
	Redis    Redis
	Business Mysql
}

//日志
type Logger struct {
	Log_path string
	Log_name string
}

//redis
type Redis struct {
	Port     int
	Host     string
	Password string
	DB       int
}

type Mysql struct {
	Host     string
	Port     int
	Password string
	Username string
	DB       string
}
