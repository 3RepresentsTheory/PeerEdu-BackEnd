package config

import (
	"fmt"
	"github.com/jinzhu/configor"
	"os"
)

var Config = struct {
	Port        string
	Dsn         string
	ReleaseMode bool `yaml:"release_mode"`
	Key         string
	LoginExpire int64 `yaml:"login_expire"`
}{}

func Init() {
	path := "config.yml"
	if _, err := os.Stat(path); err != nil {
		fmt.Println("config.yml not found, copy config_example.yml to config.yml")
		os.Exit(1)
	}
	if err := configor.Load(&Config, path); err != nil {
		panic(err)
	}
}
