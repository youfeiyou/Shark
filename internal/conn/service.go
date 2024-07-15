package conn

import (
	"golang.org/x/net/context"
	"log"
	pb "shark/pkg/proto/conn"
	"shark/pkg/proto/msg"
)

type PushServer struct {
	pb.UnimplementedConnServiceServer
}

func (*PushServer) Push(ctx context.Context, req *pb.PushReq) (*pb.PushRsp, error) {
	log.Printf("req %v", req)
	if len(req.GetPushMessages()) == 0 {
		return &pb.PushRsp{}, nil
	}
	rsp := &pb.PushRsp{}
	for _, v := range req.GetPushMessages() {
		anyct, ok := cm.connections.Load(v.GetUin())
		if !ok {
			rsp.FailedUin = append(rsp.FailedUin, v.GetUin())
			continue
		}
		ct, _ := anyct.(*connectionContext)
		go func(msg *msg.Message) {
			n, _ := ct.conn.Write(msg.GetContent())
			log.Printf("send %v bytes to uin", n)
		}(v)
	}
	return &pb.PushRsp{}, nil
}
