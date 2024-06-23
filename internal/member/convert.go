package member

import (
	db "shark/pkg/db/mysql"
	pb "shark/pkg/proto/member"
)

func ConvertToDB(member *pb.Member, filter *pb.MemberFilter) *db.Member {
	dbmember := &db.Member{}
	if filter.GetUin() != 0 {
		dbmember.Uin = member.GetUin()
	}
	if filter.GetMemberId() != 0 {
		dbmember.MemberID = member.GetMemberId()
	}
	if filter.GetPassword() != 0 {
		dbmember.PassWord = member.GetPassword()
	}
	if filter.GetName() != 0 {
		dbmember.Name = member.GetName()
	}
	if filter.GetProfile() != 0 {
		dbmember.Profile = member.GetProfile()
	}
	if filter.GetSignature() != 0 {
		dbmember.Signature = member.GetSignature()
	}
	if filter.GetPhone() != 0 {
		dbmember.Phone = member.GetPhone()
	}
	if filter.GetEmail() != 0 {
		dbmember.Email = member.GetEmail()
	}
	if filter.GetMessageNotify() != 0 {
		dbmember.MessageNotify = (uint8)(member.GetMessageNotify())
	}
	return dbmember
}

func ConvertToPb(dbmember *db.Member, filter *pb.MemberFilter) *pb.Member {
	member := &pb.Member{}
	if filter.GetUin() != 0 {
		member.Uin = dbmember.Uin
	}
	if filter.GetMemberId() != 0 {
		member.MemberId = dbmember.MemberID
	}
	if filter.GetPassword() != 0 {
		member.Password = dbmember.PassWord
	}
	if filter.GetName() != 0 {
		member.Name = dbmember.Name
	}
	if filter.GetProfile() != 0 {
		member.Profile = dbmember.Profile
	}
	if filter.GetSignature() != 0 {
		member.Signature = dbmember.Signature
	}
	if filter.GetPhone() != 0 {
		member.Phone = dbmember.Phone
	}
	if filter.GetEmail() != 0 {
		member.Email = dbmember.Email
	}
	if filter.GetMessageNotify() != 0 {
		member.MessageNotify = uint32(dbmember.MessageNotify)
	}
	if filter.GetCreateTime() != 0 {
		member.CreateTime = uint64(dbmember.CreateTime.Unix())
	}
	if filter.GetModifyTime() != 0 {
		member.ModifyTime = uint64(dbmember.LastUpdateTime.Unix())
	}
	if filter.GetIsDeleted() != 0 {
		member.IsDeleted = dbmember.IsDeleted
	}
	return member
}
