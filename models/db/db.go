package db

import (
	"github.com/shynome/cobweb/config"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func GetORM(db_config *gorm.Config) *gorm.DB {
	if db_config == nil {
		db_config = &gorm.Config{}
	}
	c := config.Get().Databases["default"]
	orm, err := gorm.Open(sqlite.Open(c.GetDSN()), db_config)
	if err != nil {
		panic(err)
	}
	return orm
}
