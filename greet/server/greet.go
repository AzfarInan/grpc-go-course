package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	pb "github.com/AzfarInan/grpc-go-course/greet/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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

func (s *Server) GreetEveryone(stream pb.GreetService_GreetEveryoneServer) error {
	log.Printf("GreetEveryone function was invoked with a streaming request")

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return nil
		}

		if err != nil {
			log.Fatalf("Failed to receive: %v", err)
			return err
		}

		firstName := req.GetFirstName()
		result := "Hello " + firstName + "!"

		err = stream.Send(&pb.GreetResponse{
			Result: result,
		})

		if err != nil {
			log.Fatalf("Failed to send: %v", err)
			return err
		}
	}
}

// / Greet with Deadlines
func (s *Server) GreetWithDeadline(ctx context.Context, req *pb.GreetRequest) (*pb.GreetResponse, error) {
	log.Printf("GreetWithDeadline function was invoked with %v", req)

	for i := 0; i < 3; i++ {
		if ctx.Err() == context.DeadlineExceeded {
			log.Println("The client canceled the request!")
			return nil, status.Error(
				codes.Canceled,
				"The client canceled the request!",
			)
		}

		time.Sleep(1 * time.Second)
	}

	return &pb.GreetResponse{
		Result: "Hello " + req.FirstName,
	}, nil
}
