package main

import (
	"fmt"
	"keeyu/message"
	"net"
	"newusr/handler"

	"github.com/hashicorp/consul/api"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":18788")
	if err != nil {
		fmt.Println("listen:", err)
		return
	}
	// consul
	client, err := api.NewClient(api.DefaultConfig())
	if err != nil {
		fmt.Println("consul client:", err)
		return
	}
	asr := api.AgentServiceRegistration{
		ID:      "newusr_",
		Tags:    []string{"newusr_"},
		Name:    "newusr_",
		Address: "127.0.0.1",
		Port:    18788,
		// Check: &api.AgentServiceCheck{
		//      CheckID:  "getimg service test",
		//      TCP:      "localhost:18787",
		//      Timeout:  "5s",
		//      Interval: "1s",
		// },
	}
	client.Agent().ServiceRegister(&asr)

	defer lis.Close()
	srv := grpc.NewServer()
	message.RegisterNewUsrServer(srv, new(handler.NewUsr))
	fmt.Println("newusr 服务启动成功")
	srv.Serve(lis)

}
