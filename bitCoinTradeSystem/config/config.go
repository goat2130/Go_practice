package config

import (
	"log"
	"os"
)

type ConfigList struct {
	ApiKey    string
	ApiSecret string
}

var Config ConfigList

func init() {
	var cfg, err = ini.Load("config.ini")
	if err != nil {
		log.Printf("Failed to read configuration file: %v", err)
		os.Exit(1)
	}

	Config = ConfigList{
		ApiKey:    cfg.Section("bitflyer").Key("api_key").String(),
		ApiSecret: cfg.Section("bitflyer").Key("api_secret").String(),
	}
}
