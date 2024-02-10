package main

import (
	"log"
	"net"

	pb "github.com/AzfarInan/grpc-go-course/calculator/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var addr string = "0.0.0.0:50051"

type Server struct {
	pb.CalculatorServer
}

func main() {
	lis, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatalf("Failed to listen on: %v", err)
	}

	log.Printf("Listening on: %v", addr)

	s := grpc.NewServer()
	pb.RegisterCalculatorServer(s, &Server{})
	// For Evans CLI
	reflection.Register(s)

	if err = s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}

}
