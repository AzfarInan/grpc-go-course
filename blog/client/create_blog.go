package main

import (
	"context"
	"log"

	pb "github.com/AzfarInan/grpc-go-course/blog/proto"
)

func createBlog(c pb.BlogServiceClient) string {
	log.Println("--------Create Blog--------")

	blog := &pb.Blog{
		AuthorId: "Azfar",
		Title:    "My First Blog",
		Content:  "Content of the first blog",
	}

	res, err := c.CreateBlog(context.Background(), blog)

	if err != nil {
		log.Fatalf("Failed to create blog: %v", err)
	}

	log.Printf("Blog has been created: %v", res)

	return res.Id
}
