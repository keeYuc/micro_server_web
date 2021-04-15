package main

import (
	"fmt"
	"keeyu/message"
	"net"
	"newusr/handler"

	"github.com/hashicorp/consul/api"
	"google.golang.org/grpc"
)

const (
	PORT = 18788
	IP   = "127.0.0.1"
	NAME = "newusr_"
)

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", IP, PORT))
	if err != nil {
		fmt.Println("listen:", err)
		return
	}
	defer lis.Close()
	// consul
	client, err := api.NewClient(api.DefaultConfig())
	if err != nil {
		fmt.Println("consul client:", err)
		return
	}
	asr := api.AgentServiceRegistration{
		// ID:      NAME,
		Tags:    []string{NAME},
		Name:    NAME,
		Address: IP,
		Port:    PORT,
		// Check: &api.AgentServiceCheck{
		//      CheckID:  "getimg service test",
		//      TCP:      "localhost:18787",
		//      Timeout:  "5s",
		//      Interval: "1s",
		// },
	}
	client.Agent().ServiceRegister(&asr)

	srv := grpc.NewServer()
	message.RegisterNewUsrServer(srv, new(handler.NewUsr))
	fmt.Println("newusr 服务启动成功")
	srv.Serve(lis)

}
