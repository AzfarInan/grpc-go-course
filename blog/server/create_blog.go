package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/AzfarInan/grpc-go-course/blog/proto"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateBlog(ctx context.Context, in *pb.Blog) (*pb.BlogId, error) {
	log.Printf("CreateBlog function was invoked with %v", in)

	data := BlogItem{
		AuthorID: in.GetAuthorId(),
		Content:  in.GetContent(),
		Title:    in.GetTitle(),
	}

	res, err := collection.InsertOne(ctx, data)

	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Internal error: %v", err),
		)
	}

	oid, ok := res.InsertedID.(primitive.ObjectID)

	if !ok {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Cannot convert to OID"),
		)
	}

	return &pb.BlogId{
		Id: oid.Hex(),
	}, nil
}
