package main

import (
	"log"

	"github.com/shubham/go-grpc/pb"
	"google.golang.org/grpc"
)

const (
	port = ":9999"
)

func main() {
	connection, err := grpc.Dial("localhost"+port, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("falied to connect %v", err)
	}

	defer connection.Close()

	client := pb.NewGreetServiceClient(connection)

	names := &pb.NamesList{
		Names: []string{"Shubham", "Jarvis", "Edith"},
	}

	// callSayHello(client)
	// callSayHelloServerStream(client, names)
	// callSayHelloClientStream(client, names)
	callSayHelloBidirectionalStream(client, names)
}
