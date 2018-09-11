package config

import (
	"github.com/BurntSushi/toml"
)

var (
	config Config
)

const (
	filepath = "./config/config.toml"
)

type Config struct {
	App AppConfig
	Rdb RdbConfig
}

type AppConfig struct {
	Name string
	Port int
}

type RdbConfig struct {
	Driver   string
	User     string
	Password string
	Schema   string
	Host     string
	Port     int
}

func init() {
	_, err := toml.DecodeFile(filepath, &config)
	if err != nil {
		panic(err)
	}
}

// GetAppConfig is return application config
func GetAppConfig() *AppConfig {
	return &config.App
}

// GetRdbConfig is return application config
func GetRdbConfig() *RdbConfig {
	return &config.Rdb
}
