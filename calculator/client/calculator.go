package main

import (
	"context"
	"log"

	pb "github.com/AzfarInan/grpc-go-course/calculator/proto"
)

func doSum(c pb.CalculatorClient) {
	log.Println("doSum invoked")

	req := &pb.AddRequest{
		A: 3,
		B: 10,
	}

	res, err := c.Add(context.Background(), req)

	if err != nil {
		log.Fatalf("Failed to add: %v", err)
	}

	log.Printf("Result: %v", res.Result)
}
