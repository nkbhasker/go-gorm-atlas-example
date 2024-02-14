package models

import (
	"github.com/nkbhasker/go-gorm-atlas-example/enums"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name   string
	Email  string
	Gender *enums.GenderEnum `gorm:"type:gender"`
}
