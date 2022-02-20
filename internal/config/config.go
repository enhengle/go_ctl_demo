package config

import "github.com/tal-tech/go-zero/rest"

/**
自定义
配置文件
*/
type Config struct {
	rest.RestConf

	Log_path string
	Log_name string

	REDIS_HOST     string
	Redis_Port     int64
	Redis_Password string
	Redis_DB       int64

	Business_Host     string
	Business_Port     int64
	Business_Password string
	Business_Username string
	Business_DB       string
}
