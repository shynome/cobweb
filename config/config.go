package config

import "github.com/GoAdminGroup/go-admin/modules/config"

func Get() config.Config {
	cfg := config.ReadFromYaml("./config.yaml")
	return cfg
}
