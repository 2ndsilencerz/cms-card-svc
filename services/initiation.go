package services

import (
	"github.com/2ndSilencerz/cms-card-svc/services/pb"
	"google.golang.org/grpc"
)

// InitAllServices to init all endpoint of grpc
func InitAllServices(grpcServer *grpc.Server) {
	pb.RegisterCardListServer(grpcServer, &server{})
}
