package main

import (
	"context"
	"fmt"
	"log"

	"github.com/ltbatista/sum-api-grpc/sum/sumpb"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Hello, I'm Sum The Client")
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer cc.Close()

	c := sumpb.NewSumServiceClient(cc)
	doUnary(c)
}

func doUnary(c sumpb.SumServiceClient) {
	fmt.Println("Starting to do an Unary RPC...")
	req := &sumpb.SumRequest{
		Sum: &sumpb.Sum{
			FirstNumber:  101,
			SecondNumber: 2,
		},
	}
	res, err := c.Sum(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling Sum RPC: %v", err)
	}
	log.Printf("Response from Greet:  %v", res.Result)
}
