package main

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	pb "shark/pkg/proto"
)

func main() {
	conn, _ := grpc.Dial("127.0.0.1:16000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	cc := pb.NewSigServiceClient(conn)
	req := &pb.SigReq{
		Uin:      884322372,
		Password: []byte("youfei877"),
		DeviceId: []byte("homo-dddd-ds88"),
	}
	rsp, err := cc.Register(context.Background(), req)
	log.Printf("rsp: %v, err: %v", rsp, err)
	rsp, err = cc.SigIn(context.Background(), req)
	log.Printf("rsp %v,err %v", rsp, err)

	checkReq := &pb.CheckReq{
		Uin:   884322372,
		Token: rsp.GetToken(),
	}
	rsp, err = cc.Check(context.Background(), checkReq)
	log.Printf("Check rsp %v,err %v", rsp.Code, err)

	rsp, err = cc.SigOut(context.Background(), checkReq)
	log.Printf("rsp %v,err %v", rsp, err)
}
