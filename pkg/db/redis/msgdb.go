package db

import (
	"context"
	"fmt"
	"google.golang.org/protobuf/proto"
	"log"
	"shark/pkg/proto/msg"
	"strconv"
)

const (
	msgdb           = "msgdb_"
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

func (msgSEQPAI) genSessionSeq(ctx context.Context, isAtMsg, isPacketMsg bool) (*msg.SessionSeq, error) {
	seq := &msg.SessionSeq{}
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
