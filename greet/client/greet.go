package main

import (
	"context"
	"log"

	pb "grpc-playground/greet/proto"
)

func doGreet(c pb.GreetServiceClient) {
	log.Println("doGreet was invoked")
	res, err := c.Greet(context.Background(), &pb.GreetRequest{
		FirstName: "Hailey",
	})
	if err != nil {
		log.Fatalf("Could not greet: %v\n", err)
	}
	log.Printf("Greeting: %s\n", res.Result)
}
