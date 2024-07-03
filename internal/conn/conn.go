package conn

import (
	"fmt"
	"github.com/panjf2000/gnet/v2"
	"golang.org/x/sync/syncmap"
	"google.golang.org/protobuf/proto"
	"log"
	"net"
	"shark/config"
	db "shark/pkg/db/redis"
	pb "shark/pkg/proto/conn"
	"strings"
	"time"
)

const (
	maxTimeoutCount = 5
	CODEC           = "codec"
)

var (
	codec        = &Codec{}
	addr, _      = getLocalHost()
	connectedCnt uint64
	cm           = &ConnManage{}
)

type ConnManage struct {
	gnet.EventHandler
	connections syncmap.Map
}

type connectionContext struct {
	uin         uint64
	lastContact uint64
	count       int8
	conn        gnet.Conn
}

func (conns *ConnManage) OnTraffic(c gnet.Conn) (action gnet.Action) {
	data, err := codec.Decode(c)
	if err == ErrIncompletePacket {
		return
	}
	if err != nil {
		log.Printf("invalid packet: %v", err)
		return gnet.Close
	}
	hb := &pb.HeatBeat{}
	if err := proto.Unmarshal(data, hb); err != nil {
		log.Printf("OnTraffic Unmarshal fail %v", err)
		return
	}
	_, ok := conns.connections.Load(hb.GetUin())
	if !ok {
		ct := &connectionContext{
			uin:         hb.GetUin(),
			lastContact: uint64(time.Now().Unix()),
			count:       0,
			conn:        c,
		}
		conns.connections.Store(hb.GetUin(), ct)
		c.SetContext(ct)
		if err := db.StatAPI.UpdateMemberStat(hb.GetUin(), true, addr); err != nil {
			log.Printf("db.StatAPI.UpdateMemberStat fail %+v", err)
			return
		}
	} else {
		v, _ := conns.connections.Load(hb.GetUin())
		ct := v.(*connectionContext)
		ct.lastContact = uint64(time.Now().Unix())
		ct.count = 0
	}
	// to do: update user status
	return
}

func (conns *ConnManage) OnClose(c gnet.Conn, err error) (action gnet.Action) {
	if err != nil {
		log.Printf("error occurred on connection=%s, %v\n", c.RemoteAddr().String(), err)
	}
	ct := c.Context().(*connectionContext)
	if ct != nil {
		conns.connections.Delete(ct.uin)
		if err := db.StatAPI.DelMemberStat(ct.uin); err != nil {
			return
		}
	}
	return
}

func (conns *ConnManage) OnOpen(c gnet.Conn) (out []byte, action gnet.Action) {
	connectedCnt++
	log.Printf("new uin connect %v %v", c.RemoteAddr(), connectedCnt)
	return
}

func (conns *ConnManage) OnTick() (delay time.Duration, action gnet.Action) {
	conns.connections.Range(func(k, v any) bool {
		uin, _ := k.(uint64)
		ct, _ := v.(*connectionContext)
		now := uint64(time.Now().Unix())
		if now-ct.lastContact > uint64(10*time.Second) {
			ct.count++
			log.Printf("%v heartbeat time out", uin)
		}
		if ct.count > maxTimeoutCount {
			log.Printf("remove %v conn", uin)
			_ = ct.conn.Close()
			conns.connections.Delete(uin)
		}
		return true
	})
	return 10 * time.Second, gnet.None
}

func Start() error {
	addr := fmt.Sprintf("tcp://:%d", config.Conf.ConnListenPort)
	if err := gnet.Run(cm, addr, gnet.WithMulticore(true)); err != nil {
		log.Printf("conn start fail %v", err)
		return err
	}
	return nil
}

func getLocalHost() (string, error) {
	conn, err := net.Dial("udp", "1.0.0.1:12345")
	if err != nil {
		return "", err
	}
	localAddr := conn.LocalAddr().String()
	colonIdx := strings.LastIndex(localAddr, ":")
	if colonIdx > 0 {
		return localAddr[:colonIdx], nil
	}
	return localAddr, nil
}
