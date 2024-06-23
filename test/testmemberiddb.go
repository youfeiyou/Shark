package main

import (
	"log"
	db "shark/pkg/db/mysql"
	"strconv"
)

func main() {
	for i := 1000; i < 1100; i++ {
		err := db.MemberIDAPI.Add(&db.MemberID{
			Uin:      uint64(i),
			MemberID: "ghhds" + strconv.FormatInt(int64(i), 10),
		})
		log.Printf("err: %v", err)
		info := db.MemberIDAPI.Get("ghhds" + strconv.FormatInt(int64(i), 10))
		log.Println(info)
	}
}
