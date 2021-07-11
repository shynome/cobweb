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
	err := models.Migrate(orm)
	if err != nil {
		log.Fatal(err)
	}
}
