package main

import (
	"context"
	"log"

	pb "github.com/AzfarInan/grpc-go-course/calculator/proto"
)

func (s *Server) Add(ctx context.Context, req *pb.AddRequest) (*pb.AddResponse, error) {
	log.Printf("Add function was invoked with %v", req)

	return &pb.AddResponse{
		Result: req.GetA() + req.GetB(),
	}, nil
}
