package main

import (
	"context"
	pb "go_lang_study/grpc/proto"
	"google.golang.org/grpc"
	"log"
	"time"
)

const (
	address = "localhost:50051"
	name    = "Astron"
)

func main() {
	conn, err := grpc.Dial(address,
		grpc.WithInsecure(),
		grpc.WithBlock())

	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewApiClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	reply, err := c.GetHello(ctx, &pb.Request{Name: name})

	if err != nil {
		log.Fatalf("GetHello error: %v", err)
	}
	log.Printf("Person: %v", reply)
}
