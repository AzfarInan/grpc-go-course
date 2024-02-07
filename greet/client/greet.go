package main

import (
	"context"
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
