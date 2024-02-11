package main

import (
	"context"
	"log"

	pb "github.com/AzfarInan/grpc-go-course/blog/proto"
)

func readBlog(c pb.BlogServiceClient, blogID string) *pb.Blog {
	log.Println("--------Read Blog--------")

	blog, err := c.ReadBlog(context.Background(), &pb.BlogId{Id: blogID})

	if err != nil {
		log.Fatalf("Failed to read blog: %v", err)
	}

	log.Printf("Blog has been read: %v", blog)

	return blog
}
