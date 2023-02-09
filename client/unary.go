package main

import (
	"context"
	"log"
	"time"

	"github.com/shubham/go-grpc/pb"
)

func callSayHello(client pb.GreetServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := client.SayHello(ctx, &pb.NoParam{})
	if err != nil {
		log.Fatalf("could not greet %v", err)
	}

	log.Print(res.Message)

}
