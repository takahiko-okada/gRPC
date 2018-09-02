package main

import (
	"fmt"
	"log"
	"net"

	"github.com/tkhkokd/grpc/greet/greetpb"
	"google.golang.org/grpc"
)

type server struct{}

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
