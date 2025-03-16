package main

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "go-backend/gRPC-go/greetpb"

	"google.golang.org/grpc"
)

func main() {
	// Connect to gRPC server
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewGreeterClient(conn)

	// Call SayHello RPC
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	resp, err := client.SayHello(ctx, &pb.HelloRequest{Name: "Go Developer"})
	if err != nil {
		log.Fatalf("Could not greet: %v", err)
	}

	fmt.Println("Response from server:", resp.GetMessage())
}
