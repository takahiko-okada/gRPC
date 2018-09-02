package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/tkhkokd/grpc/calculator/calculator   pb"
	"google.golang.org/grpc"
)

type server struct{}

func (*server) Sum(ctx context.Context, req *calculatorpb.SumRequest) (*calculatorpb.SumResponse, error) {
	fmt.Printf("Received Sum RPC: %v", req)
	first_number := req.FirstNumber
	second_number := req.SecondNumber
	sum := first_number + second_number
	res := &calculatorpb.SumResponse{
		SumResult: sum,
	}
	return res, nil
}

func main() {
	fmt.Println("Calculator server")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	// parameters => (s *grpc.Server, srv GreetServiceServer)
	calculatorpb.RegisterCalculatorServiceServer(s, &server{})
	// greetpb is the package imported, seed the definition by hovering the method
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve %v", err)
	}
}
