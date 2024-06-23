package db

import (
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"gorm.io/sharding"
	"log"
	"shark/pkg/util"
	"strconv"
)

type MemberID struct {
	Uin      uint64 `gorm:"column:uin;type:bigint;"`
	MemberID string `gorm:"primary_key;column:member_id;type:varchar(10)"`
}

const (
	shardingNum = 2
)

type MemberIDTable struct {
	db *gorm.DB
}

var MemberIDAPI MemberIDTable

func init() {
	var err error
	MemberIDAPI.db, err = InitMysql("shark")
	err = MemberIDAPI.db.Use(sharding.Register(sharding.Config{
		ShardingKey:           "member_id",
		NumberOfShards:        shardingNum,
		PrimaryKeyGenerator:   sharding.PKCustom,
		PrimaryKeyGeneratorFn: func(int64) int64 { return 0 },
		ShardingAlgorithm: func(memberid any) (suffix string, err error) {
			str, ok := memberid.(string)
			if !ok {
				return "", errors.New("invalid memberid")
			}
			index := (uint64)(util.HashString(str) % shardingNum)
			log.Printf("sharding index: %v", index)
			return strconv.FormatUint(index, 10), nil
		},
	}, "member_ids"))
	if err != nil {
		log.Printf("init memberid db api fail %v", err)
		panic("init memberid db api fail")
	}
}

func (MemberID) TableName() string {
	return "member_ids"
}

func (m *MemberIDTable) Add(mid *MemberID) error {
	result := m.db.Create(mid)
	if result.Error != nil {
		log.Printf("add fail %+v", result)
	}
	return nil
}

func (m *MemberIDTable) Get(memberid string) *MemberID {
	mids := []*MemberID{}
	result := m.db.Model(&MemberID{MemberID: memberid}).Where("member_id=?", memberid).Find(&mids)
	if result.Error != nil {
		log.Printf("MemberIDTable get fail: %v", result.Error)
		return nil
	}
	return mids[0]
}
