package main

import (
	"fmt"
	"log"
	"net"

	"github.com/2ndSilencerz/cms-card-svc/configs/utils"

	"github.com/2ndSilencerz/cms-card-svc/configs"
	"github.com/2ndSilencerz/cms-card-svc/services"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {

	port := configs.GetDefaultPort()
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	services.InitAllServices(grpcServer)
	reflection.Register(grpcServer)
	utils.PrintToConsole(fmt.Sprint("Starting service"))
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
