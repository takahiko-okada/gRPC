package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net"

	"github.com/tkhkokd/grpc/calculator/calculatorpb"
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

func (*server) PrimeNumberDecomposition(req *calculatorpb.PrimeNumberDecompositionRequest, stream calculatorpb.CalculatorService_PrimeNumberDecompositionServer) error {
	fmt.Printf("Received PrimeNumberDecomposition RPC: %v", req)

	number := req.GetNumber()
	divisor := int64(2)

	for number > 1 {
		if number%divisor == 0 {
			stream.Send(&calculatorpb.PrimeNumberDecompositionResponse{
				PrimeFactor: divisor,
			})
			number = number / divisor
		} else {
			divisor++
			fmt.Printf("Divisor has increased to %v", divisor)
		}
	}
	return nil
}

func (*server) ComputeAverage(stream calculatorpb.CalculatorService_ComputeAverageServer) error {
	fmt.Printf("ComputeAverage function was invoked with a streaming request")
	sum := int32(0)
	count := 0
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			result := float64(sum) / float64(count)
			// End of File we have finished reading the client stream
			// SendAndClose called on requests(stream) and takes a response as the parameter
			return stream.SendAndClose(&calculatorpb.ComputeAverageResponse{
				// argument parameter
				Result: result,
			})
		}
		if err != nil {
			log.Fatalf("error while reading client stream %v", err)
		}

		sum += req.GetNumber()
		count++
	}
}

func (*server) FindMaximum(stream calculatorpb.CalculatorService_FindMaximumServer) error {
	// type CalculatorService_FindMaximumServer interface {
	//   SendAndClose(*FindMaximumResponse) error
	//   Recv() (*FindMaximumRequest, error)
	//   grpc.ServerStream
	// }

	// this type contains the server stream and the methods
	fmt.Printf("FindMaximum function was invoked with a streaming RPC request\n")
	maximum := int32(0)

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Fatalf("Error while reading client stream: %v", err)
			return err
		}
		number := req.GetNumber()
		if number > maximum {
			maximum = number
			sendErr := stream.Send(&calculatorpb.FindMaximumResponse{
				Result: maximum,
			})
			if sendErr != nil {
				log.Fatalf("Error while sending data to client: %v", sendErr)
				return err
			}
		}
	}
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
