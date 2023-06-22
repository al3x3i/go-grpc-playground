package main

import (
	"context"
	pb "go-grpc-playground/protos"
	"google.golang.org/grpc"
	"log"
	"net"
)

type calculatorServer struct {
}

func (s *calculatorServer) Add(ctx context.Context, req *pb.AddRequest) (*pb.AddResponse, error) {
	sum := req.Num1 + req.Num2
	return &pb.AddResponse{Sum: sum}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	server := grpc.NewServer()
	pb.RegisterCalculatorServiceServer(server, &calculatorServer{})
	log.Println("Starting gRPC server on port 50051...")
	if err := server.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
