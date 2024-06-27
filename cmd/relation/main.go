package main

import (
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"shark/config"
	"shark/internal/relationship"
	pb "shark/pkg/proto/relation"
)

const port = 16003

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", config.Conf.RelationServicePort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterRelationServiceServer(s, &relationship.Server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
