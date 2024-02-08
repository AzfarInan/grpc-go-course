package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/AzfarInan/grpc-go-course/greet/proto"
)

func (s *Server) Greet(ctx context.Context, req *pb.GreetRequest) (*pb.GreetResponse, error) {

	log.Printf("Greet function was invoked with %v", req)

	return &pb.GreetResponse{
		Result: "Hello " + req.FirstName,
	}, nil
}

func (s *Server) GreetManyTimes(req *pb.GreetRequest, stream pb.GreetService_GreetManyTimesServer) error {

	log.Printf("GreetManyTimes function was invoked with %v", req)

	firstName := req.FirstName

	for i := 0; i < 10; i++ {
		result := "Hello " + firstName + ", number " + fmt.Sprint(i)
		res := &pb.GreetResponse{
			Result: result,
		}

		stream.Send(res)
	}

	return nil
}
