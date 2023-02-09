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
	// listen on the port
	listen, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to start server %v", err)
	}

	// create a new grpc server
	grpcServer := grpc.NewServer()

	// register the greet service
	pb.RegisterGreetServiceServer(grpcServer, &helloServer{})
	log.Printf("server started at %v", listen.Addr())

	// start the server
	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("failed to start servet %v", err)
	}

}
