package redis

import (
	"github.com/garyburd/redigo/redis"
	"greet/internal/config"
	"strconv"
)

var pool *redis.Pool

//连接线程池
func AddRedisPool(c config.ConfigOlf) {
	pool = &redis.Pool{
		Dial: func() (conn redis.Conn, e error) {
			hostAndPort := c.Redis.Host + ":" + strconv.Itoa(c.Redis.Port)
			return redis.Dial("tcp", hostAndPort)
		},
	}
}
