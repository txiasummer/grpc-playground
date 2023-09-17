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
	for _, num := range []int32{123, 234, 345, 456, 567, 678, 789} {
		log.Printf("Sending number: %v\n", num)
		stream.Send(&pb.AverageRequest{X: num})
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error while receiving response from doAverage: %v\n", err)
	}
	log.Printf("doAverage result: %f\n", res.Result)
}
