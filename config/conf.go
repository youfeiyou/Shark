package config

type Config struct {
	MysqlAddr string
	MysqlPass string
	MysqlUser string
}

var Conf Config

func init() {
	Conf = Config{
		MysqlAddr: "39.108.64.37:3306",
		MysqlPass: "123456",
		MysqlUser: "root",
	}
}
