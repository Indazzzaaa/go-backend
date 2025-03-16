package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "go-backend/gRPC-go/greetpb"

	"google.golang.org/grpc"
)

// Define server struct
type server struct {
	pb.UnimplementedGreeterServer
}

// Implement SayHello method
func (s *server) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	log.Printf("Received request from: %s", req.GetName())
	return &pb.HelloResponse{Message: fmt.Sprintf("Hello, %s!", req.GetName())}, nil
}

func main() {
	// Start listening on a TCP port
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	// Create a new gRPC server
	grpcServer := grpc.NewServer()
	pb.RegisterGreeterServer(grpcServer, &server{})

	log.Println("gRPC server listening on port 50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
