package main

import (
	"flag"
	"fmt"
	"greet/internal/config"
	"greet/internal/handler"
	"greet/internal/handler/cache"
	"greet/internal/svc"
	"greet/internal/xlogger"
	"greet/support/genv"

	"github.com/tal-tech/go-zero/core/conf"
	"github.com/tal-tech/go-zero/rest"
)

//var configFile = flag.String("greet", "etc/greet-api.yaml", "the config file") //goctl
var configFile = flag.String("develop", ".env.develop", "the config file") //配置文件启动程序

//读取配置文件启动程序
func main() {
	fmt.Println(configFile)
	flag.Parse()

	var c config.Config

	err := genv.EnvLoader(*configFile, &c)
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(c)

	//连接redis
	cache.LoadRedis(c)
	//连接db
	cache.LoadDB(c)

	ctx := svc.NewServiceContext(c)
	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	handler.RegisterHandlers(server, ctx)

	//测试日志打印
	xlogger.MainLogger.Debug("hello main Debug")
	xlogger.MainLogger.Info("hello main Info")
	xlogger.MainLogger.Warn("hello main warn")
	xlogger.GatewayLogger.Debug("hi gateway im Debug")
	xlogger.GatewayLogger.Info("hi gateway im Info")

	//程序启动
	server.Start()

}

//goctl启动程序
func testOne(c config.Config) {
	//将配置信息读取进Config结构体
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)
	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	handler.RegisterHandlers(server, ctx)

	/*xlogger.MainLogger.Debug("hello main Debug")
	xlogger.MainLogger.Info("hello main Info")
	xlogger.GatewayLogger.Debug("hi gateway im Debug")
	xlogger.GatewayLogger.Info("hi gateway im Info")
	redis.AddRedis(c)
	redis.AddRedis2(c)
	redis.AddRedisPool(c)

	db := sql.AddMysql(c.Business)
	sql.SelectAll(db)

	db2 := sql.AddMysql2(c.Business)
	sql.SelectAll2(db2)
	*/
	//http.HttpGet()
	//http.HttpParams()
	//http.HttpJson()
	//http.HttpFormDate()

	server.Start()

}
