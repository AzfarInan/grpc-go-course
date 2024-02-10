package main

import (
	"log"

	pb "github.com/AzfarInan/grpc-go-course/calculator/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var addr string = "localhost:50051"

func main() {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Failed to dial: %v", err)
	}

	defer conn.Close()

	c := pb.NewCalculatorClient(conn)
	// doSum(c)
	// doSubtract(c)
	// doMultiply(c)
	// doDivide(c)
	// doPrime(c)
	// doAverage(c)
	// doMax(c)
	// doSecondMax(c)
	// doSquareRoot(c, 10)
	doSquareRoot(c, -10)
	// doSquareRoot(c, 144)
}
