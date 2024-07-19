package db

import (
	"context"
	"github.com/pkg/errors"
	"google.golang.org/protobuf/proto"
	"log"
	pb "shark/pkg/proto/groupmember"
	"strconv"
)

const (
	gmdb = "gm_"
	meta = "meta"
)

type groupMemberAPI struct {
}

var GroupMemeberAPI groupMemberAPI

func (groupMemberAPI) AddGroupMember(ctx context.Context, group uint64, member *pb.DbGroupMember) error {
	cli := NewRedisClient(RedisAddr)
	key := gmdb + strconv.FormatUint(group, 10)
	data, err := proto.Marshal(member)
	if err != nil {
		log.Printf("addGroupMember fail %+v", err)
		return err
	}
	mdata, err := cli.HMGet(key, meta)
	metadata := mdata[meta]
	if err != nil {
		log.Printf("get meta fail %v", metadata)
		return err
	}
	metaInfo := &pb.Meta{}
	_ = proto.Unmarshal([]byte(metadata), metaInfo)
	metaInfo.MemberNum++
	metabytes, _ := proto.Marshal(metaInfo)
	fields := map[string]interface{}{
		strconv.FormatUint(member.GetUin(), 10): data,
		meta:                                    metabytes,
	}
	if err := cli.HMSet(key, fields); err != nil {
		log.Printf("addGroupMember HMSet fail %+v", err)
		return err
	}
	log.Printf("AddGroupMember success %+v", member)
	return nil
}

func (groupMemberAPI) GetGroupMember(ctx context.Context, group uint64, uin uint64) (*pb.DbGroupMember, error) {
	cli := NewRedisClient(RedisAddr)
	key := gmdb + strconv.FormatUint(group, 10)
	ret, err := cli.HMGet(key, strconv.FormatUint(uin, 10))
	if err != nil {
		return nil, err
	}
	if len(ret) == 0 {
		log.Printf("no member uin %+v", uin)
		return nil, nil
	}
	if _, ok := ret[strconv.FormatUint(uin, 10)]; !ok {
		log.Printf("no member uin %+v", uin)
		return nil, errors.New("wrong data from db")
	}
	m := &pb.DbGroupMember{}
	err = proto.Unmarshal([]byte(strconv.FormatUint(uin, 10)), m)
	if err != nil {
		log.Printf("proto.Unmarshal fail %+v", uin)
		return nil, err
	}
	return m, nil
}

func (groupMemberAPI) UpdateGroupMember(ctx context.Context, group uint64, member *pb.DbGroupMember) error {
	cli := NewRedisClient(RedisAddr)
	key := gmdb + strconv.FormatUint(group, 10)
	data, err := proto.Marshal(member)
	if err != nil {
		log.Printf("UpdateGroupMember fail %+v", err)
		return err
	}
	fields := map[string]interface{}{
		strconv.FormatUint(member.GetUin(), 10): data,
	}
	if err := cli.HMSet(key, fields); err != nil {
		log.Printf("UpdateGroupMember HMSet fail %+v", err)
		return err
	}
	log.Printf("UpdateGroupMember success %+v", member)
	return nil
}

func (groupMemberAPI) GetAllGroupMember(ctx context.Context, group uint64) ([]*pb.DbGroupMember, error) {
	cli := NewRedisClient(RedisAddr)
	key := gmdb + strconv.FormatUint(group, 10)
	ret, err := cli.HMGet(key)
	if err != nil {
		log.Printf("GetAllGroupMember fail %+v", err)
		return nil, err
	}
	if len(ret) == 0 {
		log.Printf("no uin info")
		return nil, nil
	}
	ans := make([]*pb.DbGroupMember, len(ret))
	for k, v := range ret {
		if k == meta {
			continue
		}
		m := &pb.DbGroupMember{}
		if err := proto.Unmarshal([]byte(v), m); err != nil {
			log.Printf("proto.Unmarshal fail %+v", err)
			continue
		}
		ans = append(ans, m)
	}
	return ans, nil
}
