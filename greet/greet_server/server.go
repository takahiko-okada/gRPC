package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/tkhkokd/grpc/greet/greetpb"
	"google.golang.org/grpc"
)

type server struct{}

func (*server) Greet(ctx context.Context, req *greetpb.GreetRequest) (*greetpb.GreetResponse, error) {
	fmt.Printf("Greet function was invoked with %v\n", req)
	firstName := req.GetGreeting().GetFirstName()
	// if req is not nil then GetGreeting returns the Greeting field
	// in the GreetRequest struct, of type Greeting
	result := "Hello " + firstName
	res := &greetpb.GreetResponse{
		Result: result,
	}
	return res, nil
}

func main() {
	fmt.Println("Hello World")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	// parameters => (s *grpc.Server, srv GreetServiceServer)
	greetpb.RegisterGreetServiceServer(s, &server{})
	// greetpb is the package imported, seed the definition by hovering the method
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve %v", err)
	}
}
