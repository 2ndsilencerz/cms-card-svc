package services

import (
	"github.com/2ndsilencerz/cms-card-svc/models/pb"
	"google.golang.org/grpc"
)

// InitAllServices to init all endpoint of grpc
func InitAllServices(grpcServer *grpc.Server) {
	pb.RegisterCardListServer(grpcServer, &Server{})
}
