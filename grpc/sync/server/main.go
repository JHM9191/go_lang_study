package main

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"net"
	"time"

	pb "go_lang_study/grpc/sync/proto"
)

type APIServer struct {
	pb.UnimplementedApiServer
}

func (s *APIServer) GetHello(ctx context.Context, in *pb.Request) (*pb.Reply, error) {

	log.Printf("Received: %v", in.GetName())
	time.Sleep(5 * time.Second)
	ctx.Done()
	return &pb.Reply{Message: "Hello " + in.GetName()}, nil
}

func main() {
	lis, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatalf("failed to listen %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterApiServer(grpcServer, &APIServer{})

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
