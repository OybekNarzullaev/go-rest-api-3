package apiserver

import "github.com/OybekNarzullaev/http-rest-api/store"

// config ...
type Config struct {
	BindAddr string `toml: "bind_addr"`
	LogLevel string `toml: "log_level"`
	Store *store.Config
}

func NewConfig() *Config {
	return &Config{
		BindAddr: ":8080",
		LogLevel: "debug",
		Store: store.NewConfig(),
	}
}
