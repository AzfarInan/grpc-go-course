package main

import (
	"context"
	"log"

	pb "github.com/AzfarInan/grpc-go-course/blog/proto"
)

func createBlog(c pb.BlogServiceClient) string {
	log.Println("--------Create Blog--------")

	blog := &pb.Blog{
		AuthorId: "Inan",
		Title:    "My Blog Number 2",
		Content:  "Content of the second blog",
	}

	res, err := c.CreateBlog(context.Background(), blog)

	if err != nil {
		log.Fatalf("Failed to create blog: %v", err)
	}

	log.Printf("Blog has been created: %v", res)

	return res.Id
}
