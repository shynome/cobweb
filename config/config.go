package config

import (
	"os"

	"github.com/GoAdminGroup/go-admin/modules/config"
)

func Get() *config.Config {
	// cfg := config.ReadFromYaml("./config.yaml")
	debug := os.Getenv("DEBUG") == "1"
	dbpath := os.Getenv("DB")
	if dbpath == "" {
		dbpath = "cobweb.db"
	}
	adminpath := os.Getenv("ADMIN")
	if adminpath == "" {
		adminpath = "admin"
	}

	cfg := &config.Config{
		Debug: debug,
		Databases: config.DatabaseList{
			"default": config.Database{
				Driver: "sqlite",
				File:   dbpath,
			},
		},
		UrlPrefix:     adminpath,
		AuthUserTable: "goadmin_users",
		Store: config.Store{
			Path:   "/tmp",
			Prefix: "uploads",
		},
	}
	return cfg
}
