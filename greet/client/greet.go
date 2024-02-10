package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/AzfarInan/grpc-go-course/greet/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func doGreet(c pb.GreetServiceClient) {
	log.Println("doGreet invoked")

	res, err := c.Greet(context.Background(), &pb.GreetRequest{
		FirstName: "Azfar",
	})

	if err != nil {
		log.Fatalf("Failed to call Greet: %v", err)
	}

	log.Printf("Greet response: %v", res.Result)
}

func doGreetManyTimes(c pb.GreetServiceClient) {
	log.Println("doGreetManyTimes invoked")

	req := &pb.GreetRequest{
		FirstName: "Azfar",
	}

	resStream, err := c.GreetManyTimes(context.Background(), req)

	if err != nil {
		log.Fatalf("Failed to call GreetManyTimes: %v", err)
	}

	for {
		res, err := resStream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Failed to receive response: %v", err)
		}

		log.Printf("GreetManyTimes response: %v", res.Result)
	}
}

func doLongGreet(c pb.GreetServiceClient) {
	log.Println("doLongGreet invoked")

	reqs := []*pb.GreetRequest{
		{FirstName: "Tamanna"},
		{FirstName: "Nusrat"},
		{FirstName: "Sumaiya"},
	}

	stream, err := c.LongGreet(context.Background())

	if err != nil {
		log.Fatalf("Failed to call LongGreet: %v", err)
	}

	for _, req := range reqs {
		log.Println("Sending request:", req)

		stream.Send(req)
		time.Sleep(1 * time.Second)
	}

	res, err := stream.CloseAndRecv()

	if err != nil {
		log.Fatalf("Failed to receive response: %v", err)
	}

	log.Printf("LongGreet response: %v", res.Result)
}

func doGreetEveryone(c pb.GreetServiceClient) {
	log.Println("doGreetEveryone invoked")

	reqs := []*pb.GreetRequest{
		{FirstName: "Tamanna"},
		{FirstName: "Nusrat"},
		{FirstName: "Sumaia"},
	}

	stream, err := c.GreetEveryone(context.Background())

	if err != nil {
		log.Fatalf("Failed to call GreetEveryone: %v", err)
	}

	waitc := make(chan struct{})

	go func() {
		for _, req := range reqs {
			log.Println("Sending request:", req)

			stream.Send(req)
			time.Sleep(1 * time.Second)
		}

		stream.CloseSend()
	}()

	go func() {
		for {
			res, err := stream.Recv()

			if err == io.EOF {
				break
			}

			if err != nil {
				log.Fatalf("Failed to receive response: %v", err)
				break
			}

			log.Printf("GreetEveryone response: %v", res.Result)
		}

		close(waitc)
	}()

	<-waitc
}

// / Greet with Deadlines
func doGreetWithDeadline(c pb.GreetServiceClient, timeout time.Duration) {
	log.Println("doGreetWithDeadline invoked")

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	res, err := c.GreetWithDeadline(ctx, &pb.GreetRequest{
		FirstName: "Azfar",
	})

	if err != nil {
		e, ok := status.FromError(err)

		if ok {
			if e.Code() == codes.DeadlineExceeded {
				log.Println("The request was timeout!")
				return
			} else {
				log.Fatalf("Failed to call GreetWithDeadline: %v", err)
			}
		} else {
			log.Fatalf("Failed to call GreetWithDeadline: %v", err)
		}

		return
	}

	log.Printf("GreetWithDeadline response: %v", res.Result)
}
