package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/tkhkokd/grpc/calculator/calculatorpb"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Hello I'm a Calculator client")
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer cc.Close()

	c := calculatorpb.NewCalculatorServiceClient(cc)
	// fmt.Printf("Created client: %f", c)
	// doUnary(c)

	// doServerStreaming(c)
	// doClientStreaming(c)
	doBiDiStreaming(c)
}

func doUnary(c calculatorpb.CalculatorServiceClient) {
	fmt.Println("Starting to do a Sum Unary RPC...")
	req := &calculatorpb.SumRequest{
		FirstNumber:  5,
		SecondNumber: 10,
	}

	res, err := c.Sum(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling Sum RPC: %v", err)
	}
	log.Printf("Response from Sum: %v", res.SumResult)
}

func doServerStreaming(c calculatorpb.CalculatorServiceClient) {
	fmt.Println("Starting to do a PrimeNumberDecomposition Server Streaming RPC...")
	req := &calculatorpb.PrimeNumberDecompositionRequest{
		Number: 7312987,
	}

	stream, err := c.PrimeNumberDecomposition(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling PrimeNumberDecomposition RPC: %v", err)
	}
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Something happened: %v", err)
		}
		fmt.Println(res.GetPrimeFactor())
	}
}

// CalculatorServiceClient is and interface which consists of
// available methods
func doClientStreaming(c calculatorpb.CalculatorServiceClient) {
	fmt.Println("Starting to do a Client Streaming RPC...")
	// Define multiple requests
	requests := []int32{3, 5, 9, 54, 23}

	stream, err := c.ComputeAverage(context.Background())
	if err != nil {
		log.Fatalf("error while opening stream: %v", err)
	}

	for _, req := range requests {
		stream.Send(&calculatorpb.ComputeAverageRequest{
			Number: req,
		})
	}
	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("error while receiving response from ComputeAverage: %v", err)
	}
	fmt.Printf("ComputeAverageResponse: %v\n", res)
}

func doBiDiStreaming(c calculatorpb.CalculatorServiceClient) {
	fmt.Println("Starting to do a BiDi Streaming RPC...")
	stream, err := c.FindMaximum(context.Background())
	if err != nil {
		log.Fatalf("Error while creating stream: %v", err)
		return
	}

	waitc := make(chan struct{})

	go func() {
		requests := []int32{3, 5, 9, 54, 21, 23, 100}
		for _, req := range requests {
			fmt.Printf("Sending message: %v\n", req)
			stream.Send(&calculatorpb.FindMaximumRequest{
				Number: req,
			})
			time.Sleep(1000 * time.Millisecond)
		}
		stream.CloseSend()
	}()

	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("Error while receiving: %v\n", err)
				break
			}
			maximum := res.GetResult()
			fmt.Printf("Received a new maximum of...: %v\n", maximum)
		}
		close(waitc)
	}()

	<-waitc
}
