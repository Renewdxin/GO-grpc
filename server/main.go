package main

import (
	pb "github.com/Renewdxin/GO-grpc/proto"
	"google.golang.org/grpc"
	"log"
	"net"
)

const port = ":9000"

type helloServer struct {
	pb.GreetServiceServer
}

func main() {
	listen, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to start the port %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterGreetServiceServer(grpcServer, &helloServer{})
	log.Printf("server started at %v", listen.Addr())
	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("failed to start : %v", err)
	}

}
