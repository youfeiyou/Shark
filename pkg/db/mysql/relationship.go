package db

import (
	"gorm.io/gorm"
	"gorm.io/sharding"
	"log"
	"strconv"
	"time"
)

type Relationship struct {
	Uin            uint64    `gorm:";column:uin;"`
	FriendUin      uint64    `gorm:";column:friend_uin;"`
	Status         uint8     `gorm:";column:status;"`
	CreateTime     time.Time `gorm:"column:create_time;"`
	LastUpdateTime time.Time `gorm:"column:last_modify_time;"`
}

type RelationshipTable struct {
	db *gorm.DB
}

var RelationshipTableAPI RelationshipTable

func init() {
	var err error
	RelationshipTableAPI.db, err = InitMysql("shark")
	err = RelationshipTableAPI.db.Use(sharding.Register(sharding.Config{
		ShardingKey:           "uin",
		NumberOfShards:        shardingNum,
		PrimaryKeyGenerator:   sharding.PKCustom,
		PrimaryKeyGeneratorFn: func(int64) int64 { return 0 },
		ShardingAlgorithm: func(uin any) (suffix string, err error) {
			index := uin.(uint64) % (uint64)(shardingNum)
			log.Printf("sharding index: %v", index)
			return "_" + strconv.FormatUint(index, 10), nil
		},
	}, "relationship"))
	if err != nil {
		log.Printf("init relationship db api fail %v", err)
		panic("init relationship db api fail")
	}
	if err != nil {
		panic("init relationship db api fail")
	}
}

func (Relationship) TableName() string {
	return "relationship"
}

type RelationshipAPI interface {
	Add(uin uint64, friend uint64) error
	get(uin uint64, friends []uint64) (bool, error)
	GetAllFriends(uin uint64) ([]*Relationship, error)
	Update(relationship Relationship) error
}

func (m *RelationshipTable) Add(relationship *Relationship) error {
	result := m.db.Omit("CreateTime", "LastUpdateTime").Create(relationship)
	if result.Error != nil {
		log.Printf("RelationshipTable add fail %+v", result)
	}
	return nil
}

func (m *RelationshipTable) Get(uin uint64, friends []uint64) ([]*Relationship, error) {
	var r []*Relationship
	result := m.db.Model(&Relationship{Uin: uin}).Where("friend_uin in?", friends).Find(&r)
	if result.Error != nil {
		log.Printf("RelationshipTable get fail: %v", result.Error)
		return nil, nil
	}
	return r, nil
}

func (m *RelationshipTable) GetAllFriends(uin uint64) ([]*Relationship, error) {
	var r []*Relationship
	result := m.db.Model(&Relationship{Uin: uin}).Where("uin=?", uin).Find(&r)
	if result.Error != nil {
		log.Printf("RelationshipTable GetAllFriends fail: %v", result.Error)
		return nil, nil
	}
	log.Printf("result %+v", result)
	return r, nil
}

func (m *RelationshipTable) Update(relationship *Relationship) error {
	result := m.db.Model(&Relationship{Uin: relationship.Uin}).Where("uin=? and friend_uin=?", relationship.Uin, relationship.FriendUin).Updates(relationship)
	if result.Error != nil {
		log.Printf("RelationshipTable Update fail: %v", result.Error)
		return result.Error
	}
	log.Printf("result %+v", result)
	return nil
}
