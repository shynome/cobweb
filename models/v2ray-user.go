package models

import (
	"gorm.io/gorm"
)

type V2rayUser struct {
	gorm.Model
	Username string `gorm:"uniqueIndex"`
	Uuid     string `gorm:"uniqueIndex"`
	Remark   string
}
