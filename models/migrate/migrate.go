package main

import (
	"log"

	"github.com/shynome/cobweb/models"
	"github.com/shynome/cobweb/models/db"
	"gorm.io/gorm"
)

func main() {
	orm := db.GetORM(&gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	err := orm.AutoMigrate(
		&models.V2rayUser{},
	)
	if err != nil {
		log.Fatal(err)
	}
}
