package db

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"shark/config"
)

func InitMysql(dbname string) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", config.Conf.MysqlUser, config.Conf.MysqlPass, config.Conf.MysqlAddr, dbname)
	DB, err := gorm.Open(mysql.New(mysql.Config{
		DSN: dsn,
	}), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
		return nil, err
	}
	return DB, nil
}
