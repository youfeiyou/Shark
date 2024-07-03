package main

import (
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"shark/config"
	"shark/internal/conn"
	pb "shark/pkg/proto/conn"
)

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", config.Conf.ConnServicePort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterSigServiceServer(s, &conn.PushServer{})
	log.Printf("server listening at %v", lis.Addr())
	go func() {
		if err := conn.Start(); err != nil {
			log.Fatalf("failed start serve: %v", err)
		}
	}()
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
