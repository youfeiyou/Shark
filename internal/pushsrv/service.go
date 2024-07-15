package pushsrv

import (
	"context"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"log"
	pb "shark/pkg/proto/pushsrv"
	"sync"
)

type Server struct {
	pb.UnsafePushSrvServiceServer
}

var stat sync.Map

func (s *Server) Push(ctx context.Context, req *pb.Req, opts ...grpc.CallOption) (*pb.Rsp, error) {
	rsp := &pb.Rsp{}
	if err := s.checkParams(req); err != nil {
		return nil, err
	}
	for _, v := range req.GetUins() {
		if st, ok := stat.Load(v); ok {
			log.Printf("cache stat %v", st)
		} else {
			// to do: check uin stat
		}
	}
	return rsp, nil
}

func (s *Server) checkParams(req *pb.Req) error {
	if len(req.GetUins()) == 0 {
		log.Printf("invalid uin num")
		return errors.New("invalid uin num")
	}
	if len(req.GetSessionId()) == 0 {
		log.Printf("invalid session")
		return errors.New("invalid session")
	}
	return nil
}
