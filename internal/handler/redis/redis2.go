package redis

import (
	"github.com/garyburd/redigo/redis"
	"greet/internal/config"
	"greet/internal/xlogger"
	"strconv"
)

//连接redis
func AddRedis2(c config.ConfigOlf) (rcon redis.Conn) {
	hostAndPort := c.Redis.Host + ":" + strconv.Itoa(c.Redis.Port)
	rcon, err := redis.Dial("tcp", hostAndPort)
	if err != nil {
		xlogger.MainLogger.Debug("redis连接失败")
	}
	rcon.Do("select", c.Redis.DB)
	return rcon

}
