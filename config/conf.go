package config

type Config struct {
	MysqlAddr string
	MysqlPass string
	MysqlUser string

	MemberServiceAddr   string
	RelationServicePort int16
}

var Conf Config

func init() {
	Conf = Config{
		MysqlAddr: "39.108.64.37:3306",
		MysqlPass: "shark0803",
		MysqlUser: "root",

		MemberServiceAddr:   "127.0.0.1:16002",
		RelationServicePort: 16003,
	}
}
