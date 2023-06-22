package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)
import "google.golang.org/grpc"
import pb "go-grpc-playground/protos"

func main() {
	fmt.Println("Start grpc Client")

	conn, err := grpc.Dial(":50051", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		fmt.Printf("Error: %s\n", err)
	}

	defer conn.Close()

	client := pb.NewCalculatorServiceClient(conn)

	req := &pb.AddRequest{Num1: 5, Num2: 10}
	res, err := client.Add(context.Background(), req)
	if err != nil {
		log.Fatalf("Add RPC failed: %v", err)
	}

	log.Printf("Server response: the sum is: %d", res.Sum)
}
