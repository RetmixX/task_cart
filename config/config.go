package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"os"
)

type AppConf struct {
	Port   string `env:"SERVER_PORT"`
	Mode   string `env:"MODE"`
	DbConf DbConf
}

type DbConf struct {
	User     string `env:"DB_USER"`
	Password string `env:"DB_PASSWORD"`
	Port     int    `env:"DB_PORT"`
	Name     string `env:"DB_NAME"`
	Host     string `env:"DB_HOST"`
}

func MustLoad() *AppConf {
	configPath := ".deploy/local/.env"

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		panic("config file don't exists" + configPath)
	}

	var cfg AppConf

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		panic("cannot read config file: " + err.Error())
	}

	return &cfg
}
