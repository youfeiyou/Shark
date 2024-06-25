package rpc

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"shark/config"
	pb "shark/pkg/proto/member"
	"sync"
)

var (
	connInit          func()
	memberServiceConn *grpc.ClientConn
)

func init() {
	connInit = sync.OnceFunc(func() {
		memberServiceConn, _ = grpc.Dial(config.Conf.MemberServiceAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	})
}

func GetMemberInfo(ctx context.Context, uin uint64) (*pb.Member, error) {
	connInit()
	cc := pb.NewServiceClient(memberServiceConn)
	req := &pb.Req{
		Filter: &pb.MemberFilter{
			Password: 1,
		},
	}
	rsp, err := cc.GetMember(ctx, req)
	if err != nil {
		log.Printf("rpc GetMemberInfo fail uin: %v", uin)
		return nil, err
	}
	if len(rsp.GetMember()) == 0 {
		log.Printf("no uin info uin: %v", uin)
		return nil, nil
	}
	return rsp.GetMember()[0], nil
}
