package main

import (
	"context"
	"io"
	"log"
	"time"

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

/// Prime Number Decomposition

func doPrime(c pb.CalculatorClient) {
	log.Println("doPrime invoked")

	req := &pb.PrimeRequest{
		Number: 120,
	}

	resStream, err := c.Prime(context.Background(), req)

	if err != nil {
		log.Fatalf("Failed to call Prime: %v", err)
	}

	for {
		res, err := resStream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Failed to receive response: %v", err)
		}

		log.Printf("Prime response: %v", res.Result)
	}
}

// / Compute Average
func doAverage(c pb.CalculatorClient) {
	log.Printf("doAverage invoked")

	reqs := []*pb.PrimeRequest{
		{Number: 1},
		{Number: 2},
		{Number: 3},
		{Number: 4},
	}

	stream, err := c.Average(context.Background())

	if err != nil {
		log.Fatalf("Failed to call Average: %v", err)
	}

	for _, req := range reqs {
		log.Printf("Sending request: %v", req)

		stream.Send(req)
		time.Sleep(1 * time.Second)
	}

	res, err := stream.CloseAndRecv()

	if err != nil {
		log.Fatalf("Failed to receive response: %v", err)
	}

	log.Printf("Average response: %v", res.Average)
}
