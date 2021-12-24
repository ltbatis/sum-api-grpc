package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/ltbatista/sum-api-grpc/sum/sumpb"
	"google.golang.org/grpc"
)

type server struct{}

func (*server) Sum(ctx context.Context, req *sumpb.SumRequest) (*sumpb.SumResponse, error) {
	fmt.Printf("Sum function was invoked with %v\n ", req)
	first_number := req.GetSum().FirstNumber
	second_number := req.GetSum().SecondNumber
	result := first_number + second_number
	res := &sumpb.SumResponse{
		Result: result,
	}
	return res, nil
}

func main() {
	fmt.Println("Hello, I'm Sum The Server.")
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	sumpb.RegisterSumServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
