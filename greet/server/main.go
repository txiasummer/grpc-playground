package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
	pb "grpc-playground/greet/proto"
)

const addr = "0.0.0.0:50051"

type Server struct {
	pb.GreetServiceServer
}

func main() {
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Failed to listen on %v\n", err)
	}
	log.Printf("Listening on %s\n", addr)
	s := grpc.NewServer()
	pb.RegisterGreetServiceServer(s, &Server{})
	if err = s.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v\n", err)
	}
}
