package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"not null;type:varchar(100)"`
	Email    string `gorm:"not null;type:varchar(100)"`
	Password string `gorm:"not null;type:varchar(100)"`
	Role     uint   `gorm:"not null;"`
}

type Todos struct {
	gorm.Model
	Status      uint   `gorm:"not null"`
	Description string `gorm:"not null"`
}
