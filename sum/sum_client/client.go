package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/ltbatista/sum-api-grpc/sum/sumpb"
	"google.golang.org/grpc"
)

func main() {
	first_number_arg, second_number_arg := readArgs()

	fmt.Println("Hello, I'm Sum The Client")
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer cc.Close()

	c := sumpb.NewSumServiceClient(cc)
	doUnary(c, first_number_arg, second_number_arg)
}

func readArgs() (int, int) {
	first_number_arg, _ := strconv.Atoi(os.Args[1])
	second_number_arg, _ := strconv.Atoi(os.Args[2])
	return first_number_arg, second_number_arg
}

func doUnary(c sumpb.SumServiceClient, one int, two int) {
	fmt.Println("Starting to do an Unary RPC...")
	req := &sumpb.SumRequest{
		Sum: &sumpb.Sum{
			FirstNumber:  int32(one),
			SecondNumber: int32(two),
		},
	}
	res, err := c.Sum(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling Sum RPC: %v", err)
	}
	log.Printf("Response from Sum:  %v", res.Result)
}
