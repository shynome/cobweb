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

	v2url := getV2rayUrl()
	return ServerConfig{
		Listen:   listen,
		V2rayUrl: v2url,
	}
}

func getV2rayUrl() (v2url string) {
	os.Getenv("V2RAY_URL")
	if v2url == "" {
		v2url = "/ray"
	}
	return
}
