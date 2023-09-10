package main

import (
	"context"
	"log"

	pb "grpc-playground/calculator/proto"
)

func doSum(c pb.CalculatorServiceClient) {
	req := &pb.SumRequest{
		X: 88,
		Y: 12,
	}
	res, err := c.Sum(context.Background(), req)
	log.Println("Sum was invoked with:", req)
	if err != nil {
		log.Fatalf("Could not sum: %v\n", err)
	}
	log.Println("result =", res.Result)
}
