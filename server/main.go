package main

import (
	"log"
	"net"

	pb "github.com/shubham/go-grpc/pb"
	"google.golang.org/grpc"
)

const (
	port = ":9999"
)

type helloServer struct {
	pb.GreetServiceServer
}

func main() {

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to start server %v", err)
	}

	// create a new gRPC server
	grpcServer := grpc.NewServer()

	// register the greet service
	pb.RegisterGreetServiceServer(grpcServer, &helloServer{})
	log.Printf("Server started at %v", lis.Addr())

	//list is the port, the grpc server needs to start there
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to start: %v", err)
	}
}
