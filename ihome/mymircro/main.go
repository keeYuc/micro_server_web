package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// import (
// 	"fmt"
// 	"mymicro/handler"
// 	pb "mymicro/proto"

// 	// "github.com/go-kit/kit/sd/consul"
// 	"github.com/hashicorp/consul/agent/consul"
// 	"github.com/micro/micro/v3/service"
// 	"github.com/micro/micro/v3/service/logger"
// )

// func main() {
// 	//consul
// 	config := consul.DefaultConfig()
// 	client, err := consul.NewClient(config, consul.Deps{})
// 	if err != nil {
// 		fmt.Println("错误", err)
// 		return
// 	}
// 	// Create service
// 	srv := service.New(
// 		service.Name("mymicro"),
// 		service.Version("latest"),
// 	)

// 	// Register handler
// 	pb.RegisterMymicroHandler(srv.Server(), new(handler.Mymicro))

// 	// Run service
// 	if err := srv.Run(); err != nil {
// 		logger.Fatal(err)
// 	}
// }

const (
	USERNAME = "root"
	PASSWORD = "Fucker123"
	NETWORK  = "tcp"
	SERVER   = "localhost"
	PORT     = 3306
	DATABASE = "test1"
)

func main() {
	dsn := fmt.Sprintf("%s:%s@%s(%s:%d)/%s", USERNAME, PASSWORD, NETWORK, SERVER, PORT, DATABASE)
	// dsn := fmt.Sprintf("%s:%s@%s(%s:%d)", USERNAME, PASSWORD, NETWORK, SERVER, PORT)
	fmt.Println(dsn)
	// return
	DB, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Printf("Open mysql failed,err:%v\n", err)
		return
	}
	DB.SetConnMaxLifetime(100 * time.Second) //最大连接周期，超过时间的连接就close
	DB.SetMaxOpenConns(100)                  //设置最大连接数
	DB.SetMaxIdleConns(16)                   //设置闲置连接数
}
