package db

import (
	"context"
	"fmt"
	"google.golang.org/protobuf/proto"
	"log"
	"reflect"
	"shark/pkg/proto/msg"
	"strconv"
)

const (
	msgdb           = "msgdb_"
	msgseqdb        = "msgseqdb_"
	sessiondb       = "sessiondb_"
	sessionmemberdb = "sessionmemberdb_"
)

type msgDBAPI struct {
}

type msgSEQPAI struct {
}
type sessionMemeberapi struct {
}

var (
	MsgDBAPI         msgDBAPI
	MsgSEQapi        msgSEQPAI
	SessionMemberAPI sessionMemeberapi
)

func (m msgDBAPI) AddPersonalMessage(ctx context.Context, rid uint64, message *msg.Message) error {
	log.Printf("AddPersonalMessage message %+v", message)
	sid := genPersonalSessionID(message.GetUin(), rid)
	return m.AddMessage(ctx, sid, message)
}

func (m msgDBAPI) AddGroupMessage(ctx context.Context, gid uint64, message *msg.Message) error {
	log.Printf("AddGroupMessage message %+v", message)
	sid := genGroupSessionID(gid)
	return m.AddMessage(ctx, sid, message)
}

func (msgDBAPI) AddMessage(ctx context.Context, sid string, message *msg.Message) error {
	cli := NewRedisClient(RedisAddr)
	id := strconv.FormatUint(message.GetSeq().GetMaxMsgSeq(), 10)
	data, err := proto.Marshal(message)
	if err != nil {
		log.Printf("proto.Marshal fail %v", err)
	}
	fields := map[string]interface{}{
		id: data,
	}
	if err := cli.HMSet(sid, fields); err != nil {
		log.Printf("cli.HMSet fail %v", err)
		return err
	}
	return nil
}

func (m msgDBAPI) GetGroupMessages(ctx context.Context, gid uint64, seq []uint64) (*msg.Message, error) {

	return nil, nil
}

func (m msgDBAPI) GetPersonalMessages(ctx context.Context, sendid, rid uint64, seq []uint64) (*msg.Message, error) {

	return nil, nil
}

func (msgSEQPAI) genSessionSeq(ctx context.Context, sid string, isAtMsg, isPacketMsg bool) (*msg.SessionSeq, error) {
	seq := &msg.SessionSeq{}
	cli := NewRedisClient(RedisAddr)
	argv := []interface{}{"MaxMsgSeq"}
	if isAtMsg {
		argv = append(argv, "MaxMsgAtSeq")
	}
	if isPacketMsg {
		argv = append(argv, "PacketSeq")
	}
	cmd := MsgSeqScript.Run(ctx, cli.cli, []string{sid}, argv...)
	if cmd.Err() != nil {
		return nil, cmd.Err()
	}
	s := cmd.Val().([]interface{})
	ele := reflect.ValueOf(seq).Elem()
	for i := 0; i < len(s); i += 2 {
		k := s[i].(string)
		v, _ := strconv.ParseUint(s[i+1].(string), 10, 64)
		ele.FieldByName(k).SetUint(v)
	}
	log.Printf("success gen msg seq %+v", seq)
	return seq, nil
}

func genPersonalSessionID(sid, rid uint64) string {
	a, b := sid, rid
	if sid > rid {
		a, b = b, a
	}
	return fmt.Sprintf("c_%d%d", a, b)
}

func genGroupSessionID(gid uint64) string {
	return fmt.Sprintf("g_%d", gid)
}
