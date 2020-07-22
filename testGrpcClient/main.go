package main

import (
	"context"
	"log"
	"os"
	pb "testgrpcclient/pb"
	"time"
	"fmt"

	"google.golang.org/grpc"
)
const(
	address = "localhost:9009"
	defaultName = "world"
)

func main() {
	//set up a connection to server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithTimeout(time.Second*5), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	fmt.Println("conn: ", conn)
	defer conn.Close()
	c := pb.NewGreeterClient(conn)

	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.SayHello(ctx, &pb.HelloRequest{Name: name})
	if err != nil {
		log.Fatalf("could not gret: %v", err)
	}
	log.Printf("Greeting: %s", r.Message)
}
