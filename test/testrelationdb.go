package main

import (
	"log"
	db "shark/pkg/db/mysql"
)

func main() {

	for i := 1000; i < 1010; i++ {
		_ = db.RelationshipTableAPI.Add(&db.Relationship{
			Uin:       21210,
			FriendUin: uint64(i),
		})
	}

	fs, _ := db.RelationshipTableAPI.GetAllFriends(121211)
	for _, v := range fs {
		log.Println(v)
	}
	err := db.RelationshipTableAPI.Update(&db.Relationship{
		Uin:       121211,
		FriendUin: 1998,
		Status:    1,
	})
	log.Println(err)
	fss, err := db.RelationshipTableAPI.GetAllFriends(121211)
	log.Println(err)
	for _, v := range fss {
		log.Println(v)
	}
}
