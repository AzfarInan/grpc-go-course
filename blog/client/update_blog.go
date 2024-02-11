package main

import (
	"context"
	"log"

	pb "github.com/AzfarInan/grpc-go-course/blog/proto"
)

func updateBlog(c pb.BlogServiceClient, id string) {
	log.Println("--------Update Blog--------")

	newBlog := &pb.Blog{
		Id:       id,
		AuthorId: "Updated Author",
		Title:    "My Blog Number 2 (edited)",
		Content:  "Content of the second blog (edited)",
	}

	_, err := c.UpdateBlog(context.Background(), newBlog)

	if err != nil {
		log.Fatalf("Failed to update blog: %v", err)
	}

	log.Println("Blog has been updated")
}

// Delete Blog
func deleteBlog(c pb.BlogServiceClient, id string) {
	log.Println("--------Delete Blog--------")

	_, err := c.DeleteBlog(context.Background(), &pb.BlogId{Id: id})

	if err != nil {
		log.Fatalf("Failed to delete blog: %v", err)
	}

	log.Println("Blog has been deleted")
}
