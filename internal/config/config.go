package config

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
	if configPath == "" {
		panic("CONFIG_PATH is required")
	}

	return &Config{}
}
