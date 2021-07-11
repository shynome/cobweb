package models

import (
	"github.com/GoAdminGroup/go-admin/modules/db"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	orm *gorm.DB
	err error
)

func Init(c db.Connection) {
	orm, err = gorm.Open(sqlite.Open(c.GetConfig("default").GetDSN()), &gorm.Config{})

	if err != nil {
		panic("initialize orm failed")
	}
}

func GetORM() *gorm.DB {
	return orm
}
