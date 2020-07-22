package main

import (
	"context"
	"log"
	"net"
	pb "testgrpcserver/pb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const(
	address = "localhost:50051"
	port = "9009"
	defaultName = "world"
)

type server struct{}

func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: "Hello " + in.Name}, nil
}

func main() {
	//set up a connection to server.
	lis, err := net.Listen("tcp", "localhost:9009")
	if err != nil {
		log.Fatalf("failed to lister: %v", err)
	}

	s := grpc.NewServer()

	pb.RegisterGreeterServer(s, &server{})

	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
