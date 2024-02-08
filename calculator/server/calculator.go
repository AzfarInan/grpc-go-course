package main

import (
	"context"
	"log"

	pb "github.com/AzfarInan/grpc-go-course/calculator/proto"
)

func (s *Server) Add(ctx context.Context, req *pb.CalculatorRequest) (*pb.CalculatorResponse, error) {
	log.Printf("Add function was invoked with %v", req)

	return &pb.CalculatorResponse{
		Result: req.GetA() + req.GetB(),
	}, nil
}

func (s *Server) Subtract(ctx context.Context, req *pb.CalculatorRequest) (*pb.CalculatorResponse, error) {
	log.Printf("Subtract function was invoked with %v", req)

	return &pb.CalculatorResponse{
		Result: req.GetA() - req.GetB(),
	}, nil
}

func (s *Server) Multiply(ctx context.Context, req *pb.CalculatorRequest) (*pb.CalculatorResponse, error) {
	log.Printf("Multiply function was invoked with %v", req)

	return &pb.CalculatorResponse{
		Result: req.GetA() * req.GetB(),
	}, nil
}

func (s *Server) Divide(ctx context.Context, req *pb.CalculatorRequest) (*pb.CalculatorResponse, error) {
	log.Printf("Divide function was invoked with %v", req)

	return &pb.CalculatorResponse{
		Result: req.GetA() / req.GetB(),
	}, nil
}

/// Prime Number Decomposition

func (s *Server) Prime(req *pb.PrimeRequest, stream pb.Calculator_PrimeServer) error {
	log.Printf("Prime function was invoked with %v", req)

	k := int32(2)
	N := req.GetNumber()

	for N > 1 {
		if N%k == 0 {
			res := &pb.CalculatorResponse{
				Result: k,
			}
			stream.Send(res)
			N = N / k
		} else {
			k = k + 1
		}
	}

	return nil
}
