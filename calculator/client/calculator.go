package main

import (
	"context"
	pb "grpc-playground/calculator/proto"
	"log"
)

func calculate(c pb.CalculatorServiceClient) {
	req := &pb.CalculatorRequest{
		X: 88,
		Y: 12,
	}
	res, err := c.Calculator(context.Background(), req)
	log.Println("Calculate was invoked with:", req)
	if err != nil {
		log.Fatalf("Could not calculate: %v\n", err)
	}
	log.Println("result =", res.Result)
}
