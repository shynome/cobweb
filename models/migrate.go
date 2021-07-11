package models

import "gorm.io/gorm"

func Migrate(orm *gorm.DB) error {
	return orm.AutoMigrate(
		&V2rayUser{},
	)
}
