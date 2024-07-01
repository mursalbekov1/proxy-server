package config

import "os"

type Config struct {
	env        string `yaml:"env"`
	HttpServer `yaml:"http-server"`
}

type HttpServer struct {
	port int `yaml:"port"`
}

func MustLoad() *Config {
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		panic("CONFIG_PATH is required")
	}

	return &Config{}
}
