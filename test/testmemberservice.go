package main

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	pb "shark/pkg/proto/member"
)

func main() {
	conn, _ := grpc.Dial("127.0.0.1:16002", grpc.WithTransportCredentials(insecure.NewCredentials()))
	cc := pb.NewServiceClient(conn)
	req := &pb.Req{
		Member: []*pb.Member{

			{
				Uin:       8173922,
				MemberId:  "sfawsa2175",
				Password:  []byte("SAAIUS15474121"),
				Signature: []byte("ssaaasa"),
				Email:     "mcsdaklsas@1212.com",
				Phone:     23121121,
			},
		},
		Filter: &pb.MemberFilter{
			Signature: 1,
			Email:     1,
			Phone:     1,
		},
	}
	rsp, err := cc.Register(context.Background(), req)
	log.Printf("Register rsp %v err %v", rsp, err)
}
