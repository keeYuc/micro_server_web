package main

import (
	"fmt"
	"getimg/handler"
	"keeyu/message"
	"net"

	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":8787")
	if err != nil {
		return
	}
	defer lis.Close()
	srv := grpc.NewServer()
	message.RegisterGetimgServer(srv, new(handler.Getimg))
	fmt.Println("getimg 服务启动成功")
	srv.Serve(lis)
}

// // func main() {
// 	//consul
// 	// config := stdconsul.DefaultConfig()
// 	// client, err := stdconsul.NewClient(config)
// 	// if err != nil {
// 	// 	return
// 	// }
// 	// reg := stdconsul.AgentServiceRegistration{
// 	// 	ID:      "getimg",
// 	// 	Tags:    []string{"getimg"},
// 	// 	Name:    "getimg",
// 	// 	Address: "127.0.0.1",
// 	// 	Port:    11221,
// 	// 	Check: &stdconsul.AgentServiceCheck{
// 	// 		CheckID:  "getimg test",
// 	// 		TCP:      "127.0.0.1:11221",
// 	// 		Timeout:  "1s",
// 	// 		Interval: "5s",
// 	// 	},
// 	// }
// 	// client.Agent().ServiceRegister(&reg)

// 	// Create service
// 	srv := service.New(
// 		service.Name("getimg"),
// 		service.Version("latest"),
// 		service.Address(":11221"),
// 	)

// 	// Register handler
// 	pb.RegisterGetimgHandler(srv.Server(), new(handler.Getimg))

// 	// Run service
// 	if err := srv.Run(); err != nil {
// 		fmt.Println(err)
// 		// logger.Fatal(err)
// 	}
// }
