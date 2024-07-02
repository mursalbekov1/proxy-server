package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
)

type Config struct {
	env        string `yaml:"env"`
	HttpServer `yaml:"http-server"`
}

type HttpServer struct {
	Port string `yaml:"port"`
	Host string `yaml:"host"`
}

func MustLoad() *Config {
	configPath := "config/config.yaml"

	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		log.Fatalf("cannot read config: %s", err)
	}

	return &cfg
}
