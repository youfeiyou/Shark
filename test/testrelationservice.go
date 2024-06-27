package main

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	pb "shark/pkg/proto/relation"
)

func main() {
	conn, _ := grpc.Dial("127.0.0.1:16003", grpc.WithTransportCredentials(insecure.NewCredentials()))
	cc := pb.NewRelationServiceClient(conn)
	for i := 1000; i < 2000; i = i + 50 {
		req := &pb.Req{
			Uin:       12345,
			FriendUin: uint64(i),
			Type:      pb.ReqType_ADD_TYPE,
		}
		rsp, err := cc.UpdateFriend(context.Background(), req)
		log.Printf("Register rsp %v err %v", rsp, err)
	}
	rsp, _ := cc.GetAllFriends(context.Background(), &pb.Req{Uin: 12345, Type: pb.ReqType_CHECK_TYPE})
	log.Println(len(rsp.GetInfo()))
	for _, v := range rsp.GetInfo() {
		log.Println(v.Uin, " ", v.Status)
	}
	for i := 1000; i < 2000; i = i + 50 {
		_, _ = cc.UpdateFriend(context.Background(), &pb.Req{Uin: 12345, FriendUin: uint64(i), Type: pb.ReqType_BLACK_TYPE})
	}
}
