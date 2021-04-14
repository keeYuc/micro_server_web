package model

import (
	"github.com/gomodule/redigo/redis"
)

var GlobalRedis redis.Pool

func InitRedis() {
	GlobalRedis = redis.Pool{
		MaxIdle:     20,
		MaxActive:   50,
		IdleTimeout: 60 * 5,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", "127.0.0.1:6379")
		},
	}
}
