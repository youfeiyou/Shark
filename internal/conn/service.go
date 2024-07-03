package conn

import (
	"golang.org/x/net/context"
	"log"
	pb "shark/pkg/proto/conn"
)

type PushServer struct {
	pb.UnimplementedSigServiceServer
}

func (*PushServer) Push(ctx context.Context, req *pb.PushReq) (*pb.PushRsp, error) {
	log.Printf("req %v", req)
	if len(req.GetPushMessages()) == 0 {
		return &pb.PushRsp{}, nil
	}
	for _, v := range req.GetPushMessages() {
		anyct, ok := cm.connections.Load(v.GetUin())
		if !ok {
			continue
		}
		ct, _ := anyct.(*connectionContext)
		go func(msg *pb.Message) {
			n, _ := ct.conn.Write(msg.GetContent())
			log.Printf("send %v bytes to uin", n)
		}(v)
	}
	return &pb.PushRsp{}, nil
}
