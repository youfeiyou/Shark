package main

import (
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"shark/internal/sso"
	pb "shark/pkg/proto"
)

const port = 16000

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterSigServiceServer(s, &sso.Server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
