package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/AzfarInan/grpc-go-course/calculator/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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

// / Find the maximum number
func doMax(c pb.CalculatorClient) {
	log.Printf("doMax invoked")

	stream, err := c.Max(context.Background())

	if err != nil {
		log.Fatalf("Failed to call Max: %v", err)
	}

	reqs := []*pb.MaxRequest{
		{Number: 1},
		{Number: 5},
		{Number: 11},
		{Number: 6},
		{Number: 19},
		{Number: 14},
		{Number: 20},
	}

	listOfNumbers := []int32{}
	waitc := make(chan struct{})

	go func() {
		for _, req := range reqs {
			log.Printf("Sending request: %v", req)

			stream.Send(req)

			/// Add number to list
			listOfNumbers = append(listOfNumbers, req.GetNumber())

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
				log.Fatalf("Failed to receive: %v", err)
				break
			}

			log.Printf("Max number among %v: %v", listOfNumbers, res.Max)
		}
		close(waitc)
	}()

	<-waitc
}

// Find the Second Most Max Number
// / Find the maximum number
func doSecondMax(c pb.CalculatorClient) {
	log.Printf("doSecondMax invoked")

	stream, err := c.FindSecondMax(context.Background())

	if err != nil {
		log.Fatalf("Failed to call FindSecondMax: %v", err)
	}

	reqs := []*pb.MaxRequest{
		{Number: 20},
		{Number: 1},
		{Number: 3},
		{Number: 5},
		{Number: 11},
		{Number: 18},
		{Number: 8},
		{Number: 14},
		{Number: 4},
	}

	listOfNumbers := []int32{}
	waitc := make(chan struct{})

	go func() {
		for _, req := range reqs {
			log.Printf("Sending request: %v", req)

			stream.Send(req)

			/// Add number to list
			listOfNumbers = append(listOfNumbers, req.GetNumber())

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
				log.Fatalf("Failed to receive: %v", err)
				break
			}

			log.Printf("Second Max number among %v: %v", listOfNumbers, res.Max)
		}
		close(waitc)
	}()

	<-waitc
}

// / Compute the square root
func doSquareRoot(c pb.CalculatorClient, number int32) {
	log.Printf("doSquareRoot invoked")

	req := &pb.SqrtRequest{
		Number: number,
	}

	res, err := c.Sqrt(context.Background(), req)

	if err != nil {
		e, ok := status.FromError(err)

		if ok {
			log.Printf("Error message from server: %v", e.Message())
			log.Printf("Error code from server: %v", e.Code())

			if e.Code() == codes.InvalidArgument {
				log.Fatalf("We probably sent a negative number!")
				return
			}
		} else {
			log.Fatalf("Failed to call SquareRoot: %v", err)
		}
	}

	log.Printf("SquareRoot response: %v", res)
}
