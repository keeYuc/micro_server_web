package tool

import (
	"github.com/gomodule/redigo/redis"
)

var GlobalRedis redis.Pool

//mysql 配置

const MysqlName = "root"
const MysqlPwd = "Fucker123"
const MysqlAddr = "127.0.0.1"
const MysqlPort = "3306"
const MysqlDB = "fuckhome"

//redis 配置
const RedisAddr = "127.0.0.1:6379"

//接口部分
func InitRedis() {
	GlobalRedis = redis.Pool{
		MaxIdle:     20,
		MaxActive:   50,
		IdleTimeout: 60 * 5,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", RedisAddr)
		},
	}
}
