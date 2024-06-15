package main

import (
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"shark/internal/file"
	pb "shark/pkg/proto/file"
)

const port = 16001

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterSigServiceServer(s, &file.Server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
