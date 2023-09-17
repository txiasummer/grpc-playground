package main

import (
	"context"
	pb "grpc-playground/calculator/proto"
	"log"
)

func doAverage(c pb.CalculatorServiceClient) {
	log.Println("doAverage was invoked")
	stream, err := c.Average(context.Background())
	if err != nil {
		log.Fatalf("Error while calling doAverage: %v\n", err)
	}
	for _, req := range []int32{123, 234, 345, 456, 567, 678, 789} {
		log.Printf("Sending req: %v\n", req)
		stream.Send(&pb.AverageRequest{X: req})
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error while receiving response from doAverage: %v\n", err)
	}
	log.Printf("doAverage result: %s\n", res.Result)
}
