package test

import (
	"log"

	"google.golang.org/grpc"
)

// ConnectionConfig ...
func ConnectionConfig() *grpc.ClientConn {
	conn, err := grpc.Dial(":9991", grpc.WithInsecure())
	if err != nil {
		log.Panic("no connection can be made. make sure the service is already started")
	}
	return conn
}
