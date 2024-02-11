package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/AzfarInan/grpc-go-course/blog/proto"
	"github.com/golang/protobuf/ptypes/empty"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) UpdateBlog(ctx context.Context, in *pb.Blog) (*empty.Empty, error) {
	log.Printf("UpdateBlog function was invoked with %v", in)

	oid, err := primitive.ObjectIDFromHex(in.GetId())

	if err != nil {
		return nil, status.Errorf(
			codes.InvalidArgument,
			fmt.Sprintf("Cannot parse ID"),
		)
	}

	data := &BlogItem{
		AuthorID: in.GetAuthorId(),
		Content:  in.GetContent(),
		Title:    in.GetTitle(),
	}

	res, err := collection.UpdateOne(
		ctx,
		bson.M{"_id": oid},
		bson.M{"$set": data},
	)

	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			"Could not update",
		)
	}

	if res.ModifiedCount == 0 {
		return nil, status.Errorf(
			codes.NotFound,
			"Could not find blog with Id",
		)
	}

	return &empty.Empty{}, nil
}

// / Delete Blog
func (s *Server) DeleteBlog(ctx context.Context, in *pb.BlogId) (*empty.Empty, error) {
	log.Printf("DeleteBlog function was invoked with %v", in)

	oid, err := primitive.ObjectIDFromHex(in.GetId())

	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Cannot parse ID"),
		)
	}

	res, err := collection.DeleteOne(
		ctx,
		bson.M{"_id": oid},
	)

	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			"Could not Delete",
		)
	}

	if res.DeletedCount == 0 {
		return nil, status.Errorf(
			codes.NotFound,
			"Could not find blog with Id",
		)
	}

	return &empty.Empty{}, nil
}
