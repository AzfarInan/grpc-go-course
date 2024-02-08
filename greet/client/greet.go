package main

import (
	"context"
	"io"
	"log"

	pb "github.com/AzfarInan/grpc-go-course/greet/proto"
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
