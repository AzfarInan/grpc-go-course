package main

import (
	"context"
	"io"
	"log"

	pb "github.com/AzfarInan/grpc-go-course/blog/proto"
	"google.golang.org/protobuf/types/known/emptypb"
)

func listBlogs(c pb.BlogServiceClient) {
	log.Println("--------List Blogs--------")

	stream, err := c.ListBlogs(context.Background(), &emptypb.Empty{})

	if err != nil {
		log.Fatalf("Failed to list blogs: %v", err)
	}

	for {
		res, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Failed to receive data: %v", err)
		}

		log.Printf("Blog: %v", res)
	}
}
