package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
	// protoc  --go_out=./chat chat.proto --> chat.pb.go
)

func main() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	log.Println("Starting gRPC server...")

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %s", err)
	}

}
