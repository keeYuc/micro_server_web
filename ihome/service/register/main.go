package main

import (
	"register/handler"
	pb "register/proto"

	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/logger"
)

func main() {
	// Create service
	srv := service.New(
		service.Name("register"),
		service.Version("latest"),
	)

	// Register handler
	pb.RegisterRegisterHandler(srv.Server(), new(handler.Register))

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
