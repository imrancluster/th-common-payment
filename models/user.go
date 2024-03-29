package models

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"not null;unique_index" json:"username"`
	Email    string `gorm:"not null;unique_index" json:"email"`
	Password string `gorm:"not null" json:"password"`
}
