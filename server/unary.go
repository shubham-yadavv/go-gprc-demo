package main

import (
	"context"

	"github.com/shubham/go-grpc/pb"
)

func (s *helloServer) SayHello(ctx context.Context, req *pb.NoParam) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{
		Message: "hello",
	}, nil
}
