package config

import (
	"os"
)

type ServerConfig struct {
	Listen   string
	V2rayUrl string
}

func GetServerConfig() ServerConfig {
	listen := os.Getenv("LISTEN")
	if listen == "" {
		listen = ":3005"
	}

	v2url := os.Getenv("V2RAY_URL")
	if v2url == "" {
		v2url = "/ray"
	}

	return ServerConfig{
		Listen:   listen,
		V2rayUrl: v2url,
	}
}
