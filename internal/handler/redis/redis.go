package redis

import (
	"github.com/go-redis/redis"
	"greet/internal/config"
	"strconv"
)

/**
redis操作
*/
//连接redis
func AddRedis(c config.ConfigOlf) (rcon *redis.Client) {
	hostAndPort := c.Redis.Host + ":" + strconv.Itoa(c.Redis.Port)
	client := redis.NewClient(&redis.Options{
		Addr:     hostAndPort,
		Password: c.Redis.Password,
		DB:       c.Redis.DB,
	})
	return client
}
