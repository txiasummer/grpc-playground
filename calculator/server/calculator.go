package main

import (
	"context"
	"log"

	pb "grpc-playground/calculator/proto"
)

func (s *Server) Calculator(ctx context.Context, in *pb.CalculatorRequest) (*pb.CalculatorResponse, error) {
	log.Printf("Calculator function was invoked with %v\n", in)
	return &pb.CalculatorResponse{
		Result: in.X + in.Y,
	}, nil
}
