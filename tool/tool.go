package tool

import (
	"github.com/gomodule/redigo/redis"
)

var (
	GlobalRedis redis.Pool
	redis_init  bool
)

const (
	//mysql 配置
	MysqlName = "root"
	MysqlPwd  = "Fucker123"
	MysqlAddr = "127.0.0.1"
	MysqlPort = "3306"
	MysqlDB   = "IhomeData"
	//redis 配置
	RedisAddr = "127.0.0.1:6379"
)

//接口部分
func init() {
	if !redis_init {
		GlobalRedis = redis.Pool{
			MaxIdle:     20,
			MaxActive:   50,
			IdleTimeout: 60 * 5,
			Dial: func() (redis.Conn, error) {
				return redis.Dial("tcp", RedisAddr)
			},
		}
		redis_init = true
	}
}
