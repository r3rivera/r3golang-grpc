package main

import (
	pb "github.com/r3rivera/r3app-protobuffer-repo/basic-test"

	"log"
	"net"

	"google.golang.org/grpc"
)

type server struct{}

func main() {
	log.Println("Main Server...")

	listener, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
		panic(err)
	}

	s := grpc.NewServer()
	pb.RegisterGreeterServiceServer(s, &server{})

	if err := s.Serve(listener); err != nil {
		log.Fatalf("Failed to server: %v", err)
	}
	log.Println("Main server is up and running...")
}
