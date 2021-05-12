package main

import (
	"fmt"
	"log"
	"net"

	"github.com/cmwylie19/grpc-go/greet/greetpb"
	"google.golang.org/grpc"
)

// server is used to implement helloworld.GreeterServer.
type server struct {
	// Embed the unimplemented server
	greetpb.UnimplementedGreetServiceServer
}

func main() {
	fmt.Println("hi")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Could not listen: %v", err)
	}
	s := grpc.NewServer()
	greetpb.RegisterGreetServiceServer(s, &server{})
	// greetpb.RegisterGreetServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
