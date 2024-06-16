package main

import (
	"log"
	db "shark/pkg/db/mysql"
)

func main() {
	members := []*db.Member{
		{
			Uin:      123456,
			PassWord: []byte("ZAQWSX"),
			MemberID: "ydsyyyyyy",
		},
		{
			Uin:      893456,
			MemberID: "yyyyyyy",
			PassWord: []byte("ZAQWSX"),
		},
		{
			Uin:       3121928,
			Signature: []byte("dssddsd"),
			MemberID:  "yydsyyyyy",
			Phone:     212111,
			PassWord:  []byte("ZAQWSX"),
		},
	}
	err := db.MemberAPI.AddMembers(members)
	log.Printf("info %+v, err: %v", members, err)
	info, _ := db.MemberAPI.GetMemberInfo([]uint64{123456, 893456, 3121928})
	for _, v := range info {
		log.Printf(" %+v", v)
	}
	db.MemberAPI.UpdateMemberName(123456, "smdb")
	db.MemberAPI.UpdateMemberPass(123456, []byte("xcxcxcxxc"))
	db.MemberAPI.DelMember([]uint64{123456, 893456, 3121928})

}
