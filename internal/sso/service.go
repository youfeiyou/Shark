package sso

import (
	"bytes"
	"context"
	"github.com/pkg/errors"
	"log"
	"shark/pkg/db"
	pb "shark/pkg/proto"
	"shark/pkg/util"
	"time"
)

const (
	TokenLength    = 16
	TokenLimitTime = 7 * 24 * time.Hour
	MinUin         = 1000
	MinPasswordLen = 6
)

type Server struct {
	pb.UnimplementedSigServiceServer
}

func (s *Server) SigIn(ctx context.Context, req *pb.SigReq) (*pb.Response, error) {
	log.Printf("SigIn, req: %v %v", req.GetUin(), req.GetDeviceId())
	rsp := &pb.Response{}
	info, err := db.GetMemberSig(req.GetUin())
	if err != nil {
		log.Printf("et Member  info fail: %+v", err)
		rsp.Code = util.GetMemberSigInfoFail
		rsp.Result = []byte("get member sig info fail")
		return rsp, nil
	}
	if !bytes.Equal(info.GetPassword(), util.Md5(req.GetPassword())) {
		log.Printf("invalid password")
		rsp.Code = util.InvalidPassword
		rsp.Result = []byte("invalid password")
		return rsp, nil
	}
	info.Token = util.RandString(TokenLength)
	info.ExpiredTime = uint64(time.Now().Add(TokenLimitTime).Unix())
	if err := db.UpdateMemberSig(info); err != nil {
		log.Printf("update Member  info fail: %+v", err)
		return nil, err
	}
	rsp.Code = util.Success
	rsp.Result = []byte("success")
	rsp.Token = info.GetToken()
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

func (s *Server) Register(ctx context.Context, req *pb.SigReq) (*pb.Response, error) {
	log.Printf("Register req: %+v", req)
	if req.GetUin() < MinUin || len(req.GetPassword()) < MinPasswordLen {
		log.Printf("invalid uin num,%v", req.GetUin())
		return &pb.Response{Code: util.InvalidParms, Result: []byte("invalid parms")}, nil
	}
	memberInfo, err := db.GetMemberSig(req.GetUin())
	if err != nil {
		log.Printf("GetMemberSig fail: %+v", req.GetUin())
		return nil, errors.New("request db fail")
	}
	if memberInfo != nil && memberInfo.GetUin() == req.GetUin() {
		log.Printf("uin already exists: %+v", req.GetUin())
		return &pb.Response{Code: util.UinAlreadyExists, Result: []byte("uin already exists")}, nil
	}
	info := &pb.MemberSigInfo{
		Uin:      req.GetUin(),
		Password: util.Md5(req.GetPassword()),
	}
	log.Printf("info: %+v", info)
	if err := db.UpdateMemberSig(info); err != nil {
		log.Printf("update db fail: %+v", err)
		return nil, errors.New("update db fail")
	}
	return &pb.Response{Code: util.Success}, nil
}
