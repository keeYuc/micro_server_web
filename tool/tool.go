package tool

import (
	"database/sql"
	"fmt"

	"github.com/gomodule/redigo/redis"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	GlobalRedis redis.Pool
	GlobalMysql *sql.DB
	redis_init  bool
)

const (
	//mysql 配置
	MysqlName = "keeyu"
	MysqlPwd  = "Fucker123"
	MysqlAddr = "127.0.0.1"
	MysqlPort = "3306"
	MysqlDB   = "ihome"
	//redis 配置
	RedisAddr = "127.0.0.1:6379"
)

// root:Fucker123@tcp(127.0.0.1:3306)database?parseTime=true
//接口部分
func init() {
	init_redis()
	init_mysql()
}

func init_redis() {
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

func init_mysql() error {
	//打开数据库
	//拼接链接字符串
	connString := MysqlName + ":" + MysqlPwd + "@tcp(" + MysqlAddr + ":" + MysqlPort + ")/" + MysqlDB + "?charset=utf8mb4&parseTime=True&loc=Local"
	_db, err := gorm.Open(mysql.Open(connString), &gorm.Config{})
	if err != nil {
		fmt.Println("链接数据库失败", err)
		return err
	}
	_db.AutoMigrate(new(User))
	//连接池
	db, err := _db.DB()
	if err != nil {
		fmt.Println("获取数据库设置失败", err)
		return err
	}
	//连接池设置
	//设置初始化数据库链接个数
	db.SetMaxIdleConns(50)
	db.SetMaxOpenConns(70)
	db.SetConnMaxLifetime(60 * 30)
	GlobalMysql = db
	return nil
	// db.LogMode(true)
	// //默认情况下表名是复数
	// db.SingularTable(true)
	// //表的创建
	// return db.AutoMigrate(new(User),new(House),new(Area),new(Facility),new(HouseImage),new(OrderHouse)).Error
}

type User struct {
	ID            int    //用户编号
	Name          string `gorm:"size:32;unique"`  //用户名
	Password_hash string `gorm:"size:128" `       //用户密码加密的  hash
	Mobile        string `gorm:"size:11;unique" ` //手机号
	Real_name     string `gorm:"size:32" `        //真实姓名  实名认证
	Id_card       string `gorm:"size:20" `        //身份证号  实名认证
	Avatar_url    string `gorm:"size:256" `       //用户头像路径       通过fastdfs进行图片存储
	// Houses        []*House      //用户发布的房屋信息  一个人多套房
	// Orders        []*OrderHouse //用户下的订单       一个人多次订单
}
