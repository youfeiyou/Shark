package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	pb "shark/pkg/proto"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:16000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
		return
	}
	defer conn.Close()
	cc := pb.NewSigServiceClient(conn)
	req := pb.SigReq{
		Uin: 12345678,
	}
	rsp, err := cc.SigIn(context.Background(), &req)
	fmt.Println(rsp, err)
}
