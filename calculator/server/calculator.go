package main

import (
	"context"
	"io"
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

/// Compute Average

func (s *Server) Average(stream pb.Calculator_AverageServer) error {
	log.Printf("Average function was invoked with a streaming request")

	sum := float32(0)
	count := float32(0)

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return stream.SendAndClose(&pb.AverageResponse{
				Average: sum / count,
			})
		}

		if err != nil {
			log.Fatalf("Failed to receive: %v", err)
		}

		sum += float32(req.GetNumber())
		count++
	}
}

/// Find Maximum

func (s *Server) Max(stream pb.Calculator_MaxServer) error {
	log.Printf("Max function was invoked with a streaming request")

	maxNumber := int32(0)

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return nil
		}

		if err != nil {
			log.Fatalf("Failed to receive: %v", err)
		}

		res := req.GetNumber()

		if res > maxNumber {
			maxNumber = res
		}

		err = stream.Send(&pb.MaxResponse{
			Max: maxNumber,
		})

		if err != nil {
			log.Fatalf("Failed to send: %v", err)
		}
	}
}
