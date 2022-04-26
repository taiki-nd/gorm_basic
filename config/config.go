package config

import (
	"log"
	"os"

	"gopkg.in/ini.v1"
)

type ConfigList struct {
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
		DbUserDevelop:     cfg.Section("db_development").Key("db_user_develop").String(),
		DbPasswordDevelop: cfg.Section("db_development").Key("db_password_develop").String(),
	}
}
