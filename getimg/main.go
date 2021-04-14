package main

import (
	"fmt"
	"getimg/handler"
	"keeyu/message"
	"net"

	"github.com/hashicorp/consul/api"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":18787")
	if err != nil {
		fmt.Println("listen:", err)
		return
	}
	//consul
	client, err := api.NewClient(api.DefaultConfig())
	if err != nil {
		fmt.Println("consul client:", err)
		return
	}
	asr := api.AgentServiceRegistration{
		ID:      "getimg_",
		Tags:    []string{"getimg_"},
		Name:    "getimg_",
		Address: "127.0.0.1",
		Port:    18787,
		// Check: &api.AgentServiceCheck{
		// 	CheckID:  "getimg service test",
		// 	TCP:      "localhost:18787",
		// 	Timeout:  "5s",
		// 	Interval: "1s",
		// },
	}
	client.Agent().ServiceRegister(&asr)

	//consul
	defer lis.Close()
	srv := grpc.NewServer()
	message.RegisterGetimgServer(srv, new(handler.Getimg))
	fmt.Println("getimg 服务启动成功")
	srv.Serve(lis)
}
