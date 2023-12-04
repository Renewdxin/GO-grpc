package main

import (
	"context"
	pb "github.com/Renewdxin/GO-grpc/proto"
)

func (s *helloServer) SayHello(ctx context.Context, req *pb.NoParam) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{Message: "hello"}, nil
}
