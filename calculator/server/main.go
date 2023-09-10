package main

import (
	"google.golang.org/grpc/reflection"
	"log"
	"net"

	"google.golang.org/grpc"

	pb "grpc-playground/calculator/proto"
)

const addr = "0.0.0.0:50051"

type Server struct {
	pb.CalculatorServiceServer
}

func main() {
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Failed to listen on: %v\n", err)
	}
	log.Printf("Listening on %s\n", addr)

	var opts []grpc.ServerOption
	s := grpc.NewServer(opts...)
	pb.RegisterCalculatorServiceServer(s, &Server{})
	reflection.Register(s)

	if err = s.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v\n", err)
	}
}
