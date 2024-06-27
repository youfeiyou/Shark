package relationship

import (
	"context"
	"github.com/pkg/errors"
	"log"
	db "shark/pkg/db/mysql"
	pb "shark/pkg/proto/relation"
	"shark/pkg/rpc"
)

type Server struct {
	pb.UnimplementedRelationServiceServer
}

func (Server) GetAllFriends(ctx context.Context, req *pb.Req) (*pb.Rsp, error) {
	rsp := &pb.Rsp{}
	if err := checkParams(req); err != nil {
		return rsp, err
	}
	r, err := db.RelationshipTableAPI.GetAllFriends(req.GetUin())
	if err != nil {
		log.Printf("GetAllFriends fail %v", err)
		return rsp, err
	}
	for _, v := range r {
		rsp.Info = append(rsp.Info, &pb.Relation{
			Uin:    v.FriendUin,
			Status: (pb.Status)(v.Status),
		})
	}
	return rsp, nil
}

func (Server) UpdateFriend(ctx context.Context, req *pb.Req) (*pb.Rsp, error) {
	rsp := &pb.Rsp{}
	if err := checkParams(req); err != nil {
		return rsp, err
	}
	r := &db.Relationship{
		Uin:       req.GetUin(),
		FriendUin: req.GetFriendUin(),
	}
	switch req.Type {
	case pb.ReqType_ADD_TYPE:
		r.Status = uint8(pb.Status_FRIEND)
		break
	case pb.ReqType_DELETE_TYPE:
		r.Status = uint8(pb.Status_DELETE)
		break
	case pb.ReqType_BLACK_TYPE:
		r.Status = uint8(pb.Status_BLACK)
		break
	}
	if req.GetType() == pb.ReqType_ADD_TYPE {
		if err := db.RelationshipTableAPI.Add(r); err != nil {
			log.Printf("Update fail %v", err)
			return rsp, err
		}
	} else {
		if err := db.RelationshipTableAPI.Update(r); err != nil {
			log.Printf("Update fail %v", err)
			return rsp, err
		}
	}
	return rsp, nil
}

func (Server) CheckFriend(ctx context.Context, req *pb.Req) (*pb.Rsp, error) {
	rsp := &pb.Rsp{}
	if err := checkParams(req); err != nil {
		return rsp, err
	}
	r, err := db.RelationshipTableAPI.Get(req.GetUin(), []uint64{req.GetFriendUin()})
	if err != nil || len(r) != 1 {
		log.Printf("GetAllFriends fail %v", err)
		return rsp, err
	}
	rsp.Info = append(rsp.Info, &pb.Relation{
		Uin:    r[0].FriendUin,
		Status: (pb.Status)(r[0].Status),
	})
	return rsp, nil
}

func checkParams(req *pb.Req) error {
	if req.GetUin() < 1000 {
		log.Printf("invalid parms uinlists %v", req)
		return errors.New("invalid params")
	}
	if req.GetType() <= pb.ReqType_INVALID_TYPE || req.GetType() > pb.ReqType_CHECK_TYPE {
		log.Printf("invalid parms type%v", req)
		return errors.New("invalid params type")
	}
	uins := []uint64{req.GetUin(), req.GetFriendUin()}
	dic := map[uint64]bool{}
	members, err := rpc.GetMembersInfo(context.Background(), uins)
	if err != nil {
		return err
	}
	for _, m := range members {
		log.Printf("m: %+v", m)
		dic[m.GetUin()] = true
	}
	if _, ok := dic[req.GetUin()]; !ok {
		log.Printf("invalid uin %v", req)
		return errors.New("invalid uin")
	}
	if _, ok := dic[req.GetFriendUin()]; req.GetType() == pb.ReqType_ADD_TYPE && ok {
		log.Printf("invalid friend uin,alredy exists %v", req)
		return errors.New("invalid friend uin")
	}
	if _, ok := dic[req.GetFriendUin()]; req.GetType() != pb.ReqType_ADD_TYPE && !ok {
		log.Printf("invalid friend uin ,ot exist%v", req)
		return errors.New("invalid friend uin")
	}
	return nil
}
