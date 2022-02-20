package cache

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"greet/internal/config"
)

type ConnGroup struct {
	RedisConnOne *Conn
}

type Conn struct {
	Pool *redis.Pool
}

var Connectors *ConnGroup

//初始化redis
func LoadRedis(c config.Config) {
	Connectors = &ConnGroup{
		RedisConnOne: getConn(c),
	}
}

//创建redis长连接
func getConn(c config.Config) *Conn {
	host := c.REDIS_HOST
	port := c.Redis_Port
	pwd := c.Redis_Password
	db := c.Redis_DB
	conn := new(Conn)
	addr := fmt.Sprintf("%s:%d", host, port)
	var err error
	conn.Pool, err = open(addr, pwd, db)
	if err != nil {
		panic("redis conn error")
	}
	return conn
}

func open(addr string, password string, db int64) (c *redis.Pool, err error) {
	c = &redis.Pool{
		Dial: func() (conn redis.Conn, e error) {
			c, err := redis.Dial("tcp", addr)
			if err != nil {
				return nil, err
			}
			if password != "" {
				if _, err := c.Do("AUTH", password); err != nil {
					_ = c.Close()
					return nil, err
				}
			}
			if db != 0 {
				if _, err := c.Do("SELECT", db); err != nil {
					_ = c.Close()
					return nil, err
				}
			}
			return c, err
		},
	}
	return c, nil
}
