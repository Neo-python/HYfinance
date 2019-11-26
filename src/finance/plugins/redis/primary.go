package redis

import (
	"finance/plugins"
	"fmt"
	"github.com/gomodule/redigo/redis"
	"time"
)

var (
	// 定义常量
	RedisClient         *redis.Pool
	RedisHOST           string = plugins.Config.RedisHost
	RedisDB             int    = plugins.Config.RedisDB // 0
	RedisDefaultTimeOut int    = plugins.Config.RedisDefaultTimeOut
)

func init() {

	// 建立连接池
	RedisClient = &redis.Pool{
		MaxIdle:     3,
		MaxActive:   10,
		IdleTimeout: 180 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", RedisHOST)
			if err != nil {
				return nil, err
			}
			// 选择db
			c.Do("SELECT", RedisDB)
			return c, nil
		},
	}
}

func Set(key string, value interface{}, ex int) {
	if ex == 0 {
		ex = RedisDefaultTimeOut
	}
	redis_client := RedisClient.Get()
	defer redis_client.Close()

	redis_client.Do("SET", key, value, "EX", ex)
	fmt.Println(key, value)
}

func Get(key string) (interface{}, error) {
	redis_client := RedisClient.Get()
	defer redis_client.Close()
	return redis.String(redis_client.Do("GET", key))
}

func Delete(key string) {
	redis_client := RedisClient.Get()
	defer redis_client.Close()

	redis.String(redis_client.Do("DEL", key))
}
