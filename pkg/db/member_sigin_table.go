package db

import (
	"google.golang.org/protobuf/proto"
	"log"
	pb "shark/pkg/proto"
	"strconv"
)

const (
	MemberSigTable = "MemberSigTable"
)

func UpdateMemberSig(info *pb.MemberSigInfo) error {
	buf, err := proto.Marshal(info)
	if err != nil {
		log.Fatalf("UpdateMemberSigInfo Marshal error :%v", err)
		return err
	}
	cli := NewRedisClient("127.0.0.1:6379")
	err = cli.Hmset(MemberSigTable, map[string]interface{}{
		strconv.FormatUint(info.GetUin(), 10): buf,
	})
	if err != nil {
		log.Fatalf("UpdateMemberSigInfo write db error :%v", err)
		return err
	}
	return nil
}

func GetMemberSig(uin uint64) (*pb.MemberSigInfo, error) {
	cli := NewRedisClient("127.0.0.1:6379")
	uinstr := strconv.FormatUint(uin, 10)
	info, err := cli.Hmget(MemberSigTable, uinstr)
	if err != nil {
		log.Fatalf("GetMemberSig read db error :%v", err)
		return nil, err
	}
	ms := &pb.MemberSigInfo{}
	if err := proto.Unmarshal([]byte(info[uinstr]), ms); err != nil {
		log.Fatalf("GetMemberSig Unmarshal error :%v", err)
		return nil, err
	}
	return ms, nil
}
