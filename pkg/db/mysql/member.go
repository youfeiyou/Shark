package db

import (
	"gorm.io/gorm"
	"log"
	"time"
)

//uin bigint not NULL,
//member_id bigint,
//password varchar(16),
//name varchar(10),
//profile varchar(16),
//sigure varchar(32),
//email varchar(32),
//phone bigint,
//message_notify TINYINT,
//create_time datetime NOT NULL DEFAULT CURRENT_TIMESTAMP(),
//last_modify_time datetime NOT NULL DEFAULT CURRENT_TIMESTAMP() ON UPDATE CURRENT_TIMESTAMP(),
//PRIMARY KEY(uin)

type Member struct {
	Uin            uint64    `gorm:"primary_key;column:uin;"`
	MemberID       string    `gorm:"column:member_id;"`
	PassWord       []byte    `gorm:"column:password;"`
	Name           string    `gorm:"column:name;"`
	Profile        []byte    `gorm:"column:profile;"`
	Signature      []byte    `gorm:"column:signature;"`
	Email          string    `gorm:"column:email;"`
	Phone          uint64    `gorm:"column:phone;"`
	MessageNotify  uint8     `gorm:"column:message_notify;"`
	CreateTime     time.Time `gorm:"column:create_time;"`
	LastUpdateTime time.Time `gorm:"column:last_modify_time;"`
	IsDeleted      bool      `gorm:"column:is_deleted;"`
}

type MemberDB interface {
	AddMembers([]*Member) error
	DelMember([]uint64) error
	UpdateMemberName(uint64, string) error
	GetMemberInfo([]uint64) ([]Member, error)
	UpdateMemberPass(id uint64, pass []byte) error
}

// MemberTable mysql 存储
type MemberTable struct {
	db *gorm.DB
}

var MemberAPI MemberTable

func init() {
	var err error
	MemberAPI.db, err = InitMysql("shark")

	if err != nil {
		panic("init member db api fail")
	}
}

func (Member) TableName() string {
	return "memberinfo"
}

func (m *MemberTable) AddMembers(members []*Member) error {
	result := m.db.Omit("CreateTime", "LastUpdateTime").Create(members)
	if result.Error != nil {
		log.Printf("AddMembers fail %+v", result)
	}
	return nil
}

func (m *MemberTable) DelMember(ids []uint64) error {
	result := m.db.Model(&Member{}).Where("uin IN ?", ids).Update("is_deleted", 1)
	if result.Error != nil {
		log.Printf("delMember fail %+v", result)
		return result.Error
	}
	return nil
}

func (m *MemberTable) UpdateMemberName(uin uint64, name string) error {
	result := m.db.Model(&Member{Uin: uin}).Select("Name").Update("name", name)
	if result.Error != nil {
		log.Printf("UpdateMemberName fail %+v", result)
		return result.Error
	}
	return nil
}

func (m *MemberTable) UpdateMemberPass(id uint64, pass []byte) error {
	result := m.db.Model(&Member{Uin: id}).Select("password").Update("password", pass)
	if result.Error != nil {
		log.Printf("UpdateMemberPass fail %+v", result)
		return result.Error
	}
	return nil
}

func (m *MemberTable) GetMemberInfo(idx []uint64) ([]*Member, error) {
	var info []*Member
	result := m.db.Where("uin IN ?", idx).Find(&info)
	if result.Error != nil {
		log.Printf("GetMemberInfo fail %+v", result)
		return nil, result.Error
	}
	return info, nil
}
