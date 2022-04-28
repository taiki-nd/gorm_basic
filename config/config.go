package config

import (
	"log"
	"os"

	"gopkg.in/ini.v1"
)

type ConfigList struct {
	DbSqlDevelop      string
	DbHostDevelop     string
	DbPortDevelop     string
	DbNameDevelop     string
	DbUserDevelop     string
	DbPasswordDevelop string
}

var Config ConfigList

func init() {
	cfg, err := ini.Load("config.ini")
	if err != nil {
		log.Panicf("failed to load config.ini: %v", err)
		os.Exit(1)
	}

	Config = ConfigList{
		DbSqlDevelop:      cfg.Section("db_development").Key("db_sql_develop").String(),
		DbHostDevelop:     cfg.Section("db_development").Key("db_host_develop").String(),
		DbPortDevelop:     cfg.Section("db_development").Key("db_port_develop").String(),
		DbNameDevelop:     cfg.Section("db_development").Key("db_name_develop").String(),
		DbUserDevelop:     cfg.Section("db_development").Key("db_user_develop").String(),
		DbPasswordDevelop: cfg.Section("db_development").Key("db_password_develop").String(),
	}
}
