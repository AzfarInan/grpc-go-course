package main

import (
	"context"
	"log"

	pb "github.com/AzfarInan/grpc-go-course/calculator/proto"
)

func doSum(c pb.CalculatorClient) {
	log.Println("doSum invoked")

	req := &pb.CalculatorRequest{
		A: 3,
		B: 10,
	}

	res, err := c.Add(context.Background(), req)

	if err != nil {
		log.Fatalf("Failed to add: %v", err)
	}

	log.Printf("Result: %v", res.Result)
}

func doSubtract(c pb.CalculatorClient) {
	log.Println("doSubtract invoked")

	req := &pb.CalculatorRequest{
		A: 3,
		B: 10,
	}

	res, err := c.Subtract(context.Background(), req)

	if err != nil {
		log.Fatalf("Failed to subtract: %v", err)
	}

	log.Printf("Result: %v", res.Result)
}

func doMultiply(c pb.CalculatorClient) {
	log.Println("doMultiply invoked")

	req := &pb.CalculatorRequest{
		A: 3,
		B: 10,
	}

	res, err := c.Multiply(context.Background(), req)

	if err != nil {
		log.Fatalf("Failed to multiply: %v", err)
	}

	log.Printf("Result: %v", res.Result)
}

func doDivide(c pb.CalculatorClient) {
	log.Println("doDivide invoked")

	req := &pb.CalculatorRequest{
		A: 30,
		B: 10,
	}

	res, err := c.Divide(context.Background(), req)

	if err != nil {
		log.Fatalf("Failed to divide: %v", err)
	}

	log.Printf("Result: %v", res.Result)
}
