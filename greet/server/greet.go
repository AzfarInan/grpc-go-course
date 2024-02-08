package main

import (
	"context"
	"fmt"
	"io"
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

func (s *Server) LongGreet(stream pb.GreetService_LongGreetServer) error {

	log.Printf("LongGreet function was invoked with a streaming request")

	res := ""

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return stream.SendAndClose(&pb.GreetResponse{
				Result: res,
			})
		}

		if err != nil {
			log.Fatalf("Failed to receive: %v", err)
		}

		res += "Hello " + req.GetFirstName() + "! "
	}
}
