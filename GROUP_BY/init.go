package GROUP_BY

import "github.com/go-ini/ini"

type ConfigList struct {
	Port     string
	Host     string
	User     string
	Password string
	DbName   string
}

var Config ConfigList
var DataSource string

func init() {
	cfg, _ := ini.Load("config.ini")
	Config = ConfigList{
		Port:     cfg.Section("web").Key("port").String(),
		Host:     cfg.Section("db").Key("host").String(),
		User:     cfg.Section("db").Key("user").String(),
		Password: cfg.Section("db").Key("password").String(),
		DbName:   cfg.Section("db").Key("dbname").String(),
	}
	host := "host=" + Config.Host
	port := " port=" + Config.Port
	user := " user=" + Config.User
	password := " password=" + Config.Password
	dbname := " dbname=" + Config.DbName
	DataSource = host + port + user + password + dbname + " sslmode=disable"
}
