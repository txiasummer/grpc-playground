package main

import pb "grpc-playground/calculator/proto"

type Server struct {
	pb.CalculatorServiceServer
}
