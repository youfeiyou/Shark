package member

import (
	"context"
	"github.com/pkg/errors"
	"log"
	db "shark/pkg/db/mysql"
	pb "shark/pkg/proto/member"
	"shark/pkg/util"
	"sync"
)

const PasswordLength = 6

type Server struct {
	pb.UnimplementedServiceServer
}

func (s *Server) Register(ctx context.Context, req *pb.Req) (*pb.Rsp, error) {
	var err error
	if len(req.GetMember()) != 1 {
		log.Printf("Register invalid parms")
		return &pb.Rsp{}, nil
	}
	member := req.GetMember()[0]
	if len(member.GetPassword()) < PasswordLength {
		log.Printf("Register invalid parms")
		return &pb.Rsp{}, nil
	}
	filter := &pb.MemberFilter{
		Uin:           1,
		Password:      1,
		MemberId:      1,
		Name:          req.GetFilter().GetName(),
		Profile:       req.GetFilter().GetProfile(),
		Signature:     req.GetFilter().GetSignature(),
		Email:         req.GetFilter().GetEmail(),
		MessageNotify: 1,
	}
	dbmember := ConvertToDB(member, filter)
	dbmember.PassWord = util.Md5(dbmember.PassWord)
	dbmember.Uin, err = util.GenSeq(util.UINSEQ)
	member.Uin = dbmember.Uin
	if err != nil {
		return &pb.Rsp{}, err
	}
	log.Printf("dbmenber %v", dbmember)
	if err := db.MemberAPI.AddMembers([]*db.Member{dbmember}); err != nil {
		log.Printf("Register fail,%v", err)
		return &pb.Rsp{}, err
	}
	go func() {
		_ = s.addMemberID(member.GetMemberId(), member.GetUin())
	}()
	return &pb.Rsp{Member: req.GetMember()}, nil
}

func (*Server) GetMember(ctx context.Context, req *pb.Req) (*pb.Rsp, error) {
	rsp := &pb.Rsp{}
	if len(req.GetMember()) == 0 {
		log.Printf("Register invalid parms")
		return rsp, nil
	}
	uins := make([]uint64, len(req.GetMember()))
	for _, v := range req.GetMember() {
		uins = append(uins, v.GetUin())
	}
	dbmembers, err := db.MemberAPI.GetMemberInfo(uins)
	if err != nil {
		log.Printf("Register GetMemberInfo fail")
		return rsp, nil
	}
	req.Filter.Password = 0
	for _, v := range dbmembers {
		log.Printf("members :%+v", v)
		rsp.Member = append(rsp.Member, ConvertToPb(v, req.GetFilter()))
	}
	return rsp, nil
}

func (s *Server) UpdateMember(ctx context.Context, req *pb.Req) (*pb.Rsp, error) {
	rsp := &pb.Rsp{}
	if err := s.removeNotExistsUins(ctx, req, rsp); err != nil {
		return rsp, err
	}
	filter := limitFields(req.GetFilter())
	members := req.GetMember()
	wg := &sync.WaitGroup{}
	for _, v := range members {
		member := v
		dbmember := ConvertToDB(member, filter)
		wg.Add(1)
		go func(*db.Member, *sync.WaitGroup) {
			defer wg.Done()
			if err := db.MemberAPI.UpdateMemberFields(dbmember); err != nil {
				log.Printf("UpdateMemberFields fail %v", dbmember)
				rsp.Member = append(rsp.Member, &pb.Member{})
			}
			rsp.Member = append(rsp.Member, member)
		}(dbmember, wg)
	}
	wg.Wait()
	return rsp, nil
}

func limitFields(filter *pb.MemberFilter) *pb.MemberFilter {
	return &pb.MemberFilter{
		Name:          filter.GetName(),
		Profile:       filter.GetProfile(),
		Email:         filter.GetEmail(),
		Signature:     filter.GetSignature(),
		Phone:         filter.GetPhone(),
		MessageNotify: filter.GetMessageNotify(),
	}
}

func (s *Server) removeNotExistsUins(ctx context.Context, req *pb.Req, rsp *pb.Rsp) error {
	var uins []uint64
	dic := map[uint64]struct{}{}
	for _, m := range req.GetMember() {
		uins = append(uins, m.GetUin())
	}
	dbmembers, err := db.MemberAPI.GetMemberInfo(uins)
	if err != nil {
		log.Printf("Register GetMemberInfo fail")
		return err
	}
	for _, v := range dbmembers {
		dic[v.Uin] = struct{}{}
	}
	for _, m := range req.GetMember() {
		if _, ok := dic[m.GetUin()]; !ok {
			rsp.Member = append(rsp.Member, &pb.Member{})
		}
	}
	return nil
}

func (s *Server) checkMemberID(memberID string) error {
	info := db.MemberIDAPI.Get(memberID)
	if info == nil || info.Uin > 0 {
		return errors.New("checkMemberID fail")
	}
	return nil
}

func (s *Server) addMemberID(memberID string, uin uint64) error {
	err := db.MemberIDAPI.Add(&db.MemberID{
		Uin:      uin,
		MemberID: memberID,
	})
	if err != nil {
		log.Printf("addMemberID fail %v", err)
		return err
	}
	return nil
}
