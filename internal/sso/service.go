package sso

import (
	"bytes"
	"context"
	"github.com/pkg/errors"
	"log"
	db "shark/pkg/db/redis"
	"shark/pkg/rpc"

	pb "shark/pkg/proto"
	"shark/pkg/util"
	"time"
)

const (
	TokenLength    = 16
	TokenLimitTime = 7 * 24 * time.Hour
)

type Server struct {
	pb.UnimplementedSigServiceServer
}

func (s *Server) SigIn(ctx context.Context, req *pb.SigReq) (*pb.Response, error) {
	log.Printf("SigIn, req: %v %v", req.GetUin(), req.GetDeviceId())
	rsp := &pb.Response{}
	memberinfo, err := rpc.GetMemberInfo(ctx, req.GetUin())
	if err != nil {
		log.Printf("get member info fail %v", err)
		return rsp, nil
	}
	if err != nil {
		log.Printf("et Member  info fail: %+v", err)
		rsp.Code = util.GetMemberSigInfoFail
		rsp.Result = []byte("get member sig info fail")
		return rsp, nil
	}
	if !bytes.Equal(memberinfo.GetPassword(), util.Md5(req.GetPassword())) {
		log.Printf("invalid password")
		rsp.Code = util.InvalidPassword
		rsp.Result = []byte("invalid password")
		return rsp, nil
	}
	siginfo := &pb.MemberSigInfo{
		Token:       util.RandString(TokenLength),
		ExpiredTime: uint64(time.Now().Add(TokenLimitTime).Unix()),
		DeviceId:    req.GetDeviceId(),
	}
	if err := db.UpdateMemberSig(siginfo); err != nil {
		log.Printf("update Member  info fail: %+v", err)
		return nil, err
	}
	rsp.Code = util.Success
	rsp.Result = []byte("success")
	rsp.Token = siginfo.GetToken()
	return rsp, nil
}

func (s *Server) SigOut(ctx context.Context, req *pb.CheckReq) (*pb.Response, error) {
	log.Printf("SigOut, req: %v", req.GetUin())
	memberInfo, err := db.GetMemberSig(req.GetUin())
	if err != nil {
		log.Printf("GetMemberSig fail: %+v", req.GetUin())
		return nil, errors.New("request db fail")
	}
	if memberInfo == nil {
		log.Printf("parms error: %+v", req.GetUin())
		return &pb.Response{Code: util.InvalidParms, Result: []byte("parms error")}, nil
	}
	memberInfo.Token = []byte{}
	memberInfo.ExpiredTime = 0
	if err := db.UpdateMemberSig(memberInfo); err != nil {
		log.Printf("update db fail: %+v", err)
		return nil, errors.New("update db fail")
	}
	return &pb.Response{Code: util.Success}, nil
}

func (s *Server) Check(ctx context.Context, req *pb.CheckReq) (*pb.Response, error) {
	log.Printf("Check, req: %v", req.GetUin())
	memberInfo, err := db.GetMemberSig(req.GetUin())
	if err != nil {
		log.Printf("GetMemberSig fail: %+v", req.GetUin())
		return nil, errors.New("request db fail")
	}
	if !bytes.Equal(req.GetToken(), memberInfo.GetToken()) {
		log.Printf("invalid token: %+v", req.GetUin())
		return &pb.Response{Code: util.InvalidToken}, nil
	}
	if memberInfo.GetExpiredTime() <= uint64(time.Now().Unix()) {
		log.Printf("expired token: %+v", req.GetUin())
		return &pb.Response{Code: util.ExpiredToken}, nil
	}
	return &pb.Response{Code: util.Success}, nil
}
