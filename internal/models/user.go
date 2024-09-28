package models

import (
	"gorm.io/gorm"
)

type Users struct {
	gorm.Model
	ID       uint   `gorm:"primary_key" json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func (u *Users) TableName() string {
	return "users" // Change this to your table name
}
