package config

import "os"

type V2rayShareConfig struct {
	UseDomain       string
	UsePort         string
	UseTLS          string
	UsePath         string
	UseRemarkPrefix string
}

func GetV2rayConfig() V2rayShareConfig {
	UseDomain := os.Getenv("USE_DOMAIN")
	UsePORT := os.Getenv("USE_PORT")
	UseTLS := os.Getenv("USE_TLS")
	UsePath := os.Getenv("USE_PATH")
	UseRemarkPrefix := os.Getenv("USE_REMARK_PREFIX")

	if UsePath == "" {
		UsePath = getV2rayUrl()
	}

	return V2rayShareConfig{
		UseDomain:       UseDomain,
		UsePort:         UsePORT,
		UsePath:         UsePath,
		UseTLS:          UseTLS,
		UseRemarkPrefix: UseRemarkPrefix,
	}
}
