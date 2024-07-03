package db

import (
	"fmt"
	"google.golang.org/protobuf/proto"
	"log"
	pb "shark/pkg/proto/conn"
	"strconv"
)

type MemberStatAPI interface {
	UpdateMemberStat(uin uint64, online bool, addr string) error
	GetMemberStat(uin uint64) (bool, string, error)
	DelMemberStat(uin uint64) error
}

type statRedisTble struct{}

const (
	shardingNum = 16
	tablePrefix = "MemberStat"
)

var StatAPI MemberStatAPI = &statRedisTble{}

func (s statRedisTble) UpdateMemberStat(uin uint64, online bool, addr string) error {
	cli := NewRedisClient(RedisAddr)
	stat, err := getStat(uin)
	if err != nil {
		return err
	}
	if stat == nil {
		stat = &pb.MemberStat{
			Uin: uin,
		}
		log.Printf("uin  exits")
	}
	stat.ConnAddr = addr
	stat.IsOnline = 0
	if online {
		stat.IsOnline = 1
	}
	data, _ := proto.Marshal(stat)
	fields := map[string]interface{}{
		strconv.FormatUint(uin, 10): data,
	}
	if err := cli.HMSet(genkey(uin), fields); err != nil {
		log.Printf("cli.HMSet fail %v", err)
		return err
	}
	return nil
}

func (s statRedisTble) GetMemberStat(uin uint64) (bool, string, error) {
	stat, err := getStat(uin)
	if err != nil {
		return false, "", err
	}
	if stat == nil {
		return false, "", err
	}
	online := false
	if stat.GetIsOnline() > 0 {
		online = true
	}
	return online, stat.GetConnAddr(), nil
}

func (s statRedisTble) DelMemberStat(uin uint64) error {
	cli := NewRedisClient(RedisAddr)
	if err := cli.HDel(genkey(uin), strconv.FormatUint(uin, 10)); err != nil {
		log.Printf("DelMemberStat fail %+v", err)
		return err
	}
	return nil
}

func getStat(uin uint64) (*pb.MemberStat, error) {
	cli := NewRedisClient(RedisAddr)
	result, err := cli.HMGet(genkey(uin), strconv.FormatUint(uin, 10))
	if err != nil {
		log.Printf("GetMemberStat fail %v", err)
		return nil, err
	}
	res, ok := result[strconv.FormatUint(uin, 10)]
	if !ok {
		log.Printf("not exists %d stat")
		return nil, nil
	}
	stat := &pb.MemberStat{}
	if err := proto.Unmarshal([]byte(res), stat); err != nil {
		log.Printf("GetMemberStat Unmarshal fail %v", err)
		return nil, err
	}
	return stat, nil
}

func genkey(uin uint64) string {
	return fmt.Sprintf("%s_%d", MemberSigTable, uin%shardingNum)
}
